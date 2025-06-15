# Backend Server Technical Specs
## Business Goal:
A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a Learnning record store (LRS) , providing correct and wrong score on practice vocabulary.
- A unified launchpad to launch different launching apps 

## Technical Requirements 
- The backend will be build using Go
- The database will be Sqlite3
- THe API will be built using Gin
- The API will always return JSON
- There will be no authentication or authorization

## Database Schema
Our database will be an Sqlit3 database called `words.db` that will be in the root of the project folder of `backend_go`
We have the following tables:
- words - stored vocabulary words
    - id integer
    - arabic word string
    - english word string
    - parts json
- word_groups - join table for words and groups many to many
    - id integer
    - word_id integer
    - group_id integer
- groups - thematic groups of words
    - id integer
    - name string
- study_sessions - records of study sessions grouping word_review_items
    - id integer
    - group_id integer
    - created_at datetime
    - study_activity_id integer
- study_activities - a specific study activity , linking a study session to a group 
    - id integer
    - study_session_id integer
    - group_id integer
    - created_at datetime
- word_review_items - a record of word practice, determining wether the word was correct or not.
    - word_id integer
    - study_session_id integer
    - correct boolean
    - created_at datetime


### API Endpoints
#### GET /api/dashboard/last_study_session
Returns information about the most recent study session
```json
{
  "id": integer,
  "created_at": datetime,
  "study_activity_id": integer,
  "group": {
    "id": integer,
    "name": string
  }
}
```

#### GET /api/dashboard/study_progress
Returns study progress statistics
Please note that the front end will determine progress percentage based on the total number of words available and the total number of words studied
```json
{
  "total_words_studied": integer,
  "total_available_words": integer,
}
```

#### GET /api/dashboard/quick_stats
Returns quick stats about the user's study progress
```json
{
  "success_rate": float,
  "total_study_sessions": integer,
  "total_active_groups": integer,
  "study_streak_days": integer
}
```

#### GET /api/study-activities/:id
Returns detailed information about a specific study activity.
```json
{
  "id": integer,
  "name": string,
  "description": string,
  "thumbnail_url": string,
  "launch_url": string,
  //"study_sessions": [
  //  {
    //  "id": integer,
    //  "group_id": integer,
    //  "created_at": datetime,
     // "word_review_items_count": integer,
    //  "start_time": datetime,
    //  "end_time": datetime
   // }
  //]
}
```

#### GET /api/study-activities/:id/study-sessions
Returns paginated list with 100 words per page
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": integer,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "group_id": integer,
      "created_at": datetime,
      "word_review_items_count": integer,
      "start_time": datetime,
      "end_time": datetime,
      "group": {
        "id": integer,
        "name": string
      }
    }
  ]
}
```

#### POST /api/study-activities
Creates a new study activity
##### Request Params
- group_id integer
- study_activity_id integer

##### Response
```json
{
  "id": integer,
  "group_id": integer,
  
}
```

#### GET /api/words
Returns paginated list of words
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": 100,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "arabic_word": string,
      "english_word": string,
      "correct_count": integer,
      "wrong_count": integer,
      "success_rate": float,
      "groups": [
        {
          "id": integer,
          "name": string
        }
      ]
    }
  ]
}
```

#### GET /api/words/:id
Returns detailed information about a specific word
```json
{
  "id": integer,
  "arabic_word": string,
  "english_word": string,
  "correct_count": integer,
  "wrong_count": integer,
  "success_rate": float,
  "groups": [
    {
      "id": integer,
      "name": string
    }
  ]
}
```

#### GET /api/groups
Returns paginated list of groups
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": 100,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "name": string,
      "word_count": integer
    }
  ]
}
```

#### GET /api/groups/:id
Returns detailed information about a specific group
```json
{
  "id": integer,
  "name": string,
  "total_word_count": integer,
  
}
```

#### GET /api/groups/:id/words
Returns paginated list of words in a specific group
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": 100,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "arabic_word": string,
      "english_word": string,
      "correct_count": integer,
      "wrong_count": integer,
      "success_rate": float
    }
  ]
}
```

#### GET /api/groups/:id/study-sessions
Returns paginated list of study sessions for a specific group
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": 100,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "group_id": integer,
      "created_at": datetime,
      "word_review_items_count": integer,
      "start_time": datetime,
      "end_time": datetime,
      "activity_name": string
    }
  ]
}
```

#### GET /api/study-sessions
Returns paginated list of study sessions
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": 100,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "group_id": integer,
      "study_activity_id": integer,
      "created_at": datetime,
      "word_review_items_count": integer,
      "start_time": datetime,
      "end_time": datetime,
      "activity_name": string
    }
  ]
}
```

#### GET /api/study-sessions/:id
Returns detailed information about a specific study session
```json
{
  "id": integer,
  "group_id": integer,
  "study_activity_id": integer,
  "created_at": datetime,
  "word_review_items_count": integer,
  "start_time": datetime,
  "end_time": datetime,
  "activity_name": string
}
```

#### GET /api/study-sessions/:id/words
Returns paginated list of words in a study session
```json
{
  "total_pages": integer,
  "current_page": integer,
  "items_per_page": 100,
  "total_items": integer,
  "items": [
    {
      "id": integer,
      "arabic_word": string,
      "english_word": string,
      "correct_count": integer,
      "wrong_count": integer,
      "success_rate": float
    }
  ]
}
```

#### POST /api/settings/reset-history
Reset study history
```json
{
  "message": "Study history reset successfully",
  "success": boolean
}
```

#### POST /api/settings/full-reset
Full database reset
```json
{
  "message": "Full database reset completed",
  "success": boolean
}
```

#### POST /api/study_sessions/:id/words/:word_id/review
##### Request Params
- id integer
- word_id integer
- study_session_id integer
##### Request Payload
```json
{
  "correct": boolean
}
##### Response
```json
{
  "message": "Word review updated successfully",
  "success": boolean,
  "word_id": integer,
  "study_session_id": integer,
  "correct": boolean,
  "created_at": datetime
}
```

## Mage Tasks
Mage is a task runner for Go.
Lets us list out possible tasks and run them.
### Initializing the database
This task will initalize the database called `words.db`.
### Migrate Database
This task will run a series of migrations sql files on the database.
### Seed Data
This task will import json files and transform them into target data for our database.

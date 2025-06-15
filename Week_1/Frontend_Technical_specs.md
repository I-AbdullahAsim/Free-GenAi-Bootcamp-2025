# FRONTEND TECHNICAL SPECS
## PAGES
### DASHBOARD `/dashboard`

#### PURPOSE
The purpose of this page is to provide a summary of learning and act as the default page when a user visits the web app.

#### COMPONENTS
- Last Study Session
    Shows last activity used
    Shows when last activity used
    summarizes wrong vs correct from last activity
    has a link to the group

- Study Progress
    - Total words study eg. 3/124
        -Across all study sessoins show the total words studied out of all the possible words in our database
    - display a mastery progress eg. 0%

- Quick Stats
    - Success rate eg 80%
    - Total study sessions eg. 4
    - total active groups eg. 3
    - study streak eg 4 days

- Start Studying Button
    - goes to show activities page


#### API ENDPOINTS
- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick_stats


### Study Activites Index `/study-activities`

#### PURPOSE
The purpose of this page is to show a collection of study activites with a thumbnail and its name to either launch or view the study activity.

#### COMPONENTS
- Study Activity List
    - shows a list of study activities with a thumbnail 
    - The name of the study activity to either launch or view the study activity.
    - A lauch button to take us to the launch page
    - The view page to view more information about past study sessions for this study activity
#### NEEDED API ENDPOINTS
- GET /api/study-activities

### STUDY ACTIVITY SHOW `/study-activities/:id`

#### PURPOSE
The purpose of this page is to show the details of the study activity and view past study sessions for this study activity.

#### COMPONENTS
- Name of study activity
- Thumbnail of study activity
- Description of study activity
- Launch Button
- Study Activities paginated lists
    - id
    - activity name
    - group name
    - start time
    - end time(inferred from the last word_review_item submitted)
    - number of review items 

    
#### NEEDED API ENDPOINTS
- GET /api/study-activities/:id
- GET /api/study-activities/:id/study-sessions

### Study Activity Launch `/study-activities/:id/launch`

#### PURPOSE
The purpose of this page is to launch a study activity
#### COMPONENTS
- Name of study activity

- Launch form
    - select feild or group
    - launch now button
#### BEHAVIOR
- After the form is submitted a new tab opens with the study activity based on its URL providded in the database.

Also after the form is submitted the page will redirect to the study session show page.
#### NEEDED API ENDPOINTS
- POST /api/study-activities
    
### WORDS INDEX `/words`

#### PURPOSE
The purpose of this page is to show all words in our database.
#### COMPONENTS
- Paginated Word List
    - Feilds
        - Arabic
        - English
        - Correct count
        - Wrong count
    - Pagination with 100 items per page.
    - Clicking the Arabic word will take us to the word show page.

#### NEEDED API ENDPOINTS
- GET /api/words

### WORD SHOW `/words/:id`

#### PURPOSE
The purpose of this page is to show the details of a specific word.
#### COMPONENTS
- Arabic
- English
- study Statistics
    - Correct count
    - Wrong count
    - Success rate
- Word Groups
    - Shown as a series of pills eg. tags
    - when group name is clicked it will take us to the group show page

#### NEEDED API ENDPOINTS
- GET /api/words/:id


### WORD GROUPS INDEX `/groups`

#### PURPOSE
The purpose of this page is to show a list of  groups in our database.
#### COMPONENTS
- Paginated Group List
    - Columns
        - Group name
        - Word count
    - Clicking the group name will take us to the group show page

#### NEEDED API ENDPOINTS
- GET /api/groups


### GROUP SHOW `/groups/:id`

#### PURPOSE
The purpose of this page is to show the details of a specific group.
#### COMPONENTS
- Group name
- Group statistics
    - Total word count
- Words in group (Paginated list of words)
    - should use the same component as the words index page
- study sessions (paginated list of study sessions)
    - should use the same component as the study sessions index page.    

#### NEEDED API ENDPOINTS
- GET /api/groups/:id ( the name and group stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study-sessions

### Study Session INDEX `/study-sessions`
#### PURPOSE
The purpose of this page is to show a list of study sessions in our database.
#### COMPONENTS
- Paginated Study Session List
    - Columns
        - ID
        - Activity name
        - Group name
        - Start time
        - End time
        - Number of review items
    - Clicking the study session id name will take us to the study session show page

#### NEEDED API ENDPOINTS
- GET /api/study-sessions


### STUDY SESSIONS SHOW `/study-sessions/:id`
#### PURPOSE
The purpose of this page is to show the details of a specific study session.
#### COMPONENTS
- Study Session details
    - Activity name
    - Group name
    - Start time
    - End time
    - Number of review items

- Words Review items (Paginated list of words)
    - should use the same component as the words index page

#### NEEDED API ENDPOINTS
- GET /api/study-sessions/:id
- GET /api/study-sessions/:id/words


### SETTINGS `/settings`
#### PURPOSE
The purpose of this page is to mae configurations to the study portal.
#### COMPONENTS
- Theme selection eg. light/dark
- Language selection eg. English/Arabic
- Reset History Button
    - This will reset the history of the study portal.
    - This will delete all the study sessions and word review items.
- Full Reset Button
    - This will drop all tables and recreate them with seed data.

#### NEEDED API ENDPOINTS
- POST /api/settings/reset-history
- POST /api/settings/full-reset
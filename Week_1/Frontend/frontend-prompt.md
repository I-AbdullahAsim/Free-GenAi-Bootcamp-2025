# FRONT END
We would like to build a Arabic Learning Web App

## Role / Professoin
- Frontend Developer

## Project Description
### Project Breif
We are building an Arabic learning web app that serves the following purposes:
- A portal to launch study activities
- To store, group and explore Arabic Vocabulary
- To review study progress

The web app is intened for desktop only so we dont have to be concerned with mobile layouts

### Technical Requirements 
- React.js as frontend library
- Tailwind CSS for styling
- Vite.js as the local developement server
- TypeScript for the programming language
- ShadCN for components

### Frontend Routes
This is a list of routes for our web-app we are building.
Each of these routes are a page and we will describe them in more details under the pages heading.

/dashboard
/study-activities
/study-activities/:id
/words
/words/:id
/groups
/groups/:id
/sessions
/settings

The default route / should be forwarded to /dashboard

### Global Components
#### Navigation
There will be a horizontal navigation bar with the following links:
- Dashboard
- Study Activities
- Words
- Words Groups
- Sessions
- Settings

#### BreadCrumbs
Beneath the navigation bar there will be a bread crumb that shows the current route. Example of breadcrumbs:
- Dashboard
- Study Activities -> Adventure MUD
- Words -> مرحباً
- Word Groups -> Core Words

### Pages
#### Dashboard
This page provides a summary of the Student's Progression
- Last Sessoin
#### Study Activites Index
The route for this page is /study-activities
This is a grade of cards which represents an activity.
A card has a :
- Thumbnail
- Title
- "Launch" button
- "View" button

The Launch button will open a new tab. Study activities are their own apps, but in order for them to learn they need to be provided a group_id.

eg. localhost:8081?group_id=4

THis page requires no pagination because it is unlikely to be more then 20 possible study activities

THe view button will go to Stuent Activities Show page.

#### Study Activites Show
The route for this page is /study-activities/:id

This page will show the details of a study activity.
It will have :
- Thumbnail
- Tittle
- Description
- "Launch" button


There will be a list of sessions for this activity.
- a session item will contain:
- Group name : So you know what group name was used for the sessions.
- There will be a link to the group show page.
- Start time:  When the session was created : YYYY-MM-DD HH:MM:SS format (12 hours)
- End time: When the last_word_review_item was created.
- Review items: The number of review items.

Words Index
The route for this page is /words
This is a list of tables in the following cells:
- Arabic : The arabic word
    1) This will also contain a small button for playing the sound of the word
    2) There will be a link to the word show page
- English : The english translation of the word
- Correct : Number of correct word review items
- Incorrect : Number of incorrect word review items

There should only be 50 words displayed at a time.

There needs to be pagination
- Previous button: grey out if you cannot go further back
- Page 1 of 3: With the current page bolded:
- Next button : greyed out if you cannot go further forward
All table headings should be sortable. If you click it will toggle between ascending and descending order.
An ascii arrow should indicate the direction and the column being sorted with an arrow pointing up or down.

### Words Show
The route for this page is /words/:id
This is a table of word groups with the following cells:
- Group name : The name of the group
    1) This will be a link to the word groups show page
- Words : The number of words in the group

This page contains the same sorting and pagination logic as the Words in the index page.

### Word Groups Show
The route for this page is /word-groups/:id
This has the same components as the Words index but it is scoped to only show words that are associated with this group

### Sessions Index
The route for this page is /sessions

This page will show a list of sessions similar to the Stdy Activities show page

This paeg also contains the same sorting and pagination logic as words in the index page

### Settings
The route for this page is /settings
This page will show a list of settings

Reset History Button:
This has a button that allows us to reset the entire database.
We need to confirm this action is a dialog and type the word reset me to to confirm the action.
Dark Mode Toggle : This is a toggle that changes light to dark theme


✅ 1. Database Initialization & Seeding
Tests:

 Test_DB_Migration_Execution_Order: Migration files execute in correct order (0001_, 0002_, …).

 Test_DB_Schema_Creation: All expected tables are created.

 Test_Seed_Data_Import: Seed JSON files are loaded into words, groups, word_groups correctly.

 Test_Seed_Group_Mapping: Words are assigned to the correct groups based on seed DSL.

✅ 2. Words API
GET /api/words
 Returns paginated list of words (items_per_page = 100).

 Words include correct id, arabic_word, english_word, correct_count, wrong_count, success_rate, and associated groups.

 Returns empty list when DB has no words.

GET /api/words/:id
 Returns correct word details for valid :id.

 Returns 404 for invalid :id.

✅ 3. Groups API
GET /api/groups
 Returns paginated group list with word_count.

GET /api/groups/:id
 Returns detailed info including total_word_count.

 Returns 404 for invalid group ID.

GET /api/groups/:id/words
 Returns paginated list of words associated with a group.

 Validates correct success_rate & review data per word.

GET /api/groups/:id/study-sessions
 Returns sessions filtered by group.

 Returns activity name per session.

✅ 4. Study Sessions API
GET /api/study-sessions
 Returns paginated list of study sessions.

 Validates word_review_items_count, start_time, end_time, and activity_name.

GET /api/study-sessions/:id
 Returns detailed study session by ID.

 Returns 404 for non-existent ID.

GET /api/study-sessions/:id/words
 Returns words studied in this session.

 Validates correct_count, wrong_count, success_rate.

✅ 5. Study Activity API
GET /api/study-activities/:id
 Returns detailed study activity info.

 Fields: name, description, thumbnail_url, launch_url.

GET /api/study-activities/:id/study-sessions
 Returns sessions under this activity (paginated).

 Validates each session's group and word count.

POST /api/study-activities
 Creates a new study activity.

 Requires valid group_id and study_activity_id.

 Returns new activity id.

✅ 6. Dashboard Endpoints
GET /api/dashboard/last_study_session
 Returns latest session info (check created_at ordering).

 Includes group data.

GET /api/dashboard/study_progress
 Returns correct total_words_studied and total_available_words.

GET /api/dashboard/quick_stats
 Calculates:

success_rate = total correct / total attempts

total_study_sessions = number of sessions

total_active_groups = groups involved in sessions

study_streak_days = distinct active days

✅ 7. Study History Review API
POST /api/study_sessions/:id/words/:word_id/review
 Validates payload with "correct": boolean.

 Inserts new word_review_item row.

 Returns created_at and success message.

✅ 8. Settings Endpoints
POST /api/settings/reset-history
 Clears only review data (word_review_items, study_sessions, study_activities).

 Words and groups remain.

POST /api/settings/full-reset
 Resets full DB: drops and re-creates schema.

 All tables cleared.

 Seeds reloaded (verify basic_words, animals, etc.).

✅ 9. JSON Format Tests
 All API responses must have:

"message" field if applicable.

"success": true|false.

"data" or top-level object representing the expected resource.

✅ 10. Negative & Edge Case Testing
 Invalid IDs (/api/words/999) return 404.

 Empty database does not crash any endpoint.

 Pagination handles pages beyond range (e.g. page=9999 returns empty list, not error).

 Malformed requests return 400 with JSON error.


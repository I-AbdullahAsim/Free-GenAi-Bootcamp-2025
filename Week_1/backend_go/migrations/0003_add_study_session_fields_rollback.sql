-- Remove additional fields from study_sessions table

PRAGMA foreign_keys = OFF;

CREATE TABLE study_sessions__temp (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  activity_id INTEGER NOT NULL,
  created_at DATETIME,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (activity_id) REFERENCES study_activities(id) ON DELETE CASCADE
);

INSERT INTO study_sessions__temp (
  id, user_id, activity_id, created_at, updated_at, deleted_at
)
SELECT id, user_id, activity_id, created_at, updated_at, deleted_at FROM study_sessions;

DROP TABLE study_sessions;
ALTER TABLE study_sessions__temp RENAME TO study_sessions;

PRAGMA foreign_keys = ON;

-- Remove additional fields from study_sessions table
ALTER TABLE study_sessions DROP COLUMN start_time;
ALTER TABLE study_sessions DROP COLUMN end_time;

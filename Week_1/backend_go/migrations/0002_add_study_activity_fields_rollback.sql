PRAGMA foreign_keys = OFF;

CREATE TABLE study_activities_temp (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    study_session_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (study_session_id) REFERENCES study_sessions(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

INSERT INTO study_activities_temp (id, study_session_id, group_id, created_at, updated_at, deleted_at)
SELECT id, study_session_id, group_id, created_at, updated_at, deleted_at FROM study_activities;

DROP TABLE study_activities;

ALTER TABLE study_activities_temp RENAME TO study_activities;

PRAGMA foreign_keys = ON;

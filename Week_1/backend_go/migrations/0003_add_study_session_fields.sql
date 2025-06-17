-- Add additional fields to study_sessions table if they don't exist
ALTER TABLE study_sessions ADD COLUMN start_time TIMESTAMP DEFAULT NULL;
ALTER TABLE study_sessions ADD COLUMN end_time TIMESTAMP DEFAULT NULL;

-- Update existing records with default values
UPDATE study_sessions 
SET start_time = created_at,
    end_time = datetime(created_at, '+30 minutes')
WHERE start_time IS NULL;

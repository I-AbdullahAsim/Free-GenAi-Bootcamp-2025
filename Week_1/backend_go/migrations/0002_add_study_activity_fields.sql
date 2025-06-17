-- Add additional fields to study_activities table
ALTER TABLE study_activities ADD COLUMN name TEXT;
ALTER TABLE study_activities ADD COLUMN description TEXT;
ALTER TABLE study_activities ADD COLUMN thumbnail_url TEXT;
ALTER TABLE study_activities ADD COLUMN launch_url TEXT;

-- Update existing records with default values
UPDATE study_activities 
SET name = 'Default Activity',
    description = 'Default activity description',
    thumbnail_url = '/static/default-thumbnail.png',
    launch_url = '/study/default'
WHERE name IS NULL;

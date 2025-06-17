-- Add statistics fields to words table
ALTER TABLE words ADD COLUMN correct_count INTEGER DEFAULT 0;
ALTER TABLE words ADD COLUMN wrong_count INTEGER DEFAULT 0;

-- Index for stats
CREATE INDEX IF NOT EXISTS idx_words_stats ON words(correct_count, wrong_count);

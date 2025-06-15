-- Create words table
CREATE TABLE IF NOT EXISTS words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    arabic_word TEXT NOT NULL,
    english_word TEXT NOT NULL,
    parts JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Add indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_words_arabic ON words(arabic_word);
CREATE INDEX IF NOT EXISTS idx_words_english ON words(english_word);

-- Add foreign key constraint to word_groups table
ALTER TABLE word_groups
ADD CONSTRAINT fk_word_groups_words
FOREIGN KEY (word_id) REFERENCES words(id)
ON DELETE CASCADE;
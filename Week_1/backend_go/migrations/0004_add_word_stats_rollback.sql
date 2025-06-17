-- Remove statistics fields from words table

PRAGMA foreign_keys = OFF;

CREATE TABLE words__temp (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  arabic_word TEXT NOT NULL,
  english_word TEXT NOT NULL,
  parts TEXT,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

INSERT INTO words__temp (
  id, arabic_word, english_word, parts, created_at, updated_at, deleted_at
)
SELECT id, arabic_word, english_word, parts, created_at, updated_at, deleted_at FROM words;

DROP TABLE words;
ALTER TABLE words__temp RENAME TO words;

PRAGMA foreign_keys = ON;

-- Drop stats index
DROP INDEX IF EXISTS idx_words_stats;

DROP INDEX IF EXISTS idx_highlights_event_id;

DROP TRIGGER IF EXISTS set_updated_at ON highlights;

DROP TABLE IF EXISTS highlights CASCADE;


DROP INDEX IF EXISTS idx_events_championship_id;

DROP TRIGGER IF EXISTS set_updated_at ON events;

DROP TABLE IF EXISTS events CASCADE;


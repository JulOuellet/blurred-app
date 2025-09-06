DROP INDEX IF EXISTS idx_championships_season_id;

DROP TRIGGER IF EXISTS set_updated_at ON championships;

DROP TABLE IF EXISTS championships CASCADE;


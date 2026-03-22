-- ============================================================
-- Redesign integrations: channel+championship → channel+sport
-- ============================================================

-- 1. Delete all existing integrations (cascades to youtube_inbox)
-- The old schema is channel+championship; the new schema is channel+sport.
-- All integrations must be recreated with the new schema.
DELETE FROM integrations;

-- 2. Alter integrations table
DROP INDEX IF EXISTS idx_integrations_championship_id;
ALTER TABLE integrations DROP CONSTRAINT IF EXISTS uq_channel_championship;
ALTER TABLE integrations DROP CONSTRAINT IF EXISTS fk_championship;

ALTER TABLE integrations
    DROP COLUMN championship_id,
    DROP COLUMN relevance_pattern,
    DROP COLUMN event_pattern;

ALTER TABLE integrations
    ADD COLUMN sport_id uuid NOT NULL DEFAULT 'e89d82d8-4bcc-423c-bbd0-18dbd0c1da01',
    ADD COLUMN content_filter text,
    ADD COLUMN title_exclude text,
    ADD COLUMN stage_pattern text;

ALTER TABLE integrations ALTER COLUMN sport_id DROP DEFAULT;

ALTER TABLE integrations
    ADD CONSTRAINT fk_sport FOREIGN KEY (sport_id) REFERENCES sports (id) ON DELETE CASCADE,
    ADD CONSTRAINT uq_channel_sport UNIQUE (youtube_channel_id, sport_id);

CREATE INDEX idx_integrations_sport_id ON integrations (sport_id);

-- 3. Add title_pattern to championships
ALTER TABLE championships ADD COLUMN title_pattern text;

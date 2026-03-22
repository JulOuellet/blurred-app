-- Reverse: redesign integrations back to channel+championship

-- 1. Remove title_pattern from championships
ALTER TABLE championships DROP COLUMN IF EXISTS title_pattern;

-- 2. Reverse integrations table changes
DROP INDEX IF EXISTS idx_integrations_sport_id;
ALTER TABLE integrations DROP CONSTRAINT IF EXISTS uq_channel_sport;
ALTER TABLE integrations DROP CONSTRAINT IF EXISTS fk_sport;

ALTER TABLE integrations
    DROP COLUMN sport_id,
    DROP COLUMN content_filter,
    DROP COLUMN title_exclude,
    DROP COLUMN stage_pattern;

ALTER TABLE integrations
    ADD COLUMN championship_id uuid NOT NULL DEFAULT '10000000-2026-4000-a000-000000000021',
    ADD COLUMN relevance_pattern text NOT NULL DEFAULT '(?i).*',
    ADD COLUMN event_pattern text;

ALTER TABLE integrations ALTER COLUMN championship_id DROP DEFAULT;
ALTER TABLE integrations ALTER COLUMN relevance_pattern DROP DEFAULT;

ALTER TABLE integrations
    ADD CONSTRAINT fk_championship FOREIGN KEY (championship_id) REFERENCES championships (id) ON DELETE CASCADE,
    ADD CONSTRAINT uq_channel_championship UNIQUE (youtube_channel_id, championship_id);

CREATE INDEX idx_integrations_championship_id ON integrations (championship_id);

-- 3. Re-seed integrations
INSERT INTO integrations (id, youtube_channel_id, youtube_channel_name, championship_id, lang, relevance_pattern, event_pattern)
VALUES
    ('30000000-2026-4000-a000-000000000001', 'UCfDfvvMARk4TKcC62ALi6eA', 'TNT Sports Cycling', '10000000-2026-4000-a000-000000000021', 'en-GB', '(?i)UAE Tour.*Stage.*Highlights', '(?i)Stage\s+(\d+)'),
    ('30000000-2026-4000-a000-000000000002', 'UC77UtoyivVHkpApL0wGfH5w', 'Lanterne Rouge Cycling', '10000000-2026-4000-a000-000000000021', 'en-GB', '(?i)UAE Tour\s+\d{4}\s+Stage\s+\d+', '(?i)Stage\s+(\d+)');

-- Fix integration lang from 'en' (invalid) to 'en-GB'
UPDATE integrations
SET lang = 'en-GB'
WHERE id = '30000000-2026-4000-a000-000000000001';

-- Fix highlights created from this integration
UPDATE highlights
SET lang = 'en-GB'
WHERE youtube_id IN (
    SELECT youtube_video_id
    FROM youtube_inbox
    WHERE integration_id = '30000000-2026-4000-a000-000000000001'
      AND status = 'completed'
);

-- Migrate highlight source from 'official' to the channel name
UPDATE highlights
SET source = 'TNT Sports Cycling'
WHERE source = 'official'
  AND youtube_id IN (
    SELECT youtube_video_id
    FROM youtube_inbox
    WHERE integration_id = '30000000-2026-4000-a000-000000000001'
      AND status = 'completed'
);

-- Drop redundant source column (youtube_channel_name is used instead)
ALTER TABLE integrations DROP COLUMN source;

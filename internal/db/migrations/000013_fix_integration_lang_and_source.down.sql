-- Re-add source column
ALTER TABLE integrations ADD COLUMN source varchar(255);

-- Restore source for TNT integration
UPDATE integrations
SET source = 'official'
WHERE id = '30000000-2026-4000-a000-000000000001';

-- Revert highlight source back to 'official'
UPDATE highlights
SET source = 'official'
WHERE source = 'TNT Sports Cycling'
  AND youtube_id IN (
    SELECT youtube_video_id
    FROM youtube_inbox
    WHERE integration_id = '30000000-2026-4000-a000-000000000001'
      AND status = 'completed'
);

-- Revert lang
UPDATE integrations
SET lang = 'en'
WHERE id = '30000000-2026-4000-a000-000000000001';

UPDATE highlights
SET lang = 'en'
WHERE youtube_id IN (
    SELECT youtube_video_id
    FROM youtube_inbox
    WHERE integration_id = '30000000-2026-4000-a000-000000000001'
      AND status = 'completed'
);

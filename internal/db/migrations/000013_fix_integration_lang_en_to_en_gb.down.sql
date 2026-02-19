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

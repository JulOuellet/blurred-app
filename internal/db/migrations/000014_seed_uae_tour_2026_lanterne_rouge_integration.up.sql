-- ============================================================
-- UAE Tour 2026 – Lanterne Rouge Cycling Integration
-- Championship ID: 10000000-2026-4000-a000-000000000021
-- Channel: @LanterneRougeCycling
-- ============================================================
INSERT INTO integrations (
    id,
    youtube_channel_id,
    youtube_channel_name,
    championship_id,
    lang,
    relevance_pattern,
    event_pattern)
VALUES (
    '30000000-2026-4000-a000-000000000002',
    'UC77UtoyivVHkpApL0wGfH5w',
    'Lanterne Rouge Cycling',
    '10000000-2026-4000-a000-000000000021',
    'en-GB',
    '(?i)UAE Tour\s+\d{4}\s+Stage\s+\d+',
    '(?i)Stage\s+(\d+)');

-- ============================================================
-- UAE Tour 2026 – TNT Sports Cycling Integration
-- Championship ID: 10000000-2026-4000-a000-000000000021
-- Channel: @TNTSportsCycling
-- ============================================================
INSERT INTO integrations (
    id,
    youtube_channel_id,
    youtube_channel_name,
    championship_id,
    lang,
    source,
    relevance_pattern,
    event_pattern)
VALUES (
    '30000000-2026-4000-a000-000000000001',
    'UCfDfvvMARk4TKcC62ALi6eA',
    'TNT Sports Cycling',
    '10000000-2026-4000-a000-000000000021',
    'en',
    'official',
    '(?i)UAE Tour.*Stage.*Highlights',
    '(?i)Stage\s+(\d+)');

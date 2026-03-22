-- ============================================================
-- Seed title_pattern for all championships
-- Patterns are case-insensitive regexes to match video titles.
-- They handle common dash variants (-, –, —) and spacing.
-- ============================================================

-- 2025 season
UPDATE championships SET title_pattern = '(?i)Tour[\s\-–—]*de[\s\-–—]*France'
WHERE id = '4a965e4a-037d-4cb1-9ac1-499a8b5a93bc';

-- 2026 season
UPDATE championships SET title_pattern = '(?i)UAE[\s\-–—]*Tour'
WHERE id = '10000000-2026-4000-a000-000000000021';

UPDATE championships SET title_pattern = '(?i)Omloop[\s\-–—]*Het[\s\-–—]*Nieuwsblad'
WHERE id = '10000000-2026-4000-a000-000000000001';

UPDATE championships SET title_pattern = '(?i)Strade[\s\-–—]*Bianche'
WHERE id = '10000000-2026-4000-a000-000000000002';

UPDATE championships SET title_pattern = '(?i)Paris[\s\-–—]*Nice'
WHERE id = '10000000-2026-4000-a000-000000000003';

UPDATE championships SET title_pattern = '(?i)Tirreno[\s\-–—]*Adriatico'
WHERE id = '10000000-2026-4000-a000-000000000004';

UPDATE championships SET title_pattern = '(?i)Milan[oa]?[\s\-–—]*San[\s\-–—]*Remo'
WHERE id = '10000000-2026-4000-a000-000000000005';

UPDATE championships SET title_pattern = '(?i)E3[\s\-–—]*(Saxo|Harelbeke|Classic)'
WHERE id = '10000000-2026-4000-a000-000000000017';

UPDATE championships SET title_pattern = '(?i)Gent[\s\-–—]*Wevelgem'
WHERE id = '10000000-2026-4000-a000-000000000018';

UPDATE championships SET title_pattern = '(?i)Tour[\s\-–—]*of[\s\-–—]*Flanders|Ronde[\s\-–—]*van[\s\-–—]*Vlaanderen'
WHERE id = '10000000-2026-4000-a000-000000000006';

UPDATE championships SET title_pattern = '(?i)Itzulia|Basque[\s\-–—]*Country'
WHERE id = '10000000-2026-4000-a000-000000000020';

UPDATE championships SET title_pattern = '(?i)Paris[\s\-–—]*Roubaix'
WHERE id = '10000000-2026-4000-a000-000000000007';

UPDATE championships SET title_pattern = '(?i)Amstel[\s\-–—]*Gold'
WHERE id = '10000000-2026-4000-a000-000000000008';

UPDATE championships SET title_pattern = '(?i)Fl[èeé]che[\s\-–—]*Wallonne'
WHERE id = '10000000-2026-4000-a000-000000000019';

UPDATE championships SET title_pattern = '(?i)Li[èeé]ge[\s\-–—]*Bastogne[\s\-–—]*Li[èeé]ge'
WHERE id = '10000000-2026-4000-a000-000000000009';

UPDATE championships SET title_pattern = '(?i)Giro[\s\-–—]*d[''`'']?[\s\-–—]*Italia'
WHERE id = '10000000-2026-4000-a000-000000000010';

UPDATE championships SET title_pattern = '(?i)Crit[éeè]rium[\s\-–—]*du[\s\-–—]*Dauphin[éeè]|Dauphin[éeè]'
WHERE id = '10000000-2026-4000-a000-000000000011';

UPDATE championships SET title_pattern = '(?i)Tour[\s\-–—]*de[\s\-–—]*Suisse'
WHERE id = '10000000-2026-4000-a000-000000000012';

UPDATE championships SET title_pattern = '(?i)Tour[\s\-–—]*de[\s\-–—]*France'
WHERE id = '10000000-2026-4000-a000-000000000013';

UPDATE championships SET title_pattern = '(?i)Vuelta[\s\-–—]*(a[\s\-–—]*Espa[ñn]a)?'
WHERE id = '10000000-2026-4000-a000-000000000014';

UPDATE championships SET title_pattern = '(?i)GP[\s\-–—]*Qu[éeè]bec|Grand[\s\-–—]*Prix.*Qu[éeè]bec'
WHERE id = '10000000-2026-4000-a000-000000000022';

UPDATE championships SET title_pattern = '(?i)GP[\s\-–—]*Montr[éeè]al|Grand[\s\-–—]*Prix.*Montr[éeè]al'
WHERE id = '10000000-2026-4000-a000-000000000023';

UPDATE championships SET title_pattern = '(?i)World[\s\-–—]*Championships|Worlds'
WHERE id = '10000000-2026-4000-a000-000000000015';

UPDATE championships SET title_pattern = '(?i)(Il[\s\-–—]*)?Lombardia'
WHERE id = '10000000-2026-4000-a000-000000000016';

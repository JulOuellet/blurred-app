DELETE FROM events
WHERE championship_id IN (
    '10000000-2026-4000-a000-000000000020', -- Itzulia Basque Country
    '10000000-2026-4000-a000-000000000010', -- Giro d'Italia
    '10000000-2026-4000-a000-000000000011', -- Critérium du Dauphiné
    '10000000-2026-4000-a000-000000000012', -- Tour de Suisse
    '10000000-2026-4000-a000-000000000013', -- Tour de France
    '10000000-2026-4000-a000-000000000014'  -- Vuelta a España
);

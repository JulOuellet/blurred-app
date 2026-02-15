INSERT INTO seasons (
    id,
    name,
    start_date,
    end_date,
    sport_id)
VALUES (
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    '2026',
    '2026-01-01T00:00:00Z',
    '2026-12-31T23:59:59Z',
    'e89d82d8-4bcc-423c-bbd0-18dbd0c1da01');

-- ============================================================
-- Championships (one per race/competition)
-- ============================================================
-- UAE Tour (stage race, ~7 days)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000021',
    'UAE Tour',
    'UCI World Tour',
    '2026-02-16T00:00:00Z',
    '2026-02-22T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 UAE Tour, a one-week stage race in the United Arab Emirates.',
    'https://images.unsplash.com/photo-1753703986159-7dd19429b1c0');

-- Omloop Het Nieuwsblad (one-day)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000001',
    'Omloop Het Nieuwsblad',
    'UCI World Tour',
    '2026-02-28T00:00:00Z',
    '2026-02-28T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Omloop Het Nieuwsblad, a Spring Classic and the traditional opener of the European road season.',
    'https://upload.wikimedia.org/wikipedia/commons/8/83/Muur-van-Geraardsbergen.jpg');

-- 2. Strade Bianche (one-day)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000002',
    'Strade Bianche',
    'UCI World Tour',
    '2026-03-07T00:00:00Z',
    '2026-03-07T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Strade Bianche, a Spring Classic on the white gravel roads of Tuscany.',
    'https://upload.wikimedia.org/wikipedia/commons/d/de/Strade_Bianche-26-2_%2851964437028%29.jpg');

-- 3. Paris–Nice (stage race, ~8 days)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000003',
    'Paris–Nice',
    'UCI World Tour',
    '2026-03-08T00:00:00Z',
    '2026-03-15T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Paris–Nice, a one-week stage race from Paris to the French Riviera.',
    'https://upload.wikimedia.org/wikipedia/commons/4/43/Eze-Village-PACA-France.jpg');

-- 4. Tirreno–Adriatico (stage race, ~7 days)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000004',
    'Tirreno–Adriatico',
    'UCI World Tour',
    '2026-03-09T00:00:00Z',
    '2026-03-15T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Tirreno–Adriatico, a one-week stage race across Italy.',
    'https://images.unsplash.com/photo-1531686888376-83ee7d64f5eb');

-- 5. Milano–San Remo (one-day, Monument #1)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000005',
    'Milano–San Remo',
    'UCI World Tour',
    '2026-03-21T00:00:00Z',
    '2026-03-21T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Milano–San Remo, a Monument and the longest one-day race in professional cycling.',
    'https://images.unsplash.com/photo-1697532355892-7af404c638b2');

-- 6. E3 Saxo Classic (one-day, Cobbled Classic)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000017',
    'E3 Saxo Classic',
    'UCI World Tour',
    '2026-03-27T00:00:00Z',
    '2026-03-27T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 E3 Saxo Classic, a Spring Classic on the cobbled hills of Flanders.',
    'https://upload.wikimedia.org/wikipedia/commons/b/b1/Paterberg_in_de_gelijknamige_straat_-_Belgi%C3%AB.jpg');

-- 7. Gent-Wevelgem (one-day, Cobbled Classic)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000018',
    'Gent–Wevelgem',
    'UCI World Tour',
    '2026-03-29T00:00:00Z',
    '2026-03-29T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Gent–Wevelgem, a Spring Classic in Flanders featuring the iconic Kemmelberg.',
    'https://images.unsplash.com/photo-1578345728913-ce0d20741dc0');

-- 8. Tour of Flanders (one-day, Monument #2)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000006',
    'Tour of Flanders',
    'UCI World Tour',
    '2026-04-05T00:00:00Z',
    '2026-04-05T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Tour of Flanders, a Monument and the biggest one-day race in Belgian cycling.',
    'https://upload.wikimedia.org/wikipedia/commons/c/ca/Melden_-_Koppenberg_-_View.jpg');

-- Itzulia Basque Country (stage race, ~6 days)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000020',
    'Itzulia Basque Country',
    'UCI World Tour',
    '2026-04-06T00:00:00Z',
    '2026-04-11T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Itzulia Basque Country, a one-week stage race through the mountains of the Basque Country.',
    'https://upload.wikimedia.org/wikipedia/commons/a/ad/Txindoki_mountain_from_Lazkaomendi.jpg');

-- Paris–Roubaix (one-day, Monument #3)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000007',
    'Paris–Roubaix',
    'UCI World Tour',
    '2026-04-12T00:00:00Z',
    '2026-04-12T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Paris–Roubaix, a Monument known as the Hell of the North for its brutal cobblestone sectors.',
    'https://images.unsplash.com/photo-1748940452876-f58f4ee63a6c');

-- 8. Amstel Gold Race (one-day, Ardennes Classic)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000008',
    'Amstel Gold Race',
    'UCI World Tour',
    '2026-04-19T00:00:00Z',
    '2026-04-19T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Amstel Gold Race, a Spring Classic and the first of the Ardennes trilogy in the Netherlands.',
    'https://images.unsplash.com/photo-1597513033476-7fa6d9628cdd');

-- 11. La Flèche Wallonne (one-day, Ardennes Classic)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000019',
    'La Flèche Wallonne',
    'UCI World Tour',
    '2026-04-22T00:00:00Z',
    '2026-04-22T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 La Flèche Wallonne, a Spring Classic famous for its finish atop the Mur de Huy.',
    'https://upload.wikimedia.org/wikipedia/commons/e/e6/2017_Mur_de_Huy_4.jpg');

-- 12. Liège–Bastogne–Liège (one-day, Monument #4)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000009',
    'Liège–Bastogne–Liège',
    'UCI World Tour',
    '2026-04-26T00:00:00Z',
    '2026-04-26T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Liège–Bastogne–Liège, a Monument and the oldest classic in professional cycling.',
    'https://upload.wikimedia.org/wikipedia/commons/3/32/Cote_de_la_Redoute.jpg');

-- 10. Giro d'Italia (Grand Tour, ~3 weeks)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000010',
    'Giro d''Italia',
    'UCI World Tour',
    '2026-05-09T00:00:00Z',
    '2026-05-31T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of all 21 stages of the 2026 Giro d''Italia, the first Grand Tour of the season.',
    'https://images.unsplash.com/photo-1712729946173-eb9794baba50');

-- 11. Critérium du Dauphiné (stage race, ~8 days)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000011',
    'Critérium du Dauphiné',
    'UCI World Tour',
    '2026-06-07T00:00:00Z',
    '2026-06-14T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Critérium du Dauphiné, a one-week stage race in the French Alps and key preparation for the Tour de France.',
    'https://images.unsplash.com/photo-1666870981668-077673495e31');

-- 12. Tour de Suisse (stage race, ~9 days)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000012',
    'Tour de Suisse',
    'UCI World Tour',
    '2026-06-17T00:00:00Z',
    '2026-06-25T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Tour de Suisse, a one-week stage race through the Swiss Alps.',
    'https://images.unsplash.com/photo-1594741210511-baa4ccb7698d');

-- 13. Tour de France (Grand Tour, ~3 weeks)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000013',
    'Tour de France',
    'UCI World Tour',
    '2026-07-04T00:00:00Z',
    '2026-07-26T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of all 21 stages of the 2026 Tour de France, the most prestigious Grand Tour in cycling. Grand Départ from Barcelona, Spain.',
    'https://images.unsplash.com/photo-1710330362971-83ec53835f92');

-- 14. Vuelta a España (Grand Tour, ~3 weeks)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000014',
    'Vuelta a España',
    'UCI World Tour',
    '2026-08-22T00:00:00Z',
    '2026-09-13T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of all 21 stages of the 2026 Vuelta a España, the final Grand Tour of the season. Grand Départ from Monaco.',
    'https://images.unsplash.com/photo-1695068446447-d8a646a60585');

-- Grand Prix Cycliste de Québec (one-day)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000022',
    'Grand Prix Cycliste de Québec',
    'UCI World Tour',
    '2026-09-11T00:00:00Z',
    '2026-09-11T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Grand Prix Cycliste de Québec, a one-day race on the circuit of the Plains of Abraham.',
    'https://images.unsplash.com/photo-1660712319764-1badd42b2bc3');

-- Grand Prix Cycliste de Montréal (one-day)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000023',
    'Grand Prix Cycliste de Montréal',
    'UCI World Tour',
    '2026-09-13T00:00:00Z',
    '2026-09-13T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Grand Prix Cycliste de Montréal, a one-day race on the circuit of Mont Royal.',
    'https://images.unsplash.com/photo-1715191307694-ee4e57b473d6');

-- UCI Road World Championships
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000015',
    'UCI Road World Championships',
    'UCI Events',
    '2026-09-20T00:00:00Z',
    '2026-09-27T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 UCI Road World Championships in Montreal, Canada, where national teams compete for the rainbow jersey.',
    'https://images.unsplash.com/photo-1575540576545-3db8561c29e4');

-- 16. Il Lombardia (one-day, Monument #5)
INSERT INTO championships (
    id,
    name,
    organization,
    start_date,
    end_date,
    season_id,
    description,
    reference_img_url)
VALUES (
    '10000000-2026-4000-a000-000000000016',
    'Il Lombardia',
    'UCI World Tour',
    '2026-10-10T00:00:00Z',
    '2026-10-10T23:59:59Z',
    'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
    'Watch highlights of the 2026 Il Lombardia, a Monument and the final classic of the season in the hills of Lombardy.',
    'https://upload.wikimedia.org/wikipedia/commons/0/04/Magreglio_Santuario_di_Madonna_di_Ghisallo_Esterno_Facciata_1.jpg');

-- ============================================================
-- Events (one per one-day race, placeholder for stage races)
-- Stage race events will be added as the season progresses.
-- ============================================================
INSERT INTO events (
    id,
    name,
    date,
    championship_id)
VALUES
    -- One-day races
    (
        '20000000-2026-4000-a000-000000000001',
        'Omloop Het Nieuwsblad',
        '2026-02-28T00:00:00Z',
        '10000000-2026-4000-a000-000000000001'),
    (
        '20000000-2026-4000-a000-000000000002',
        'Strade Bianche',
        '2026-03-07T00:00:00Z',
        '10000000-2026-4000-a000-000000000002'),
    (
        '20000000-2026-4000-a000-000000000005',
        'Milano–San Remo',
        '2026-03-21T00:00:00Z',
        '10000000-2026-4000-a000-000000000005'),
    (
        '20000000-2026-4000-a000-000000000017',
        'E3 Saxo Classic',
        '2026-03-27T00:00:00Z',
        '10000000-2026-4000-a000-000000000017'),
    (
        '20000000-2026-4000-a000-000000000018',
        'Gent–Wevelgem',
        '2026-03-29T00:00:00Z',
        '10000000-2026-4000-a000-000000000018'),
    (
        '20000000-2026-4000-a000-000000000006',
        'Tour of Flanders',
        '2026-04-05T00:00:00Z',
        '10000000-2026-4000-a000-000000000006'),
    (
        '20000000-2026-4000-a000-000000000007',
        'Paris–Roubaix',
        '2026-04-12T00:00:00Z',
        '10000000-2026-4000-a000-000000000007'),
    (
        '20000000-2026-4000-a000-000000000008',
        'Amstel Gold Race',
        '2026-04-19T00:00:00Z',
        '10000000-2026-4000-a000-000000000008'),
    (
        '20000000-2026-4000-a000-000000000019',
        'La Flèche Wallonne',
        '2026-04-22T00:00:00Z',
        '10000000-2026-4000-a000-000000000019'),
    (
        '20000000-2026-4000-a000-000000000009',
        'Liège–Bastogne–Liège',
        '2026-04-26T00:00:00Z',
        '10000000-2026-4000-a000-000000000009'),
    (
        '20000000-2026-4000-a000-000000000022',
        'Grand Prix Cycliste de Québec',
        '2026-09-11T00:00:00Z',
        '10000000-2026-4000-a000-000000000022'),
    (
        '20000000-2026-4000-a000-000000000023',
        'Grand Prix Cycliste de Montréal',
        '2026-09-13T00:00:00Z',
        '10000000-2026-4000-a000-000000000023'),
    (
        '20000000-2026-4000-a000-000000000016',
        'Il Lombardia',
        '2026-10-10T00:00:00Z',
        '10000000-2026-4000-a000-000000000016');


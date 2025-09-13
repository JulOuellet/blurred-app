INSERT INTO sports (id, name, created_at, updated_at)
    VALUES ('e89d82d8-4bcc-423c-bbd0-18dbd0c1da01', 'Road Cycling', '2025-09-06T21:40:56.014764Z', '2025-09-06T21:40:56.014764Z');

INSERT INTO seasons (id, name, start_date, end_date, sport_id, created_at, updated_at)
    VALUES ('fb566050-e7ed-464c-9a19-b984b5dcdd34', '2025', '0001-01-01T00:00:00Z', '0001-01-01T00:00:00Z', 'e89d82d8-4bcc-423c-bbd0-18dbd0c1da01', '2025-09-06T21:41:32.198157Z', '2025-09-06T21:41:32.198157Z');

INSERT INTO championships (id, name, organization, start_date, end_date, season_id, description, reference_img_url, created_at, updated_at)
    VALUES ('4f37ff2c-d000-4486-88c6-3a059d4c96e9', 'Grand Prix cycliste de Québec', 'UCI', '2025-09-12T00:00:00Z', '2025-09-12T23:59:59Z', 'fb566050-e7ed-464c-9a19-b984b5dcdd34', 'Watch highlights of the Grand Prix cycliste de Québec taking place on September 12 in the roads of Quebec City.', 'https://images.unsplash.com/photo-1613061323515-60cf45de72d7', '2025-09-06T22:31:24.760113Z', '2025-09-06T22:31:24.760113Z'),
    ('4a965e4a-037d-4cb1-9ac1-499a8b5a93bc', 'Tour de France', 'UCI', '2025-07-05T00:00:00Z', '2025-07-27T23:59:59Z', 'fb566050-e7ed-464c-9a19-b984b5dcdd34', 'Watch highlights of all 21 stages of the 2025 Tour de France, running from July 5 to July 27.', 'https://images.unsplash.com/photo-1640890810035-aea65fdb1291', '2025-09-06T21:42:10.475667Z', '2025-09-06T21:42:10.475667Z');


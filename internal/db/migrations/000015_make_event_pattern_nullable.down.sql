UPDATE integrations SET event_pattern = '' WHERE event_pattern IS NULL;
ALTER TABLE integrations ALTER COLUMN event_pattern SET NOT NULL;

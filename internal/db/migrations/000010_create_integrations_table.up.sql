CREATE TABLE integrations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    youtube_channel_id varchar(255) NOT NULL,
    youtube_channel_name varchar(255),
    championship_id uuid NOT NULL,
    lang varchar(10) NOT NULL,
    source varchar(255),
    relevance_pattern text NOT NULL,
    event_pattern text NOT NULL,
    active boolean NOT NULL DEFAULT true,
    last_polled_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_championship FOREIGN KEY (championship_id) REFERENCES championships (id) ON DELETE CASCADE,
    CONSTRAINT uq_channel_championship UNIQUE (youtube_channel_id, championship_id)
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON integrations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

CREATE INDEX idx_integrations_championship_id ON integrations (championship_id);

CREATE INDEX idx_integrations_active ON integrations (active)
WHERE
    active = true;

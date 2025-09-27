CREATE TABLE events (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) NOT NULL,
    date timestamptz,
    championship_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_championship FOREIGN KEY (championship_id) REFERENCES championships (id) ON DELETE CASCADE
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON events
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

CREATE INDEX idx_events_championship_id ON events (championship_id);


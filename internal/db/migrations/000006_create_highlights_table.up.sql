CREATE TABLE highlights (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) NOT NULL,
    url text NOT NULL,
    youtube_id varchar(50),
    lang varchar(10) NOT NULL,
    media_type varchar(50) NOT NULL,
    source varchar(255),
    event_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON highlights
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

CREATE INDEX idx_highlights_event_id ON highlights (event_id);


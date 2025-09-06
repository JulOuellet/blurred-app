CREATE TABLE seasons (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) NOT NULL,
    start_date timestamptz,
    end_date timestamptz,
    sport_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_sport FOREIGN KEY (sport_id) REFERENCES sports (id) ON DELETE CASCADE
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON seasons
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

CREATE INDEX idx_seasons_sport_id ON seasons (sport_id);


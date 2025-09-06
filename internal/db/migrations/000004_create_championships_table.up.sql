CREATE TABLE championships (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name varchar(255) NOT NULL,
    organization varchar(255),
    start_date timestamptz,
    end_date timestamptz,
    season_id uuid NOT NULL,
    description text,
    reference_img_url text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_season FOREIGN KEY (season_id) REFERENCES seasons (id) ON DELETE CASCADE
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON championships
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

CREATE INDEX idx_championships_season_id ON championships (season_id);


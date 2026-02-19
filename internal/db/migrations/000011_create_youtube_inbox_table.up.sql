CREATE TABLE youtube_inbox (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    integration_id uuid NOT NULL,
    youtube_video_id varchar(50) NOT NULL,
    video_title text NOT NULL,
    published_at timestamptz,
    status varchar(20) NOT NULL DEFAULT 'pending',
    failure_reason text,
    retry_count integer NOT NULL DEFAULT 0,
    processed_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_integration FOREIGN KEY (integration_id) REFERENCES integrations (id) ON DELETE CASCADE,
    CONSTRAINT uq_youtube_video UNIQUE (youtube_video_id)
);

CREATE INDEX idx_inbox_status ON youtube_inbox (status)
WHERE
    status IN ('pending', 'failed');

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON youtube_inbox
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

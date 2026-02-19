package youtube

import "time"

type FeedEntry struct {
	VideoID     string
	Title       string
	PublishedAt time.Time
	ChannelName string
}

type VideoDetails struct {
	Title           string
	DurationSeconds int
	Description     string
}

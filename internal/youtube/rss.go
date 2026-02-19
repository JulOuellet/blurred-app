package youtube

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const rssFeedURL = "https://www.youtube.com/feeds/videos.xml?channel_id=%s"

type atomFeed struct {
	XMLName xml.Name   `xml:"feed"`
	Title   string     `xml:"title"`
	Entries []atomEntry `xml:"entry"`
}

type atomEntry struct {
	VideoID   string `xml:"videoId"`
	Title     string `xml:"title"`
	Published string `xml:"published"`
	Author    struct {
		Name string `xml:"name"`
	} `xml:"author"`
}

func FetchFeed(channelID string) ([]FeedEntry, error) {
	url := fmt.Sprintf(rssFeedURL, channelID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RSS feed returned status %d", resp.StatusCode)
	}

	var feed atomFeed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	entries := make([]FeedEntry, 0, len(feed.Entries))
	for _, e := range feed.Entries {
		publishedAt, _ := time.Parse(time.RFC3339, e.Published)
		entries = append(entries, FeedEntry{
			VideoID:     e.VideoID,
			Title:       e.Title,
			PublishedAt: publishedAt,
			ChannelName: e.Author.Name,
		})
	}

	return entries, nil
}

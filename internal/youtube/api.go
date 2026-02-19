package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

const videosAPIURL = "https://www.googleapis.com/youtube/v3/videos"

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

type videosAPIResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
		ContentDetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
	} `json:"items"`
}

func (c *Client) GetVideoDetails(videoID string) (*VideoDetails, error) {
	url := fmt.Sprintf("%s?part=snippet,contentDetails&id=%s&key=%s", videosAPIURL, videoID, c.apiKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call YouTube API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("YouTube API returned status %d", resp.StatusCode)
	}

	var apiResp videosAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode YouTube API response: %w", err)
	}

	if len(apiResp.Items) == 0 {
		return nil, fmt.Errorf("video not found: %s", videoID)
	}

	item := apiResp.Items[0]
	duration := parseISO8601Duration(item.ContentDetails.Duration)

	return &VideoDetails{
		Title:           item.Snippet.Title,
		DurationSeconds: duration,
		Description:     item.Snippet.Description,
	}, nil
}

var iso8601Re = regexp.MustCompile(`PT(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)

func parseISO8601Duration(d string) int {
	matches := iso8601Re.FindStringSubmatch(d)
	if matches == nil {
		return 0
	}

	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])
	seconds, _ := strconv.Atoi(matches[3])

	return hours*3600 + minutes*60 + seconds
}

package layouts

import (
	"os"
	"strings"
)

// publicBaseURL returns the canonical origin of the deployed site (e.g.
// "https://blurred.app"), used to build the absolute URLs that OpenGraph
// scrapers require. Empty when PUBLIC_BASE_URL is unset.
func publicBaseURL() string {
	return strings.TrimSuffix(os.Getenv("PUBLIC_BASE_URL"), "/")
}

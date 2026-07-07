package layouts

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"
)

var assetVersion = func() string {
	if sha := os.Getenv("RAILWAY_GIT_COMMIT_SHA"); len(sha) >= 8 {
		return sha[:8]
	}
	return strconv.FormatInt(time.Now().Unix(), 10)
}()

// AssetVersion returns a value that changes on every deploy, appended to
// static asset URLs so CDN and browser caches (Cloudflare caches CSS for
// hours) can never serve a stylesheet from a previous release.
func AssetVersion() string {
	return assetVersion
}

// PublicBaseURL returns the canonical origin of the deployed site (e.g.
// "https://blurred.watch"), used to build the absolute URLs that OpenGraph
// scrapers and search engines require. Empty when PUBLIC_BASE_URL is unset.
func PublicBaseURL() string {
	return strings.TrimSuffix(os.Getenv("PUBLIC_BASE_URL"), "/")
}

// PageTitle turns a page's name into a full document title that carries the
// brand and, on the home page, the search keywords the site should rank for.
func PageTitle(title string) string {
	if title == "" || title == "Blurred" {
		return "Blurred — Spoiler-Free Sports Highlights"
	}
	if strings.Contains(title, "Blurred") {
		return title
	}
	return title + " | Blurred"
}

type contextKey string

const requestPathKey contextKey = "requestPath"

// WithRequestPath stores the current request path so templates can build
// canonical URLs without every handler having to pass it down.
func WithRequestPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, requestPathKey, path)
}

func canonicalURL(ctx context.Context) string {
	base := PublicBaseURL()
	if base == "" {
		return ""
	}
	path, _ := ctx.Value(requestPathKey).(string)
	if path == "" || path == "/" {
		return base + "/"
	}
	return base + path
}

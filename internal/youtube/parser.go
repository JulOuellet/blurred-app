package youtube

import (
	"regexp"
	"strconv"
)

func MatchesRelevancePattern(title string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(title)
}

func ExtractEventNumber(title string, pattern *regexp.Regexp) (int, bool) {
	matches := pattern.FindStringSubmatch(title)
	if len(matches) < 2 {
		return 0, false
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, false
	}

	return num, true
}

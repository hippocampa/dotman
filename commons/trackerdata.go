package commons

import (
	"regexp"
)

type TrackerData struct {
	Source string       `json:"source"`
	Stored string       `json:"stored"`
	Status StatusString `json:"status"`
}

var tracker TrackerData

func normalizeHomePath(path *string) {
	re := regexp.MustCompile(`^/home/[^/]+/(.*)$`)
	if re.MatchString(*path) {
		matches := re.FindStringSubmatch(*path)
		if len(matches) > 1 {
			*path = "$HOME/" + matches[1]
		}
	}
}

func NewTrackerData(src string,
	dest string,
	status StatusString) {
	normalizeHomePath(&src)
	normalizeHomePath(&dest)
	tracker = TrackerData{Source: src,
		Stored: dest, Status: status}
}

func SaveTrackerData() {
	// TODO
}

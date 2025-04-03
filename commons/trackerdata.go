package commons

type TrackerData struct {
	Source string       `json:"source"`
	Stored string       `json:"stored"`
	Status StatusString `json:"status"`
}

func NewTrackerData(src string, dest string, status StatusString) *TrackerData {
	return &TrackerData{
		Source: src,
		Stored: dest,
		Status: status,
	}
}

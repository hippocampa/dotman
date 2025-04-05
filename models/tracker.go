package models

var (
	TrackerDataList []TrackerData
	Data            TrackerData
)

type TrackerData struct {
	Source string `json:"source"`
	Stored string `json:"stored"`
	Status string `json:"status"`
}

func NewTrackerData(src, stored, status string) *TrackerData {
	return &TrackerData{
		Source: src,
		Stored: stored,
		Status: status,
	}
}

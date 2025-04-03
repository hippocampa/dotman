package commons

import (
	"encoding/json"
	"os"
	"regexp"
)

type TrackerData struct {
	Source string       `json:"source"`
	Stored string       `json:"stored"`
	Status StatusString `json:"status"`
}

var tracker TrackerData
var trackerDataList []TrackerData

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

func ReadJSON(path string) error {
	indexjsonloc := path + "/index.json"
	fileData, err := os.ReadFile(indexjsonloc)
	if err != nil {
		// If file doesn't exist, initialize empty trackerDataList
		if os.IsNotExist(err) {
			trackerDataList = []TrackerData{}
			return nil
		}
		return err
	}
	if len(fileData) == 0 {
		trackerDataList = []TrackerData{}
		return nil
	}
	err = json.Unmarshal(fileData, &trackerDataList)
	if err != nil {
		return err
	}
	return nil
}

func SaveTrackerData() error {
	homepath, err := GetDotFilesDir()
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	err = os.MkdirAll(homepath, 0755)
	if err != nil {
		return err
	}

	// Read existing data first
	err = ReadJSON(homepath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Append the new tracker data
	trackerDataList = append(trackerDataList, tracker)

	// Save updated list
	indexjsonloc := homepath + "/index.json"
	jsonData, err := json.MarshalIndent(trackerDataList, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(indexjsonloc, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

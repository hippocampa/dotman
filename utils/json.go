package utils

import (
	"encoding/json"
	"os"

	"github.com/hippocampa/dotshadow/models"
)

func ReadJSON(path string) error {
	fileData, err := os.ReadFile(path)
	if err != nil {
		// If file doesn't exist, initialize empty trackerDataList
		if os.IsNotExist(err) {
			models.TrackerDataList = []models.TrackerData{}
			return nil
		}
		return err
	}
	if len(fileData) == 0 {
		models.TrackerDataList = []models.TrackerData{}
		return nil
	}
	err = json.Unmarshal(fileData, &models.TrackerDataList)
	if err != nil {
		return err
	}
	return nil
}

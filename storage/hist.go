package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hippocampa/dotshadow/models"
	"github.com/hippocampa/dotshadow/utils"
)

func UpdateHist(src, dest, status string) error {
	data := models.NewTrackerData(utils.NormalizeHomePath(src), utils.NormalizeHomePath(dest), "linked")

	dotfilesRoot, err := utils.GetDotfilesRoot()
	if err != nil {
		return err
	}
	if err = utils.ReadJSON(dotfilesRoot + "/index.json"); err != nil {
		return err
	}
	models.TrackerDataList = append(models.TrackerDataList, *data)
	jsonData, err := json.MarshalIndent(models.TrackerDataList, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(dotfilesRoot+"/index.json", jsonData, 0o644); err != nil {
		return fmt.Errorf("error saving index.json: %w", err)
	}

	return nil
}

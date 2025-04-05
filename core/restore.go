package core

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hippocampa/dotshadow/models"
	"github.com/hippocampa/dotshadow/storage"
	"github.com/hippocampa/dotshadow/utils"
	"github.com/urfave/cli/v2"
)

func Restore(context *cli.Context) error {
	if context.Bool("all") {
		if err := restoreAll(); err != nil {
			return err
		}
	} else if context.Args().Len() == 0 {
		return fmt.Errorf("missing <file>: either provide a file or use --all")
	}
	return nil
}

func restoreAll() error {
	utils.PrintAction.Println(" RESTORING DOTFILES ")

	dotfileRoot, err := utils.GetDotfilesRoot()
	if err != nil {
		return err
	}
	histFile, err := os.ReadFile(dotfileRoot + "/index.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(histFile, &models.TrackerDataList); err != nil {
		return err
	}

	for _, data := range models.TrackerDataList {
		if err := storage.Link(utils.DenormalizeHomePath(data.Stored),
			utils.DenormalizeHomePath(data.Source)); err != nil {
			return err
		}
	}

	return nil
}

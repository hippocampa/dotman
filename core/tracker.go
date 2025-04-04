package core

import (
	"fmt"
	"path/filepath"

	"github.com/hippocampa/dotshadow/utils"
	"github.com/urfave/cli/v2"
)

func Track(context *cli.Context) error {
	source := context.Args().Get(0)
	fmt.Println(filepath.Abs(source))
	_, err := utils.GetDotfilesRoot()
	if err != nil {
		return fmt.Errorf("GetDotfilesRoot: %w", err)
	}

	return nil
}

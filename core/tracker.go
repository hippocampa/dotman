package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hippocampa/dotshadow/storage"
	"github.com/hippocampa/dotshadow/utils"
	"github.com/urfave/cli/v2"
)

func Track(context *cli.Context) error {
	source := context.Args().Get(0)
	dotfilesRoot, err := utils.GetDotfilesRoot()
	if err != nil {
		return fmt.Errorf("GetDotfilesRoot: %w", err)
	}

	sourceRelative, err := utils.GetHomeRelativePath(source)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(filepath.Join(dotfilesRoot, sourceRelative)), 0o755); err != nil {
		return fmt.Errorf("error creating folder: %w", err)
	}

	if err := storage.Copy(source,
		filepath.Join(dotfilesRoot, sourceRelative)); err != nil {
		return fmt.Errorf("error copying files: %w", err)
	}

	return nil
}

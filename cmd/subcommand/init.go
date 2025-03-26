package subcommand

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hippocampa/dotshadow/commons"
	"github.com/urfave/cli/v2"
)

func InitCmd() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  "Create dotfiles folder",
		Action: InitAction,
	}
}

func InitAction(context *cli.Context) error {
	dotfilesDir, err := commons.GetDotFilesDir()
	if err != nil {
		return err
	}
	// create .dotfiles in homedirectory
	if err := os.Mkdir(dotfilesDir, 0o755); err != nil {
		return err
	}

	// crate index.json
	if _, err := os.Create(filepath.Join(dotfilesDir, "index.json")); err != nil {
		return err
	}
	fmt.Printf("Intialized:\n.dotfiles\t%v\nindex.json\t%v\n", dotfilesDir, filepath.Join(dotfilesDir, "index.json"))

	return nil
}

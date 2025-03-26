package subcommand

import (
	"os"
	"os/user"
	"path/filepath"

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
	// get current user
	currUser, err := user.Current()
	if err != nil {
		return err
	}
	userHomedir := currUser.HomeDir

	if err := os.Mkdir(filepath.Join(userHomedir, "dotfiles"), 0o755); err != nil {
		return nil
	}
	return nil
}

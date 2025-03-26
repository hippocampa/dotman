package subcommand

import (
	"github.com/urfave/cli/v2"
)

func AddCmd() *cli.Command {
	return &cli.Command{
		Name:   "add",
		Usage:  "Add dotfiles",
		Action: AddAction,
	}
}

func AddAction(context *cli.Context) error {
	/*Usage
		dotshadow add <dir> to add new file <dir> to the dotfiles directory
	  return error if the directory is already exist
	*/

	return nil
}

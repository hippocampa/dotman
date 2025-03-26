package subcommand

import (
	"github.com/urfave/cli/v2"
)

func SyncCmd() *cli.Command {
	return &cli.Command{
		Name:   "sync",
		Usage:  "Synchronize dotfiles",
		Action: SyncAction,
	}
}

func SyncAction(context *cli.Context) error {
	/*Usage
	dotshadow sync
	Synchronize all the dotfiles to reflect recent changes
	*/
	return nil
}

package subcommand

import (
	"github.com/urfave/cli/v2"
)

func RestoreCmd() *cli.Command {
	return &cli.Command{
		Name:      "restore",
		Usage:     "Restore one or more tracked dotfiles",
		ArgsUsage: "[file]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Restore all tracked files",
			},
			&cli.BoolFlag{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "Copy the file instead of symlinking it",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force overwrite if destination file already exists",
			},
		},
	}
}

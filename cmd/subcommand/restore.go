package subcommand

import (
	"github.com/hippocampa/dotshadow/core"
	"github.com/urfave/cli/v2"
)

func RestoreCmd() *cli.Command {
	return &cli.Command{
		Name:      "restore",
		Usage:     "Restore one or all tracked dotfiles",
		ArgsUsage: "<file>",
		Action:    core.Restore,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Restore all tracked files",
			},
		},
	}
}

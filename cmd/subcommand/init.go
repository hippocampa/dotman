package subcommand

import (
	"github.com/hippocampa/dotshadow/core"
	"github.com/urfave/cli/v2"
)

func InitCmd() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  "Create dotfiles folder",
		Action: core.Init,
	}
}

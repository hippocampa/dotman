package subcommand

import (
	"github.com/hippocampa/dotshadow/core"
	"github.com/urfave/cli/v2"
)

func TrackCmd() *cli.Command {
	return &cli.Command{
		Name:      "track",
		Usage:     "Track a single dotfile",
		ArgsUsage: "<file>",
		Action:    core.Track,
	}
}

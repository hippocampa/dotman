package cmd

import "github.com/urfave/cli/v2"

type Cmd struct {
	CliApp *cli.App
}

func New() *Cmd {
	return &Cmd{CliApp: &cli.App{
		Name:     "dotman",
		Commands: []*cli.Command{},
	}}
}

func (c *Cmd) AddSubCmd(subcmd *cli.Command) {
	c.CliApp.Commands = append(c.CliApp.Commands, subcmd)
}

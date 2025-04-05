package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/hippocampa/dotshadow/utils"
	"github.com/urfave/cli/v2"
)

func Init(context *cli.Context) error {
	banner := `
  _____        _                         
 |  __ \      | |                        
 | |  | | ___ | |_ _ __ ___   __ _ _ __  
 | |  | |/ _ \| __| '_ ` + "`" + ` _ \ / _` + "`" + ` | '_ \ 
 | |__| | (_) | |_| | | | | | (_| | | | |
 |_____/ \___/ \__|_| |_| |_|\__,_|_| |_|
                                         
                                         `
	fmt.Println(banner)
	dotfilesRoot, err := utils.GetDotfilesRoot()
	if err != nil {
		return fmt.Errorf("GetDotfilesRoot: %w", err)
	}
	// create .dotfiles in homedirectory
	if err := os.Mkdir(dotfilesRoot, 0o755); err != nil {
		return err
	}

	// crate index.json
	if _, err := os.Create(filepath.Join(dotfilesRoot, "index.json")); err != nil {
		return err
	}
	utils.PrintAction.Println(" INITIALIZING ")
	color.Green(" .dotfiles\t%v\n index.json\t%v\n",
		dotfilesRoot, filepath.Join(dotfilesRoot, "index.json"))

	return nil
}

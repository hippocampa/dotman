package subcommand

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hippocampa/dotshadow/commons"
	"github.com/urfave/cli/v2"
)

func TrackCmd() *cli.Command {
	return &cli.Command{
		Name:      "track",
		Usage:     "Track files",
		Action:    trackAction,
		ArgsUsage: "<FILE PATH>",
	}
}

func trackAction(context *cli.Context) error {
	dotfilesDir, err := commons.GetDotFilesDir()
	if err != nil {
		return err
	}
	source := context.Args().Get(0)
	fmt.Println("Tracking: ", source)

	if err := track(source, dotfilesDir); err != nil {
		return err
	}

	return nil
}

func track(source string, destination string) error {
	// TODO: Update index.js
	userHomeDir, err := commons.GetHomeDir()
	if err != nil {
		return err
	}
	relativePath := filepath.Join(strings.TrimPrefix(source, userHomeDir))
	destDir := filepath.Join(destination, filepath.Dir(relativePath))
	destFull := filepath.Join(destination, relativePath)
	folderErr := os.MkdirAll(destDir, 0o755)

	if folderErr != nil {
		return folderErr
	}

	// if err := copyFile(source, destFull); err != nil {
	// 	return err
	// }

	// if err := createSymLink(source, destFull); err != nil {
	// 	return err
	// }
	data := commons.NewTrackerData(source, destFull, commons.LINKED)
	fmt.Println(data)

	return nil
}

// TODO: Do i need to check if the symlink is exist?
func createSymLink(source, dest string) error {
	if _, err := os.Lstat(dest); err == nil {
		// File/symlink exists, remove it
		err := os.Remove(dest)
		if err != nil {
			log.Fatal("Failed to remove existing file:", err)
		}
	}
	if err := os.Symlink(source, dest); err != nil {
		return err
	}
	fmt.Printf("Symlink Created:\n%v to %v\n", source, dest)
	return nil
}

func copyFile(source string, destination string) error {
	// read file from source
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}
	fmt.Printf("Copying %v to %v\n", source, destination)

	return nil
}

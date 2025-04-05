package storage

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

func Copy(source string, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	color.Green("COPYING:")
	color.Magenta("%v to %v\n", source, destination)

	return nil
}

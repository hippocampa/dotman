package storage

import (
	"fmt"
	"log"
	"os"
)

func Link(source string, dest string) error {
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

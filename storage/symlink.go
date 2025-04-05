package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func Link(target string, linkname string) error {
	target = filepath.Clean(target)
	linkname = filepath.Clean(linkname)

	// Prevent self-linking
	if target == linkname {
		return fmt.Errorf("refusing to link file to itself: %s", target)
	}

	if fi, err := os.Lstat(linkname); err == nil {
		if fi.Mode()&os.ModeSymlink != 0 {
			currentTarget, err := os.Readlink(linkname)
			if err == nil && currentTarget != target {
				color.Yellow(" [WARN] Replacing existing symlink at %s → %s", linkname, currentTarget)
			}
		} else {
			color.Yellow(" [WARN]️ Overwriting existing file at %s", linkname)
		}

		if err := os.Remove(linkname); err != nil {
			return fmt.Errorf("failed to remove existing file/symlink: %w", err)
		}
	}

	if err := os.Symlink(target, linkname); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	// Success log

	color.New(color.FgGreen).Print(" [SYMLINKED] ")
	color.Green("%s → %s\n", linkname, target)

	return nil
}

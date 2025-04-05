package storage

import "fmt"

func Commit(source, dest string) error {
	if err := Copy(source, dest); err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	if err := Link(source, dest); err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	// update json
	return nil
}

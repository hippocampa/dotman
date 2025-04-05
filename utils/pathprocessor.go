package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GetHomeRelativePath(path string) (string, error) {
	dotfilesRoot, err := GetUserHomeDir()
	if err != nil {
		return "", fmt.Errorf("GetDotfilesRoot: %w", err)
	}
	relativePath := filepath.Join(strings.TrimPrefix(path, dotfilesRoot))
	return relativePath, nil
}

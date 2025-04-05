package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
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

func NormalizeHomePath(path string) string {
	re := regexp.MustCompile(`^/home/[^/]+/(.*)$`)
	if re.MatchString(path) {
		matches := re.FindStringSubmatch(path)
		if len(matches) > 1 {
			path = "$HOME/" + matches[1]
		}
	}
	return path
}

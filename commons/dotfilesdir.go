package commons

import (
	"os/user"
	"path/filepath"
)

func GetHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homedir := currentUser.HomeDir
	return homedir, nil
}

func GetDotFilesDir() (string, error) {
	homedir, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".dotfiles"), nil
}

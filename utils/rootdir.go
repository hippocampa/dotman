package utils

import (
	"os/user"
	"path/filepath"
)

func GetUserHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homedir := currentUser.HomeDir
	return homedir, nil

}

func GetDotfilesRoot() (string, error) {
	homedir, err := GetUserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".dotfiles"), nil

}

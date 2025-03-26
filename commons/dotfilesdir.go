package commons

import (
	"os/user"
	"path/filepath"
)

func GetDotFilesDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homedir := currentUser.HomeDir
	return filepath.Join(homedir, ".dotfiles"), nil
}

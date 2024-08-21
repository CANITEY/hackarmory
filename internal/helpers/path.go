package helpers

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"

)

func CheckPath(p string) bool {
	pathVar := os.Getenv("PATH")
	pathSlice := strings.Split(pathVar, ":")
	return slices.Contains(pathSlice, p)
}

func AddPath(p string) (bool, error) {
	if ok := CheckPath(p); !ok {
		return true, nil
	}
	shellName := os.Getenv("SHELL")

	switch {
	case strings.Contains(shellName, "fish"):
		exec := executer{}
		if err := exec.Execute("fish_add_path /opt/mycoolthing/bin", ""); err != nil {
			return false, err
		}

		return true, nil
	case strings.Contains(shellName, "zsh"):
		if err := writePath(".zshrc", p); err != nil {
			return false, err
		}
	case strings.Contains(shellName, "bash"):
		if err := writePath(".bashrc", p); err != nil {
			return false, err
		}
	}

	return true, nil
}

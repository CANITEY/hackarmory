package installs

import (
	"os"
	"path"
)

type Commands map[string]func() error

var commands = Commands{
	"subfinder": subfinder,
	"amass": amass,
	"gobuster": gobuster,
}

// creates the directory that the tools will reside in
func CreateToolsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	toolsPath := path.Join(home, "tools/")
	if err := os.MkdirAll(toolsPath, os.ModePerm); err != nil {
		return "", nil
	}


	return toolsPath, nil
}

func Install(queue []string) []error {

	return nil
}

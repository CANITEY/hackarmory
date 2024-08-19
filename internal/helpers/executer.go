package helpers

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

)

type executer struct {
	rootPath, shortPath string
}

func NewExecuter(absPath string) *executer {
	shortPath := path.Join(absPath, "path")
	return &executer{
		rootPath: absPath,
		shortPath: shortPath,
	}
}

func (e *executer) Execute(command, relPath string) (error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path.Join(e.rootPath, relPath)

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (e *executer) CreateSymLink(toolsPath, shortcutName string) error {
	toolPath, _ := filepath.Abs(toolsPath)
	shortPath, _ := filepath.Abs(path.Join(e.shortPath, shortcutName))
	if err := os.Symlink(toolPath, shortPath); err != nil {
		return err
	}

	return nil
}

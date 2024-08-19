package helpers

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

type executer struct {
	rootPath, ShortPath string
}

func NewExecuter(absPath string) *executer {
	shortPath := path.Join(absPath, "path")
	return &executer{
		rootPath: absPath,
		ShortPath: shortPath,
	}
}

func (e *executer) Execute(command, relPath string) (error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = path.Join(e.rootPath, relPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

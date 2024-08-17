package helpers

import (
	"os/exec"
	"path"
	"strings"
)

type executer string

func NewExecuter(absPath string) *executer {
	return (*executer)(&absPath)
}

func (e *executer) Execute(command, relPath string) (error) {
	cmd := exec.Cmd{}
	cmd.Dir = path.Join(string(*e), relPath)
	cmd.Args = strings.Split("git clone https://github.com/aboul3la/Sublist3r.git", " ")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

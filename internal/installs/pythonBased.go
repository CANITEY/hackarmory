package installs

import (
	"os/exec"
	"path"
	"strings"
)

func arjun() (error) {
	cmd := exec.Cmd{}
	toolsDir, err := GetToolsDir()
	if err != nil {
		return err
	}
	cmd.Dir  = toolsDir
	cmd.Args = strings.Split("git clone https://github.com/s0md3v/Arjun", " ")
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd.Dir  = path.Join(toolsDir, "Arjun")
	cmd.Args = strings.Split("python3 setup.py install", " ")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

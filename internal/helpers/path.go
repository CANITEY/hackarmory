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
	// TODO: add the function that adds path to path variable
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
		// TODO: Make the writer write between the # START and # END
		home, _ := os.UserHomeDir()
		zshrc := path.Join(home, ".zshrc")
		file, err := os.OpenFile(zshrc, os.O_CREATE | os.O_RDWR, 0644)
		if err != nil {
			return false, err
		}
		data := fmt.Sprintf("\n\n# START of hackarmory modification \nexport PATH=$PATH:%v\n\n# END of hackarmory\n", p)
		file.WriteString(data)
	case strings.Contains(shellName, "bash"):
		home, _ := os.UserHomeDir()
		zshrc := path.Join(home, ".bashrc")
		file, err := os.OpenFile(zshrc, os.O_CREATE | os.O_RDWR, 0644)
		if err != nil {
			return false, err
		}
		data := fmt.Sprintf("\n\n# START of hackarmory addition \nexport PATH=$PATH:%v\n\n# END of hackarmory\n", p)
		file.WriteString(data)
	}

	return true, nil
}

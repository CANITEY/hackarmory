package helpers

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
	"bufio"
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


func writePath(conFile, p string) error {
	home, _ := os.UserHomeDir()
	rcLocation := path.Join(home, conFile)
	// READING FILE
	file, err := os.OpenFile(rcLocation, os.O_CREATE | os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// READ FILE INTO BUFFER
	buf := strings.Builder{}

	// SIGNAL TO INSURE THAT THE LINE IS WRITTEN
	ok := false


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// MODIFY OUR CONFIG LINES
		if scanner.Text() == "# START of hackarmory modification" {
			ok = true
			buf.WriteString(fmt.Sprintf("%v\n", scanner.Text()))
			scanner.Scan()
			data := fmt.Sprintf("\nexport PATH=$PATH:%v\n", p)
			buf.WriteString(data)
			continue
		}
		buf.WriteString(fmt.Sprintf("%v\n", scanner.Text()))
	}

	if !ok {
		data := fmt.Sprintf("# START of hackarmory modification\nexport PATH=$PATH:%v\n#END of hackarmory modification\n", p)
		buf.WriteString(data)
	}

	// WRITING THE FILE
	file, err = os.OpenFile(rcLocation, os.O_TRUNC | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(buf.String()); err != nil {
		return err
	}
	
	return nil
}

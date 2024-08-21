package helpers

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func WritePath(conFile, p string) error {
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
	file, err = os.OpenFile(conFile, os.O_TRUNC | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(buf.String()); err != nil {
		return err
	}
	
	return nil
}

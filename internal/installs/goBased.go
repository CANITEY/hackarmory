package installs

import (
	"os/exec"
	"strings"
)

func subfinder() error {
	cmd := exec.Command("go", strings.Split("install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func amass() error {
	cmd := exec.Command("go", strings.Split("install -v github.com/owasp-amass/amass/v4/...@master", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func gobuster() error {
	cmd := exec.Command("go", strings.Split("install github.com/OJ/gobuster/v3@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

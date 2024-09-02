package installs

import (
	"os/exec"
	"strings"

)
func katana() error {
	cmd := exec.Command("go", strings.Split("install github.com/projectdiscovery/katana/cmd/katana@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
func hakrawler() error {
	cmd := exec.Command("go", strings.Split("install github.com/hakluke/hakrawler@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func gospider() error {
	cmd := exec.Command("go", strings.Split("install github.com/jaeles-project/gospider@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func gowitness() error {
	cmd := exec.Command("go", strings.Split("install github.com/sensepost/gowitness@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

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


func nabuu() error {
	cmd := exec.Command("go", strings.Split("install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func nuclei() error {
	cmd := exec.Command("go", strings.Split("install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func gxss() error {
	cmd := exec.Command("go", strings.Split("install github.com/KathanP19/Gxss@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func assetfinder() error {
	cmd := exec.Command("go", strings.Split("install github.com/tomnomnom/assetfinder@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func httpx() error {
	cmd := exec.Command("go", strings.Split("install -v github.com/projectdiscovery/httpx/cmd/httpx@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func httprobe() error {
	cmd := exec.Command("go", strings.Split("install github.com/tomnomnom/httprobe@latest", " ")...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

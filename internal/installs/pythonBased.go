package installs

import (
	"path"

	"github.com/CANITEY/hackarmory/internal/helpers"
)

func arjun() error {
	toolsDir, err := GetToolsDir()
	if err != nil {
		return err
	}
	exec := helpers.NewExecuter(toolsDir)
	err = exec.Execute("git clone https://github.com/s0md3v/Arjun", "")
	if err != nil {
		return err
	}

	err = exec.Execute("python3 setup.py install", "Arjun/")
	if err != nil {
		return err
	}
	return nil
}

func sublister() error {
	toolsDir, err := GetToolsDir()
	if err != nil {
		return err
	}

	exec := helpers.NewExecuter(toolsDir)
	err = exec.Execute("git clone https://github.com/aboul3la/Sublist3r.git", "")
	if  err != nil {
		return err
	}
	err = exec.Execute("pip install -r requirements.txt --break-system-packages", "Sublist3r/")
	if  err != nil {
		return err
	}
	toolPath := path.Join(toolsDir, "Sublist3r", "sublist3r.py")
	return exec.CreateSymLink(toolPath, "sublist3r")
}

func dirsearch() error {
	toolsDir, err := GetToolsDir()
	if err != nil {
		return err
	}

	exec := helpers.NewExecuter(toolsDir)
	err = exec.Execute("git clone https://github.com/maurosoria/dirsearch.git --depth 1", "")
	if err != nil {
		return err
	}
	err = exec.Execute("pip install -r requirements.txt --break-system-packages", "dirsearch/")
	if  err != nil {
		return err
	}
	toolPath := path.Join(toolsDir, "dirsearch", "dirsearch.py")
	return exec.CreateSymLink(toolPath, "dirsearch")
}

package installs

import (
	"os"
	"path"
	"path/filepath"

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

func Sublister() error {
	toolsDir, err := GetToolsDir()
	if err != nil {
		return err
	}

	exec := helpers.NewExecuter(toolsDir)
	err = exec.Execute("git clone https://github.com/aboul3la/Sublist3r.git", "")
	if  err != nil {
		return err
	}
	err = exec.Execute("pip install -r requirements.txt --break-system-packages", "Sublist3r")
	if  err != nil {
		return err
	}

	// create a shortcat on path
	toolPath, _ := filepath.Abs(path.Join(toolsDir, "Sublist3r", "sublist3r.py"))
	shortPath, _ := filepath.Abs(path.Join(exec.ShortPath, "sublist3r"))
	if err := os.Symlink(toolPath, shortPath); err != nil {
		return err
	}

	return nil
}

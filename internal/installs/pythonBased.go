package installs

import (
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
		return err
	}

	cmd.Dir  = path.Join(toolsDir, "Arjun")
	cmd.Args = strings.Split("python3 setup.py install", " ")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

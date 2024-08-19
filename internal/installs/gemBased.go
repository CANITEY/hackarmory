package installs

import (
	"os"

	"github.com/CANITEY/hackarmory/internal/helpers"
)

func wpscan() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	exec := helpers.NewExecuter(home)
	err = exec.Execute("gem install wpscan", "")
	if err != nil {
		return nil
	}

	// TODO: add the gem path to your path variable


	return nil
}

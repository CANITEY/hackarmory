package checks

import (
	"os/exec"
)

type msg struct {
	Pass string
	Err error
}

func CheckDependency(dep string) (string, error) {
	_, err := exec.Command("which", dep).Output()
	if err != nil {
		return dep, err
	}

	return dep, nil
}

func CheckDependencies(deps ...string) (map[string]msg) {
	var output = make(map[string]msg)
	for _, dep := range deps {
		if out, err := CheckDependency(dep); err != nil {
			output[dep] = msg{
				Err: err,
			}
		} else {
			output[dep] = msg{
				Pass: out,
			}
		}
	}
	return output
}

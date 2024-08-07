package checks

import (
	"os/exec"
)

type msg struct {
	Pass string
	Err error
}

func CheckDependency(dep string) (string, error) {
	output, err := exec.Command(dep, "--version").Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
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

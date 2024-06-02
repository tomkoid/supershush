package tools

import "os/exec"

func RunCommand(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}

package tools

import "os/exec"

func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

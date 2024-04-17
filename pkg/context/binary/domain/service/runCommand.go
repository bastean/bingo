package service

import (
	"os/exec"
	"strings"
)

func RunCommand(args ...string) {
	if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
		FailOnError(err, strings.Join(args, " "))
	}
}

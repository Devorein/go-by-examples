package examples

import (
	"os"
	"os/exec"
	"syscall"
)

func ExecingProcesses() {
	bin, err := exec.LookPath("dir")
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(bin, args, env)

	if execErr != nil {
		panic(execErr)
	}
}

package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func gccExecution(files ...string) error {
	fmt.Println("GCC: ", files)
	var arg []string
	arg = append(arg, "-o", "app")
	arg = append(arg, files...)
	cmd := exec.Command("gcc", arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Cannot compile by GCC.\nFiles = %v\nErrors = %v\nStderr = %v", files, err, stderr.String())
	}
	return nil
}

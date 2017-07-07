package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func gccExecution(file string) error {
	appName := convertFromSourceToAppName(file)
	name := file
	args := [][]string{
		[]string{"-o", appName, name},
		[]string{"-O", "-o", appName, name, "-lm"},
	}
	for _, arg := range args {
		cmd := exec.Command("gcc", arg...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("=== MISTAKE IN GCC ===")
			fmt.Printf("Cannot compile by gcc file : %v\n", name)
			fmt.Printf("Arguments for compile      : %v\n", arg)
			fmt.Printf("App name                   : %v\n", appName)
			fmt.Printf("Error                      : %v\n", stderr.String())
			continue
		}
		return nil
	}
	return fmt.Errorf("Cannot compile by GCC")
}

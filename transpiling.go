package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func c2goTranspiling(file string, goFile string) error {
	cmd := exec.Command(c2go, transpile, "-o", goFile, file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("=== MISTAKE ===\n")
		s := fmt.Sprintf("Command : c2go %v -o %v %v\n", transpile, goFile, file)
		s += fmt.Sprintf("Cannot compile by c2go file with name : %v\nGo name : %v\nError: %v\n\n", file, goFile, stderr.String())
		fmt.Println(s)
		return fmt.Errorf(stderr.String())
	}
	return nil
}

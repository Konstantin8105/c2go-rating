package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func c2goTranspiling(file string) error {
	cmd := exec.Command("c2go", "transpile", "-o", "/tmp/1.go", file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("c2go : %v\n%v\n%v", err, out.String(), stderr.String())
	}
	return nil
}

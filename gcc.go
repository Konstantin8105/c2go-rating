package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func gccExecution(appName string, files ...string) error {
	argTemplates := [][]string{
		[]string{"-o", appName, "<files>"},
		[]string{"-O", "-o", appName, "<files>", "-lm"},
		[]string{"-o", appName, "<files>", "-pthread", "-ldl"},
	}
	err := make([]error, len(argTemplates))
	for i, template := range argTemplates {
		var arg []string
		for _, tt := range template {
			if tt == "<files>" {
				arg = append(arg, files...)
			} else {
				arg = append(arg, tt)
			}
		}
		cmd := exec.Command("gcc", arg...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		if err[i] = cmd.Run(); err[i] != nil {
			continue
		}
		return nil
	}
	return fmt.Errorf("Cannot compile by GCC.\nAppname = %v.\nFiles = %v\nErrors = %v", appName, files, err)
}

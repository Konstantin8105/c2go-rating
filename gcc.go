package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func gccExecution(files ...string) error {
	fmt.Println("GCC: ", files)

	// Generate output file
	dir, err := ioutil.TempDir("", "c2go-rating-gcc")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir) // clean up

	var arg []string
	arg = append(arg, "-o", dir+"/app")
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

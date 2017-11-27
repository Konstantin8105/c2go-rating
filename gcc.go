package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func gccExecution(files ...string) (err error) {
	if *onlyFlag != "" && *onlyFlag != "gcc" {
		return nil
	}
	defer func() {
		if err != nil {
			cErrGCC <- err
		}
	}()
	fmt.Println("GCC  : ", files)

	// Generate output file
	dir, err := ioutil.TempDir("", "c2go-rating-gcc")
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err2 := os.RemoveAll(dir)
			err = fmt.Errorf("%v\n%v", err, err2)
		} else {
			err = os.RemoveAll(dir)
		}
	}() // clean up

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

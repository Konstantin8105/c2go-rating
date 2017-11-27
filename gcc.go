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
		err2 := os.RemoveAll(dir)
		if err != nil {
			err = fmt.Errorf("%v\n%v", err, err2)
		} else {
			err = err2
		}
	}() // clean up

	return runGCC(dir + "/app")
}

func gccExecutionWithResult(files ...string) (_ string, err error) {
	if *onlyFlag != "" && *onlyFlag != "gcc" {
		return "", nil
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
		return "", err
	}
	defer func() {
		err2 := os.RemoveAll(dir)
		if err != nil {
			err = fmt.Errorf("%v\n%v", err, err2)
		} else {
			err = err2
		}
	}() // clean up

	app := dir + "/app"
	err = runGCC(app, files...)
	if err != nil {
		return "", err
	}

	return run(app)
}

func runGCC(app string, files ...string) (err error) {
	var arg []string
	arg = append(arg, "-o", app)
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

func run(app string, args ...string) (result string, err error) {
	cmd := exec.Command(app, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Cannot run app.\nApp = %v\nArg = %v\nErrors = %v\nStderr = %v", app, args, err, stderr.String())
	}
	return out.String(), nil
}

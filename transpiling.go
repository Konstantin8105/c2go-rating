package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func c2goTranspiling(files ...string) error {
	fmt.Println("C2GO : ", files)

	// Generate output file
	dir, err := ioutil.TempDir("", "c2go-rating-c2go")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir) // clean up

	var arg []string
	arg = append(arg, "transpile", "-o", dir+"/1.go")
	arg = append(arg, files...)
	cmd := exec.Command("c2go", arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("c2go : %v\n%v\n%v", err, out.String(), stderr.String())
	}
	return nil
}

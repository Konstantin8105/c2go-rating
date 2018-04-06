package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func c4goTranspiling(files ...string) error {
	return c4goTranspilingWithResult("", files...)
}

func c4goTranspilingWithResult(result string, files ...string) (err error) {
	if *onlyFlag != "" && *onlyFlag != "c4go" {
		return nil
	}
	defer func() {
		cErrc4go <- err
	}()
	// fmt.Println("c4go : ", files)

	// Generate output file
	dir, err := ioutil.TempDir("", "c4go-rating-c4go")
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

	var arg []string
	arg = append(arg, "transpile", "-o", dir+"/1.go")
	arg = append(arg, files...)
	cmd := exec.Command("c4go", arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("c4go : %v\n%v\n%v", err, out.String(), stderr.String())
	}

	// Calculate warnings
	content, err := ioutil.ReadFile(dir + "/1.go")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var counter int
	for _, l := range lines {
		if strings.HasPrefix(l, "// Warning") {
			counter++
		}
	}
	cWarning <- counter

	if result == "" {
		_, err = run("go", []string{"build", dir + "/1.go"}...)
		return
	}

	// compare result
	var c4goResult string
	c4goResult, err = run("go", []string{"run", dir + "/1.go"}...)
	if err != nil {
		return
	}

	if result != c4goResult {
		err = fmt.Errorf("results is different:\n%v\n%v", result, c4goResult)
		return
	}

	return nil
}

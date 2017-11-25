package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func csmithExecute(file string) error {
	cmd := exec.Command("/bin/bash", "-c", "csmith")
	fmt.Println("file ", file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("csmith : %v\n%v\n%v", err, out.String(), stderr.String())
	}

	f, err := os.Create(file)
	defer f.Close()
	_, err = f.Write(out.Bytes())
	if err != nil {
		fmt.Println("err ", err)
		return err
	}
	return nil
}

func main() {
	var data []error
	cErr := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for e := range cErr {
			data = append(data, e)
		}
		wg.Done()
	}()
	singleCcode(cErr)
	triangle(cErr)
	csmith(cErr)
	close(cErr)
	wg.Wait()

	var fail int
	for _, d := range data {
		if d != nil {
			fmt.Println("------")
			fmt.Println(d)
			fail++
		}
	}
	fmt.Println("Fail   results : ", fail)
	fmt.Println("Amount results : ", len(data))
}

func singleCcode(cErr chan<- error) {
	sourceFolder := "./testdata/SingleCcode/"

	// Get all files
	files, err := filepath.Glob(sourceFolder + "*.c")
	if err != nil {
		panic(fmt.Errorf("Error: %v", err))
	}

	// Check in gcc
	for _, file := range files {
		if err := gccExecution(file); err != nil {
			cErr <- err
		}
	}

	// Transpiling by c2go
	for _, file := range files {
		if err := c2goTranspiling(file); err != nil {
			cErr <- err
		}
		cErr <- nil
	}
}

func triangle(cErr chan<- error) {
	sourceFolder := "./testdata/triangle/"
	file := sourceFolder + "triangle.c"

	// Check in gcc
	if err := gccExecution(file, "-lm"); err != nil {
		cErr <- err
	}

	// Transpiling by c2go
	if err := c2goTranspiling(file); err != nil {
		cErr <- err
	} else {
		cErr <- nil
	}
}

func csmith(cErr chan<- error) {
	sourceFolder := "./testdata/csmith/"

	// Get all files
	files, err := filepath.Glob(sourceFolder + "*.c")
	if err != nil {
		panic(fmt.Errorf("Error: %v", err))
	}

	minAmount := 300
	if len(files) < minAmount {
		ch := make(chan string, 10)
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for c := range ch {
					// Generate c code
					// csmith > test.c;
					csmithExecute(c)
				}
				wg.Done()
			}()
		}
		for i := 0; i < minAmount; i++ {
			ch <- fmt.Sprintf("./testdata/csmith/%d.c", i)
		}
		close(ch)
		wg.Wait()

		csmith(cErr)
		return
	}

	// Check in gcc
	for _, file := range files {
		if err := gccExecution("-I./testdata/csmith-git/runtime/", file); err != nil {
			cErr <- err
		}
	}

	// Transpiling by c2go
	for _, file := range files {
		if err := c2goTranspiling("-clang-flag", "-I./testdata/csmith-git/runtime/", file); err != nil {
			cErr <- err
		}
		cErr <- nil
	}
}

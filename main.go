package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	var data []error
	cErr := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for e := range cErr {
			data = append(data, e)
			if e != nil {
				fmt.Println("------")
				fmt.Println(e)
			}
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

	for _, file := range files {
		// Check in gcc
		if err := gccExecution(file); err != nil {
			cErr <- err
			continue
		}
		// Transpiling by c2go
		if err := c2goTranspiling(file); err != nil {
			cErr <- err
			continue
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

func csmith(cErr chan<- error) {
	sourceFolder := "./testdata/csmith/"

	// Get all files
	files, err := filepath.Glob(sourceFolder + "*.c")
	if err != nil {
		panic(fmt.Errorf("Error: %v", err))
	}

	minAmount := 30
	if len(files) < minAmount {
		ch := make(chan string, 10)
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for c := range ch {
					// Generate c code
					// csmith > test.c;
					_ = csmithExecute(c)
				}
				wg.Done()
			}()
		}
		files, err := filepath.Glob(sourceFolder + "*.c")
		if err != nil {
			panic(fmt.Errorf("Error: %v", err))
		}
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		prefix := r1.Intn(10000)
		for i := len(files); i < minAmount; i++ {
			ch <- fmt.Sprintf("./testdata/csmith/%d_%d.c", prefix, i)
		}
		close(ch)
		wg.Wait()

		csmith(cErr)
		return
	}

	for _, file := range files {
		// Check in gcc
		if err := gccExecution("-I./testdata/csmith-git/runtime/", file); err != nil {
			cErr <- err
			continue
		}
		// Transpiling by c2go
		if err := c2goTranspiling("-clang-flag", "-I./testdata/csmith-git/runtime/", file); err != nil {
			cErr <- err
			continue
		}
		cErr <- nil
	}
}

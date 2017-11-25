package main

import (
	"fmt"
	"path/filepath"
	"sync"
)

const (
	// Folder with sqlite files for testing
	sqliteFolder string = "./sqlite/"

	// Folder with GSL files for testing
	gslFolder string = "./gsl/"

	// Folder for output files of GSL
	gslOutput string = "./gslOutput/"
)

// generate C codes
func generate() {
	// Generate c code
	// csmith > test.c;

	// Check by GCC
	// gcc -I./csmith/runtime -O -w test.c -o a.out;
	// ./a.out;

	// Check transpiling c2go
	// c2go transpile -clang-flag="-I./csmith/runtime" ./test.c;

	// Check compile
	// go run ./test.go;
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

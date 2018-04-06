package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

type part struct {
	gcc  []string
	c4go []string
}

var (
	// cErrc4go - channel of errors
	cErrc4go = make(chan error, 20)

	// cErrc4go - channel of errors
	cErrGCC = make(chan error, 20)

	// cInput - just run
	cInput = make(chan part, 20)

	// cInputWithChecking - run and check result
	cInputWithChecking = make(chan part, 20)

	// cWarning - channel for summary warnings in Go code
	cWarning = make(chan int, 20)
)

var (
	partFlag = flag.String("part", "", "Choose: single, triangle, csmith. If nothing is choosed, then start all.")
	onlyFlag = flag.String("only", "", "Choose: gcc, c4go.")
)

func main() {
	flag.Parse()

	var datac4go []error
	var dataGCC []error
	var warnings int

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	for con := 0; con < 10; con++ {
		wg.Add(1)
		go func() {
			for inp := range cInput {
				// Check in gcc
				if err := gccExecution(inp.gcc...); err != nil {
					printReport("GCC", inp, err)
					continue
				}
				// Transpiling by c4go
				if err := c4goTranspiling(inp.c4go...); err != nil {
					printReport("c4go", inp, err)
				}
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for inp := range cInputWithChecking {
				if *onlyFlag == "c4go" {
					cInput <- inp
					continue
				}
				// Check in gcc
				result, err := gccExecutionWithResult(inp.gcc...)
				if err != nil {
					printReport("GCC", inp, err)
					continue
				}
				// Transpiling by c4go
				if err := c4goTranspilingWithResult(result, inp.c4go...); err != nil {
					printReport("c4go", inp, err)
				}
			}
			wg.Done()
		}()
	}

	wg.Add(1)
	go func() {
		switch *partFlag {
		case "":
			folderCcode("./testdata/SingleCcode/")
			folderCcode("./testdata/ac-book/")
			folderCcode("./testdata/book-c-the-examples-and-tasks/")
			folderCcode("./testdata/books-examples/")
			folderCcode("./testdata/C-Deitel-Book/")
			folderCcode("./testdata/c_programming_language_book/")
			folderCcode("./testdata/k-and-r/")
			folderCcode("./testdata/K-and-R-exercises-and-examples")
			folderCcode("./testdata/programming-in-c")
			triangle()
			csmith()
		case "single":
			folderCcode("./testdata/SingleCcode/")
		case "triangle":
			triangle()
		case "csmith":
			csmith()
		}
		close(cInput)
		close(cInputWithChecking)
		wg.Done()
	}()

	wg2.Add(1)
	go func() {
		for e := range cErrc4go {
			datac4go = append(datac4go, e)
		}
		wg2.Done()
	}()
	wg2.Add(1)
	go func() {
		for e := range cErrGCC {
			dataGCC = append(dataGCC, e)
		}
		wg2.Done()
	}()
	wg2.Add(1)
	go func() {
		for e := range cWarning {
			warnings += e
		}
		wg2.Done()
	}()

	wg.Wait()

	close(cErrc4go)
	close(cErrGCC)
	close(cWarning)
	wg2.Wait()

	var fail int
	for _, d := range datac4go {
		if d != nil {
			fail++
		}
	}
	fmt.Println("Fail results   gcc  : ", len(dataGCC))
	fmt.Println("Fail results   c4go : ", fail)
	fmt.Println("Amount warnings     : ", warnings)
	fmt.Println("Amount results c4go : ", len(datac4go))
}

var m sync.Mutex

func printReport(name string, inp part, err error) {
	m.Lock()
	fmt.Println("+=====================+")
	fmt.Println("Name : ", name)
	fmt.Println("Args : ", inp)
	fmt.Println("Err  : ", err)
	m.Unlock()
}

func folderCcode(sourceFolder string) {

	{
		files, err := ioutil.ReadDir(sourceFolder)
		if err != nil {
			cErrc4go <- err
		}

		for _, f := range files {
			if f.IsDir() {
				folderCcode(sourceFolder + f.Name() + "/")
			}
		}
	}

	// Get all files
	files, err := filepath.Glob(sourceFolder + "*.c")
	if err != nil {
		panic(fmt.Errorf("Error: %v", err))
	}

	for _, file := range files {
		cInput <- part{
			gcc:  []string{file, "-lm"},
			c4go: []string{file},
		}
	}
}

func triangle() {
	sourceFolder := "./testdata/triangle/"
	file := sourceFolder + "triangle.c"
	cInput <- part{
		gcc:  []string{file, "-lm"},
		c4go: []string{file},
	}
}

func csmithExecute(file string) (err error) {
	cmd := exec.Command("/bin/bash", "-c", "csmith")
	fmt.Println("file ", file)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("csmith : %v\n%v\n%v", err, out.String(), stderr.String())
	}

	f, err := os.Create(file)
	defer func() { err = f.Close() }()
	_, err = f.Write(out.Bytes())
	if err != nil {
		fmt.Println("err ", err)
		return err
	}
	return nil
}

func csmith() {
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

		csmith()
		return
	}

	for _, file := range files {
		cInputWithChecking <- part{
			gcc:  []string{"-I./testdata/csmith-git/runtime/", file},
			c4go: []string{"-clang-flag", "-I./testdata/csmith-git/runtime/", file},
		}
	}
}

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

// cErrC2GO - channel of errors
var cErrC2GO = make(chan error, 20)

// cErrC2GO - channel of errors
var cErrGCC = make(chan error, 20)

type part struct {
	gcc  []string
	c2go []string
}

var cInput = make(chan part, 20)

var cWarning = make(chan int, 20)

func main() {
	part := flag.String("part", "", "Choose: single, triangle, csmith")
	flag.Parse()

	var dataC2GO []error
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
					fmt.Println(err)
					continue
				}
				// Transpiling by c2go
				if err := c2goTranspiling(inp.c2go...); err != nil {
					fmt.Println(err)
				}
			}
			wg.Done()
		}()
	}

	wg.Add(1)
	go func() {
		switch *part {
		case "":
			folderCcode("./testdata/SingleCcode/")
			// folderCcode("./testdata/ac-book/")
			// folderCcode("./testdata/apue2e/")
			// folderCcode("./testdata/book-c-the-examples-and-tasks/")
			// folderCcode("./testdata/books-examples/")
			// folderCcode("./testdata/C-Deitel-Book/")
			// folderCcode("./testdata/c_programming_language_book/")
			// folderCcode("./testdata/k-and-r/")
			// folderCcode("./testdata/K-and-R-exercises-and-examples")
			// folderCcode("./testdata/programming-in-c")
			// triangle()
			// csmith()
		case "single":
			folderCcode("./testdata/SingleCcode/")
		case "triangle":
			triangle()
		case "csmith":
			csmith()
		}
		close(cInput)
		wg.Done()
	}()

	wg2.Add(1)
	go func() {
		for e := range cErrC2GO {
			dataC2GO = append(dataC2GO, e)
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

	close(cErrC2GO)
	close(cErrGCC)
	close(cWarning)
	wg2.Wait()

	var fail int
	for _, d := range dataC2GO {
		if d != nil {
			fail++
		}
	}
	fmt.Println("Fail results   gcc  : ", len(dataGCC))
	fmt.Println("Fail results   c2go : ", fail)
	fmt.Println("Amount warnings     : ", warnings)
	fmt.Println("Amount results c2go : ", len(dataC2GO))
}

func folderCcode(sourceFolder string) {

	{
		files, err := ioutil.ReadDir(sourceFolder)
		if err != nil {
			cErrC2GO <- err
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
			gcc:  []string{file},
			c2go: []string{file},
		}
	}
}

func triangle() {
	sourceFolder := "./testdata/triangle/"
	file := sourceFolder + "triangle.c"
	cInput <- part{
		gcc:  []string{file, "-lm"},
		c2go: []string{file},
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
		cInput <- part{
			gcc:  []string{"-I./testdata/csmith-git/runtime/", file},
			c2go: []string{"-clang-flag", "-I./testdata/csmith-git/runtime/", file},
		}
	}
}

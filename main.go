package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const (
	// Folder with C code execute file
	// one file - one application
	single string = "./SingleCcode/"

	// c2go application name
	c2go string = "./c2go"

	// transpile flag for c2go
	transpile string = "transpile"
)

func convertFromSourceToAppName(sourceName string) string {
	return sourceName[0:(len(sourceName) - len(filepath.Ext(sourceName)))]
}

func addPrefix(prefix, s string) string {
	return prefix + s
}

func removeGCCfiles(folderName string) {
	files, _ := ioutil.ReadDir(folderName)
	for _, file := range files {
		if filepath.Ext(file.Name()) == "" || filepath.Ext(file.Name()) == ".exe" {
			// Remove application
			err := os.Remove(addPrefix(single, file.Name()))
			if err != nil {
				panic(fmt.Errorf("cannot remove file: %v", addPrefix(folderName, file.Name())))
			}
		}
	}
}

func removeGoFiles(folderName string) {
	files, _ := ioutil.ReadDir(folderName)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".go" {
			// Remove go files
			err := os.Remove(addPrefix(single, file.Name()))
			if err != nil {
				panic(fmt.Errorf("cannot remove file: %v", addPrefix(folderName, file.Name())))
			}
		}
	}
}

func main() {

	// Single Application

	// Remove Go files
	removeGoFiles(single)

	// Remove the gcc result
	removeGCCfiles(single)

	// Get all files
	files, _ := ioutil.ReadDir(single)
	fmt.Println("List of files:")
	for _, file := range files {
		fmt.Printf("%v ", file.Name())
	}
	fmt.Println("")
	fmt.Println("Amount of files : ", len(files))

	// Check in gcc
	// example: gcc -o hello hello.c
	var mistakeFilesGCC []string

	var wg sync.WaitGroup

	type message struct {
		errBody  string
		fileName string
	}

	chanMessage := make(chan message, 10)

	go func() {
		for m := range chanMessage {
			fmt.Println(m.errBody)
			mistakeFilesGCC = append(mistakeFilesGCC, m.fileName)
		}
	}()

	for _, file := range files {
		go func(file os.FileInfo) {
			wg.Add(1)
			defer wg.Done()
			appName := addPrefix(single, convertFromSourceToAppName(file.Name()))
			name := addPrefix(single, file.Name())
			cmd := exec.Command("gcc", "-o", appName, name)
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			///////////////////////
			removeGCCfiles(single) // TODO : try to remove
			///////////////////////
			cmd2 := exec.Command("gcc", "-o", "-std=gnu99", appName, name)
			var out2 bytes.Buffer
			var stderr2 bytes.Buffer
			cmd2.Stdout = &out2
			cmd2.Stderr = &stderr2
			err2 := cmd2.Run()

			if err != nil && err2 != nil {
				m := fmt.Sprintf("=== MISTAKE =======\n")
				m += fmt.Sprintf("Cannot compile by c2go file with name : %v\nApp name: %v\n Error: %v\nError: %v\n", name, appName, stderr.String(), stderr2.String())
				chanMessage <- message{
					errBody:  m,
					fileName: name,
				}
			}
		}(file)
	}
	wg.Wait()
	close(chanMessage)
	// Remove the gcc result
	removeGCCfiles(single)

	// c2go must exist
	_, err := os.Stat(c2go)
	if err != nil {
		panic(fmt.Errorf("Application c2go is not found"))
	}

	// Transpiling by c2go
	var mistakeC2Go int
	var mistakeFilesC2GO []string
	var errC2GO []string
	for _, file := range files {
		goName := addPrefix(single, convertFromSourceToAppName(file.Name())+".go")
		name := addPrefix(single, file.Name())
		cmd := exec.Command(c2go, transpile, "-o", goName, name)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			s := fmt.Sprintf("Command : clang %v -o %v %v\n", transpile, goName, name)
			s += fmt.Sprintf("Cannot compile by gcc file with name : %v\nGo name : %v\nError: %v\n\n", name, goName, stderr.String())
			mistakeC2Go++
			fmt.Printf("=== MISTAKE : %v =======\n", mistakeC2Go)
			fmt.Println(s)
			mistakeFilesC2GO = append(mistakeFilesC2GO, name)
			errC2GO = append(errC2GO, stderr.String())
		}
	}
	// Remove Go files
	removeGoFiles(single)

	// Mistakes is not allowable
	fmt.Println("Amount mistake source by gcc: ", len(mistakeFilesGCC))
	for _, m := range mistakeFilesGCC {
		fmt.Println("\tMistake file : ", m)
	}
	// Calculate rating
	fmt.Println("Amount mistake c2go results: ", mistakeC2Go)
	for i, m := range mistakeFilesC2GO {
		fmt.Println("\tMistake file : ", m)
		fmt.Println("\tError: ", strings.Split(errC2GO[i], "\n")[0])
	}
	fmt.Printf("Result: %v is Ok at %v source c files - %.5v procent. \n", len(files)-mistakeC2Go, len(files), float64(len(files)-mistakeC2Go)/float64(len(files))*100.0)

	// multifile checking
	// main files:
	// studentlistmain.c  queue.h queue.c
	// selectionMain.c intArray.h intArray.c
	// triangle....
}

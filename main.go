package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func removeGCCfiles() {
	files, _ := ioutil.ReadDir(single)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == "" || filepath.Ext(file.Name()) == ".exe" {
			// Remove application
			fileName := single + file.Name()
			err := os.Remove(fileName)
			if err != nil {
				panic(fmt.Errorf("cannot remove file of gcc: %v, %v", file.Name(), err))
			}
		}
	}
}

func removeGoFiles() {
	files, _ := ioutil.ReadDir(single)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".go" {
			// Remove go files
			fileName := single + file.Name()
			err := os.Remove(fileName)
			if err != nil {
				panic(fmt.Errorf("cannot remove file of Go: %v, %v", file.Name(), err))
			}
		}
	}
}

func main() {

	// Single Application

	// Remove Go files
	removeGoFiles()
	defer removeGoFiles()

	// Remove the gcc result
	removeGCCfiles()
	defer removeGCCfiles()

	// Get all files
	files, err := filepath.Glob(single + "*.c")
	if err != nil {
		panic(fmt.Errorf("Error: %v", err))
	}
	fmt.Println("List of files:")
	for _, file := range files {
		fmt.Printf("%v ", file)
	}
	fmt.Println("")
	fmt.Println("Amount of files : ", len(files))

	// Check in gcc
	// example: gcc -o hello hello.c
	var mistakeFilesGCC []string

	for _, file := range files {
		appName := convertFromSourceToAppName(file)
		name := file
		{
			cmd := exec.Command("gcc", "-o", appName, name)
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("=== MISTAKE IN GCC ===")
				fmt.Printf("Cannot compile by gcc file with name : %v\n", name)
				fmt.Printf("App name                             : %v\n", appName)
				fmt.Printf("Error                                : %v\n", stderr.String())
				mistakeFilesGCC = append(mistakeFilesGCC, file)
			} else {
				continue
			}
		}
		{
			cmd := exec.Command("gcc", "-o", "-std=gnu99", appName, name)
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("=== MISTAKE IN GCC WITH STD=GNU99 ===")
				fmt.Printf("Cannot compile by gcc file with name : %v\n", name)
				fmt.Printf("App name                             : %v\n", appName)
				fmt.Printf("Error                                : %v\n", stderr.String())
				mistakeFilesGCC = append(mistakeFilesGCC, file)
			}
		}
	}

	// c2go must exist
	_, err = os.Stat(c2go)
	if err != nil {
		panic(fmt.Errorf("Application c2go is not found"))
	}

	// Transpiling by c2go
	type result struct {
		fileName string
		err      string
	}
	var results []result
	for _, file := range files {
		goName := convertFromSourceToAppName(file) + ".go"
		name := file

		cmd := exec.Command(c2go, transpile, "-o", goName, name)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			s := fmt.Sprintf("Command : clang %v -o %v %v\n", transpile, goName, name)
			s += fmt.Sprintf("Cannot compile by c2go file with name : %v\nGo name : %v\nError: %v\n\n", name, goName, stderr.String())
			results = append(results, result{
				fileName: name,
				err:      stderr.String(),
			})
			fmt.Printf("=== MISTAKE : %v =======\n", len(results))
			fmt.Println(s)
		}
	}

	// Mistakes is not allowable
	fmt.Println("Amount mistake source by gcc: ", len(mistakeFilesGCC))
	for _, m := range mistakeFilesGCC {
		fmt.Println("\tMistake file : ", m)
	}
	// Calculate rating
	fmt.Println("Amount mistake c2go results: ", len(results))
	for _, r := range results {
		fmt.Println("\tMistake file : ", r.fileName)
		fmt.Println("\tError: ", strings.Split(r.err, "\n")[0])
	}
	fmt.Printf("Result: %v is Ok at %v source c files - %.5v procent. \n", len(files)-len(results), len(files), float64(len(files)-len(results))/float64(len(files))*100.0)

	// multifile checking
	// main files:
	// studentlistmain.c  queue.h queue.c
	// selectionMain.c intArray.h intArray.c
	// triangle....
}

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
	singleFolder string = "./SingleCcode/"

	// file sqlite for testing
	sqliteFolder string = "./sqlite/"

	// c2go application name
	c2go string = "./c2go"

	// transpile flag for c2go
	transpile string = "transpile"
)

func convertFromSourceToAppName(sourceName string) string {
	return sourceName[0:(len(sourceName) - len(filepath.Ext(sourceName)))]
}

func removeGCCfiles(folder string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == "" || filepath.Ext(file.Name()) == ".exe" {
			// Remove application
			fileName := folder + file.Name()
			err := os.Remove(fileName)
			if err != nil {
				panic(fmt.Errorf("cannot remove file of gcc: %v, %v", file.Name(), err))
			}
		}
	}
}

func removeGoFiles(folder string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".go" {
			// Remove go files
			fileName := folder + file.Name()
			err := os.Remove(fileName)
			if err != nil {
				panic(fmt.Errorf("cannot remove file of Go: %v, %v", file.Name(), err))
			}
		}
	}
}

func gccExecution(file string) error {
	appName := convertFromSourceToAppName(file)
	name := file
	args := [][]string{
		[]string{"-o", appName, name},
		[]string{"-O", "-o", appName, name, "-lm"},
	}
	for _, arg := range args {
		cmd := exec.Command("gcc", arg...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("=== MISTAKE IN GCC ===")
			fmt.Printf("Cannot compile by gcc file : %v\n", name)
			fmt.Printf("Arguments for compile      : %v\n", arg)
			fmt.Printf("App name                   : %v\n", appName)
			fmt.Printf("Error                      : %v\n", stderr.String())
			continue
		}
		return nil
	}
	return fmt.Errorf("Cannot compile by GCC")
}

func main() {

	// Checking applications and folders
	// for start the checking
	_, err := os.Stat(c2go)
	if err != nil {
		panic(fmt.Errorf("application c2go is not found"))
	}

	_, err = os.Stat(singleFolder)
	if err != nil {
		panic(fmt.Errorf("folder %v is not found", singleFolder))
	}

	_, err = os.Stat(sqliteFolder)
	if err != nil {
		panic(fmt.Errorf("folder %v is not found", sqliteFolder))
	}

	// Single file C source code
	{
		// Remove Go files
		removeGoFiles(singleFolder)
		defer removeGoFiles(singleFolder)

		// Remove the gcc result
		removeGCCfiles(singleFolder)
		defer removeGCCfiles(singleFolder)

		// Single Application

		// Get all files
		files, err := filepath.Glob(singleFolder + "*.c")
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
			if err := gccExecution(file); err != nil {
				mistakeFilesGCC = append(mistakeFilesGCC, file)
			}
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
	}
	for i := 0; i < 5; i++ {
		fmt.Println("=*=")
	}
	{
		// sqlite
		fmt.Println("Analising : SQLITE\n\n")
		// gcc -pthread  *.c -ldl

		// Remove Go files
		removeGoFiles(sqliteFolder)
		defer removeGoFiles(sqliteFolder)

		// Remove the gcc result
		removeGCCfiles(sqliteFolder)
		defer removeGCCfiles(sqliteFolder)

		// checking by GCC
		args := []string{"-o", "sqlite3", sqliteFolder + "shell.c", sqliteFolder + "sqlite3.c", "-pthread", "-ldl"}
		cmd := exec.Command("gcc", args...)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("=== MISTAKE IN GCC ===")
			fmt.Printf("Cannot compile by gcc file : %v\n", sqliteFolder)
			fmt.Printf("Error                      : %v\n", stderr.String())
			fmt.Printf("Args                       : %v\n", args)
			return
		}

		files, err := filepath.Glob(sqliteFolder + "*.c")
		if err != nil {
			panic(fmt.Errorf("Error: %v", err))
		}

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
				fmt.Printf("=== MISTAKE ===\n")
				fmt.Println(s)

			}
		}
	}

	// multifile checking
	// main files:
	// studentlistmain.c  queue.h queue.c
	// selectionMain.c intArray.h intArray.c
}

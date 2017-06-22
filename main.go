package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
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
	var mistakeSource int
	for _, file := range files {
		appName := addPrefix(single, convertFromSourceToAppName(file.Name()))
		name := addPrefix(single, file.Name())
		cmd := exec.Command("gcc", "-o", appName, name)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			mistakeSource++
			fmt.Printf("=== MISTAKE : %v =======\n", mistakeSource)
			fmt.Printf("Cannot compile by c2go file with name : %v\nApp name: %v\n Error: %v\n\n", name, appName, stderr.String())
		}
	}
	// Mistakes is not allowable
	fmt.Println("Amount mistake source: ", mistakeSource)
	if mistakeSource > 0 {
		os.Exit(1)
	}
	// Remove the gcc result
	removeGCCfiles(single)

	// c2go must exist
	_, err := os.Stat(c2go)
	if err != nil {
		panic(fmt.Errorf("Application c2go is not found"))
	}

	// Transpiling by c2go
	var mistakeC2Go int
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
			mistakeC2Go++
			fmt.Printf("=== MISTAKE : %v =======\n", mistakeC2Go)
			fmt.Printf("Cannot compile by gcc file with name : %v\nGo name : %v\nError: %v\n\n", name, goName, stderr.String())
		}
	}
	// Remove Go files
	removeGoFiles(single)

	// Calculate rating
	fmt.Println("Amount mistake c2go results: ", mistakeC2Go)
	fmt.Printf("Result: %v is Ok at %v source c files\n", len(files)-mistakeC2Go, len(files))
}

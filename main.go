package main

import (
	"fmt"
	"os"
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

type result struct {
	fileName string
	err      error
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

	// saving results of gcc
	var mistakeFilesGCC []string

	// Saving results of c2go
	var results []result

	// Amount C source codes in c2go
	var countFiles int

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

		// Check in gcc
		for _, file := range files {
			appName := convertFromSourceToAppName(file)
			if err := gccExecution(appName, file); err != nil {
				mistakeFilesGCC = append(mistakeFilesGCC, file)
			}
		}

		// Transpiling by c2go
		for _, file := range files {
			countFiles++
			goFile := convertFromSourceToAppName(file) + ".go"
			if err := c2goTranspiling(file, goFile); err != nil {
				results = append(results, result{
					fileName: file,
					err:      err,
				})
			}
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Println("=*=")
	}
	{
		// sqlite
		fmt.Println("Analising : SQLITE")
		// gcc -pthread  *.c -ldl

		// Remove Go files
		removeGoFiles(sqliteFolder)
		defer removeGoFiles(sqliteFolder)

		// Remove the gcc result
		removeGCCfiles(sqliteFolder)
		defer removeGCCfiles(sqliteFolder)

		files, err := filepath.Glob(sqliteFolder + "*.c")
		if err != nil {
			panic(fmt.Errorf("Error: %v", err))
		}

		// checking by GCC
		if err := gccExecution("sqlite3", files...); err != nil {
			fmt.Println("=== MISTAKE IN GCC ===")
			fmt.Printf("Cannot compile by gcc file : %v\n", sqliteFolder)
			fmt.Printf("Error                      : %v\n", err)
			return
		}

		for _, file := range files {
			countFiles++
			goFile := convertFromSourceToAppName(file) + ".go"
			if err := c2goTranspiling(file, goFile); err != nil {
				results = append(results, result{
					fileName: file,
					err:      err,
				})
			}
		}
	}

	// multifile checking
	// main files:
	// studentlistmain.c  queue.h queue.c
	// selectionMain.c intArray.h intArray.c

	// Mistakes is not allowable
	fmt.Println("Amount mistake source by gcc: ", len(mistakeFilesGCC))
	for _, m := range mistakeFilesGCC {
		fmt.Println("\tMistake file : ", m)
	}
	// Calculate rating
	fmt.Println("Amount mistake c2go results: ", len(results))
	for _, r := range results {
		fmt.Println("\tMistake file : ", r.fileName)
		fmt.Println("\tError: ", strings.Split(r.err.Error(), "\n")[0])
	}
	fmt.Printf("Result: %v is Ok at %v source c files - %.5v procent. \n", countFiles-len(results), countFiles, float64(countFiles-len(results))/float64(countFiles)*100.0)
}

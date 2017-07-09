package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	// Folder with C code execute file
	// one file - one application
	singleFolder string = "./SingleCcode/"

	// Folder with sqlite files for testing
	sqliteFolder string = "./sqlite/"

	// Folder with GSL files for testing
	gslFolder string = "./gsl/"

	// Folder for output files of GSL
	gslOutput string = "./gslOutput/"

	// c2go application name
	c2go string = "./c2go"
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

	_, err = os.Stat(gslFolder)
	if err != nil {
		panic(fmt.Errorf("folder %v is not found", gslFolder))
	}

	err = os.RemoveAll(gslOutput)
	if err != nil {
		panic(fmt.Errorf("cannot remove %v folder", gslOutput))
	}

	err = os.Mkdir(gslOutput, 0777)
	if err != nil {
		panic(fmt.Errorf("cannot create %v folder", gslOutput))
	}

	// saving results of gcc
	var mistakeFilesGCC []string

	// Saving results of c2go
	var results []result

	// Amount C source codes in c2go
	var countFiles int

	if false {
		// Single file C source code
		fmt.Println("Analising : Single C source codes")

		// Remove Go files
		removeGoFiles(singleFolder)
		defer removeGoFiles(singleFolder)

		// Remove the gcc result
		removeGCCfiles(singleFolder)
		defer removeGCCfiles(singleFolder)

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
	if false {
		// sqlite
		fmt.Println("Analising : SQLITE")

		// Remove Go files
		removeGoFiles(sqliteFolder)
		defer removeGoFiles(sqliteFolder)

		// Remove the gcc result
		removeGCCfiles(sqliteFolder)
		defer removeGCCfiles(sqliteFolder)

		// Get all files
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
	{
		fmt.Println("GSL")
		// Copy files *.c and *.h to GSL output folder
		files, err := filepath.Glob(gslFolder + "*.c")
		if err != nil {
			panic(fmt.Errorf("Error: %v", err))
		}
		fmt.Println("Files = ", files)
		for _, file := range files {
			outputFile := gslOutput + strings.Split(file, "/")[1]
			fmt.Println("outputFile = ", outputFile)
			copyFile(file, outputFile)
		}
		// Create config.h file
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
		fmt.Println("\t\tError: ", strings.Split(r.err.Error(), "\n")[0])
	}
	fmt.Printf("Result: %v is Ok at %v source c files - %.5v procent. \n", countFiles-len(results), countFiles, float64(countFiles-len(results))/float64(countFiles)*100.0)
}

// Copy - copy files
func copyFile(inputFileName, outputFileName string) (err error) {

	if len(inputFileName) == 0 {
		return fmt.Errorf("inputFileName is zero: %s", inputFileName)
	}

	if len(outputFileName) == 0 {
		return fmt.Errorf("inputFileName is zero: %s", outputFileName)
	}

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer func() {
		errFile := inputFile.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%v ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer func() {
		errFile := outputFile.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%v ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}

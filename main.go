package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
		var files []string
		getInternalDirectory(gslFolder, &files)
		for _, file := range files {
			if path.Ext(file) == ".c" || path.Ext(file) == ".h" {
				ff := strings.Split(file, "/")
				outputFile := gslOutput + ff[len(ff)-1]
				copyFile(file, outputFile)
			}
		}

		// Create config.h file
		copyFile(gslFolder+"config.h.in", gslOutput+"config.h")
		prepareConfig(gslOutput + "config.h")

		// Editing of includes

		// Transpiling
		filesC, err := filepath.Glob(gslOutput + "*.c")
		if err != nil {
			panic(fmt.Errorf("Error: %v", err))
		}
		for _, ff := range filesC {
			countFiles++
			goFile := convertFromSourceToAppName(ff) + ".go"
			if err := c2goTranspiling(ff, goFile); err != nil {
				results = append(results, result{
					fileName: ff,
					err:      err,
				})
			}
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
		fmt.Println("\t\tError: ", strings.Split(r.err.Error(), "\n")[0])
	}
	fmt.Printf("Result: %v is Ok at %v source c files - %.5v procent. \n", countFiles-len(results), countFiles, float64(countFiles-len(results))/float64(countFiles)*100.0)
}

func getInternalDirectory(folder string, filesSummary *[]string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(fmt.Errorf("cannot read dir %v", folder))
	}
	for _, file := range files {
		if file.IsDir() {
			var f []string
			var folderName string = folder + file.Name() + "/"
			getInternalDirectory(folderName, &f)
			for i := range f {
				f[i] = folderName + f[i]
			}
			*filesSummary = append(*filesSummary, f...)
		} else {
			*filesSummary = append(*filesSummary, file.Name())
		}
	}
}

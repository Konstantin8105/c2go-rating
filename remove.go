package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func removeGCCfiles(folder string) {
	removeTempFiles(".exe")
	removeTempFiles("")
}

func removeGoFiles(folder string) {
	removeTempFiles(".go")
}

func removeTempFiles(folder string, extension string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == extension {
			// Remove go files
			fileName := folder + file.Name()
			err := os.Remove(fileName)
			if err != nil {
				panic(fmt.Errorf("cannot remove file : %v. Error = %v", file.Name(), err))
			}
		}
	}
}

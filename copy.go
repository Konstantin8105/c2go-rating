package main

import (
	"fmt"
	"io"
	"os"
)

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

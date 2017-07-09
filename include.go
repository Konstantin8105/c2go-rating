package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func changeInclude(file string) {
	inFile, _ := os.Open(file)
	defer func() {
		err := inFile.Close()
		if err != nil {
			fmt.Println("File ====> ", file)
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#include") {
			if strings.Contains(line, "<config.h>") {
				fmt.Println("#include \"config.h\"")
			}
			if strings.Contains(line, "gsl") {
				fmt.Println("#include \"" + line + "\"")
			}
		}
	}
	//TODO : change file
}

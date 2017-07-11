package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO : create the filter
func changeInclude(inputC, outputC string) error {
	inFile, err := os.Open(inputC)
	if err != nil {
		return fmt.Errorf("cannot open file %v\nerr = %v", inputC, err)
	}
	defer func() {
		err := inFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	outFile, err := os.Create(outputC)
	if err != nil {
		return fmt.Errorf("cannot create file %v\nerr = %v", outputC, err)
	}
	defer func() {
		err := outFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	//includeRegexp := regexp.MustCompile(`(^\s*\#\s*include\s*<gsl([^<>]+)>)`)

	for scanner.Scan() {
		line := scanner.Text()
		//if includeRegexp.Match([]byte(line)) {
		//	fmt.Println(line)
		//}
		if strings.Contains(line, "#include") {
			if strings.Contains(line, "<config.h>") {
				fmt.Fprintf(outFile, "#include \"config.h\"\n")
				continue
			}
			t := "gsl/"
			if strings.Contains(line, t) {
				index := strings.Index(line, t) + len(t)
				fmt.Fprintf(outFile, "#include \"%v\"\n", line[index:len(line)-1])
				continue
			}
		}
		fmt.Fprintln(outFile, line)
	}
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config = map[string]string{
	"HAVE_C99_INLINE": "0",
}

func prepareConfig(configName string) {
	inFile, _ := os.Open(configName)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#undef") {
			parts := strings.Split(line, " ")
			value, ok := config[parts[1]]
			if ok {
				fmt.Println("#define " + parts[1] + " " + value)
				continue
			}
			if strings.HasPrefix(parts[1], "HAVE_DECL_") {
				fmt.Println("#define " + parts[1] + " " + "0")
				continue
			}
			fmt.Println("--> ", line)
		}
	}
	panic("dd")
}

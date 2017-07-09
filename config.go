package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config = map[string]string{
	"HAVE_C99_INLINE":            "0",
	"HAVE_GNUX86_IEEE_INTERFACE": "1",
}

var com = []string{
	"IEEE",
	"inline",
	"size_t",
	"volatile",
}

// TODO : add other
func prepareConfig(configFileName string) {
	inFile, _ := os.Open(configFileName)
	defer func() {
		err := inFile.Close()
		if err != nil {
			panic(err)
		}
	}()
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
			for _, c := range com {
				if strings.Contains(line, c) {
					line = comment(line)
				}
			}
			fmt.Println("--> ", line)
		}
	}
}

func comment(line string) string {
	return "/* " + line + " */"
}

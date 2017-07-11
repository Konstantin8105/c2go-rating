package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config = map[string]string{
	"HAVE_GNUX86_IEEE_INTERFACE":        "1",
	"GSL_DISABLE_DEPRECATED":            "1",
	"HAVE_C99_INLINE":                   "0",
	"HAVE_DLFCN_H":                      "0",
	"HAVE_DOPRNT":                       "0",
	"HAVE_EXIT_SUCCESS_AND_FAILURE":     "1",
	"HAVE_EXTENDED_PRECISION_REGISTERS": "1",
	"HAVE_FPU_X86_SSE":                  "0",
	"HAVE_IEEE_COMPARISONS":             "1",
	"HAVE_IEEE_DENORMALS":               "1",
	"HAVE_INLINE":                       "1",
	"HAVE_INTTYPES_H":                   "0",
	"HAVE_LIBM":                         "0",
	"HAVE_MEMCPY":                       "1",
	"HAVE_MEMMOVE":                      "1",
	"HAVE_MEMORY_H":                     "1",
	"HAVE_PRINTF_LONGDOUBLE":            "1",
	"HAVE_STDINT_H":                     "0",
	"HAVE_STDLIB_H":                     "0",
	"HAVE_STRDUP":                       "0",
	"HAVE_STRINGS_H":                    "0",
	"HAVE_STRING_H":                     "0",
	"HAVE_STRTOL":                       "0",
	"HAVE_STRTOUL":                      "0",
	"HAVE_SYS_STAT_H":                   "0",
	"HAVE_SYS_TYPES_H":                  "0",
	"HAVE_UNISTD_H":                     "0",
	"HAVE_VPRINTF":                      "0",
	"PACKAGE":                           "GoGSL",
	"PACKAGE_BUGREPORT":                 "\"\"",
	"PACKAGE_NAME":                      "GoGSL",
}

var com = []string{
	"IEEE",
	"inline",
	"size_t",
	"volatile",
	//"HAVE_IEEEFP_H",
	"HIDE_INLINE_STATIC",
	"__CHAR_UNSIGNED__",
}

// TODO : add more config parameters

func prepareConfig(inputConfig, outputConfig string) {
	inFile, err := os.Open(inputConfig)
	if err != nil {
		panic(fmt.Errorf("cannot open file %v", inputConfig))
	}
	// close file
	defer func() {
		err := inFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	outFile, err := os.Create(outputConfig)
	if err != nil {
		panic(fmt.Errorf("cannot create file %v", outputConfig))
	}
	// close file
	defer func() {
		err := outFile.Close()
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
				fmt.Fprintf(outFile, "#define %v %v\n", parts[1], value)
				continue
			}
			if strings.HasPrefix(parts[1], "HAVE_DECL_") {
				fmt.Fprintf(outFile, "#define %v 0\n", parts[1])
				continue
			}
			for _, c := range com {
				if strings.Contains(line, c) {
					line = comment(line)
				}
			}
		}
		fmt.Fprintf(outFile, "%v\n", line)
	}
}

func comment(line string) string {
	return "/* " + line + " */"
}

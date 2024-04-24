package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	numberFlag := flag.Bool("n", false, "Prints the line numbers.")
	lineBreakNums := flag.Bool("b", false, "Don't print line numbers for empty lines.")
	flag.Parse()

	lineCount := 0
	switch len(flag.Args()) {
	case 0:
		filePrinter(os.Stdin, &lineCount, *numberFlag, *lineBreakNums)
	case 1:
		var input io.Reader
		if flag.Arg(0) != "-" {
			file := fileSource(flag.Arg(0))
			defer file.Close()
			input = file
		} else {
			input = os.Stdin
		}
		filePrinter(input, &lineCount, *numberFlag, *lineBreakNums)
	default:
		for _, arg := range flag.Args() {
			file := fileSource(arg)
			defer file.Close()
			filePrinter(file, &lineCount, *numberFlag, *lineBreakNums)
		}
	}
}

func fileSource(pathToFile string) *os.File {
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "rcat: could not open file: ")
		os.Exit(1)
	}
	return file
}

func filePrinter(input io.Reader, lineCounter *int, printLineNumbers bool, noLineBreaks bool) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if printLineNumbers || (len(scanner.Text()) != 0 && noLineBreaks) {
			*lineCounter++
			fmt.Fprintln(os.Stdout, *lineCounter, " ", scanner.Text())
		} else {
			fmt.Fprintln(os.Stdout, scanner.Text())
		}
	}
}

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
	// lineBreakNums := flag.Bool("b", false, "Don't print line numbers for empty lines.")
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		filePrinter(os.Stdin, *numberFlag)
	case 1:
		var input io.Reader
		if flag.Arg(0) != "-" {
			file := fileSource(flag.Arg(0))
			defer file.Close()
			input = file
		} else {
			input = os.Stdin
		}
		filePrinter(input, *numberFlag)
	default:
		for _, arg := range flag.Args() {
			file := fileSource(arg)
			defer file.Close()
			filePrinter(file, *numberFlag)
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

func filePrinter(input io.Reader, printLineNumbers bool) {
	scanner := bufio.NewScanner(input)
	lineCounter := 0
	for scanner.Scan() {
		lineCounter++
		if printLineNumbers {
			fmt.Fprintln(os.Stdout, lineCounter, " ", scanner.Text())
		} else {
			fmt.Fprintln(os.Stdout, scanner.Text())
		}
	}
}

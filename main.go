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
	flag.Parse()

	for _, arg := range flag.Args() {

		var input io.Reader

		if arg != "-" {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "rcat: could not open file: ")
				os.Exit(1)
			}
			defer file.Close()

			input = file
		} else {
			input = os.Stdin
		}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			fmt.Fprintln(os.Stdout, scanner.Text())
		}
	}
}

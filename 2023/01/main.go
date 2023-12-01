package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Printf("exit due to: %v\n", err)
		os.Exit(1)
	}
}

func run(w io.Writer, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected one argument pointing to the calibration document, instead got %d args", len(args)-1)
	}

	file := args[1]
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %v", file, err)
	}

	cal := decodeCalibrationDocument(f)
	fmt.Fprintf(w, "The calibration value is %d\n", cal)

	return nil
}

// decodeCalibration decodes and sums all the calibration values hidden by the artsy elf.
func decodeCalibrationDocument(r io.Reader) int {
	var cal int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		v := decodeCalibration(sc.Text())
		cal += v
	}
	return cal
}

var digits = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

// decodeCalibration decodes the calibration value hidden by the artsy elf.
func decodeCalibration(line string) int {
	var first, last rune
	var letters []rune

	for _, c := range line {
		if unicode.IsDigit(c) {
			if first == 0 {
				first = c
				last = c
			} else {
				last = c
			}
			letters = nil
		} else {
			letters = append(letters, c)
			for digit, v := range digits {
				idx := strings.Index(string(letters), digit)

				if idx != -1 {
					if first == 0 {
						first = v
						last = v
					} else {
						last = v
					}

					letters = letters[len(letters)-1:]
					break
				}
			}
		}
	}

	cal, _ := strconv.Atoi(string([]rune{first, last}))
	return cal
}

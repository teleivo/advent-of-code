package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if err := run(); err != nil {
		fmt.Errorf("exit with error due to: %v", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}

// decodeCalibration decodes and sums all the calibration values hidden by the artsy elf.
func decodeCalibrationDocument(r io.Reader) int {
	var sum int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		v := decodeCalibration(sc.Text())
		sum += v
	}
	return sum
}

// decodeCalibration decodes the calibration value hidden by the artsy elf.
func decodeCalibration(line string) int {
	var first, last rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			if first == 0 {
				first = c
				last = c
			} else {
				last = c
			}
		}
	}

	cal, _ := strconv.Atoi(string([]rune{first, last}))
	return cal
}

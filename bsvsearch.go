package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bsv := NewBible("Biblia Sacra Vulgata")
	content := loadLinesFromFile("./vuldat_test.txt")
	bsv, _ = bsv.Load(content)

	fmt.Printf("Bible: %v\n", bsv)
}

func loadLinesFromFile(fileLocation string) (lines []string) {
	f, err := os.Open(fileLocation)
	defer f.Close()
	if err != nil {
		fmt.Errorf("Error while trying to open file: %w", err)
		os.Exit(1)
	}

	fScanner := bufio.NewScanner(f)

	lines = make([]string, 0)
	for fScanner.Scan() {
		newLine := fScanner.Text()

		lines = append(lines, newLine)
	}
	if err := fScanner.Err(); err != nil {
		fmt.Errorf("Error reading standard input: %w", err)
		os.Exit(1)
	}

	return lines
}

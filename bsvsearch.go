package main

import (
	"fmt"
)

func main() {
	bsv := NewBible("Biblia Sacra Vulgata")
	content := loadLinesFromFile("./vuldat.txt")
	_, err := bsv.Load(content)
	if err != nil {
		fmt.Printf("Failed to load contents of bible. Quitting: %v", err)
	}
}

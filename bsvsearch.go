package main

import (
	"fmt"
)

func main() {
	bsv := NewBible("Biblia Sacra Vulgata")
	content := loadLinesFromFile("./vuldat.txt")
	bsv, _ = bsv.Load(content)

	fmt.Printf("Jo3 verses:\n")
	jo3Spec := BookSpecification{"Jo3"}
	bf := BetterFilter{}
	for _, v := range bf.Filter(bsv.verses, jo3Spec) {
		fmt.Printf(" - %v is in Jo3", v)
	}
}

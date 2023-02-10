package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type verse struct {
	book    string
	chapter int
	number  int
	content string
}

func main() {
	verses, err := loadVersesFromFile("./vuldat.txt")
	if err != nil {
		fmt.Println("Error occurred in loadVersesFromFile() call", err)
	}

	for i := range verses {
		fmt.Println(verses[i])
	}
}

func createVerseFromLine(line string) verse {
	ss := strings.Split(line, "|")
	chapter, err := strconv.ParseInt(ss[1], 10, 0)
	if err != nil {
		log.Fatalln(err)
	}
	number, err := strconv.ParseInt(ss[2], 10, 0)
	if err != nil {
		log.Fatalln(err)
	}

	return verse{
		book:    ss[0],
		chapter: int(chapter),
		number:  int(number),
		content: ss[3],
	}
}

func loadVersesFromFile(bibleFileLocation string) (verses []string, err error) {
	bf, err := os.Open(bibleFileLocation)
	defer bf.Close()
	if err != nil {
		return []string{}, err
	}

	bfScanner := bufio.NewScanner(bf)

	verses = make([]string, 0)
	for bfScanner.Scan() {
		verse := bfScanner.Text()

		verses = append(verses, verse)
	}
	if err := bfScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}

	return verses, err
}

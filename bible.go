package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bible struct {
	translation string
	verses      []Verse
}

type Verse struct {
	book    string
	chapter int
	number  int
	content string
}

func NewBible(translation string) Bible {
	emptyVerses := make([]Verse, 0, 36000)
	return Bible{translation, emptyVerses}
}

func (b Bible) Load(content []string) (Bible, error) {
	for _, str := range content {
		bo, ch, nu, co := parseLine(str)

		v := Verse{
			book:    bo,
			chapter: ch,
			number:  nu,
			content: co,
		}

		b.verses = append(b.verses, v)
	}

	return b, nil
}

// If results aren't being returned as expected, check capitalization!
func (b Bible) FilterByBook(bookName string) []Verse {
	result := make([]Verse, 0)

	for i, v := range b.verses {
		if v.book == bookName {
			result = append(result, b.verses[i])
		}
	}

	return result
}

// If results aren't being returned as expected, check capitalization!
func (b Bible) FilterByBookAndChapter(bookName string, chapterNumber int) []Verse {
	result := make([]Verse, 0)

	for i, v := range b.verses {
		if v.book == bookName && v.chapter == chapterNumber {
			result = append(result, b.verses[i])
		}
	}

	return result
}

func parseLine(line string) (book string, chapter int, number int, content string) {
	ss := strings.Split(line, "|")
	book = ss[0]
	c, err := strconv.ParseInt(ss[1], 10, 0)
	chapter = int(c)
	if err != nil {
		log.Fatalln(err)
	}
	n, err := strconv.ParseInt(ss[2], 10, 0)
	number = int(n)
	if err != nil {
		log.Fatalln(err)
	}
	content = ss[3]

	return book, chapter, number, content
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

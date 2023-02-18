package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bible struct {
	Translation string
	Verses      []Verse
}

type Verse struct {
	Book    string
	Chapter int
	Number  int
	Content string
}

func NewBible(translation string) Bible {
	emptyVerses := make([]Verse, 0, 36000)
	return Bible{translation, emptyVerses}
}

func (b Bible) Load(content []string) (Bible, error) {
	for _, str := range content {
		bo, ch, nu, co := parseLine(str)

		v := Verse{
			Book:    bo,
			Chapter: ch,
			Number:  nu,
			Content: co,
		}

		b.Verses = append(b.Verses, v)
	}

	return b, nil
}

// If results aren't being returned as expected, check capitalization!
func (b Bible) FilterByBook(bookName string) []Verse {
	result := make([]Verse, 0)

	for i, v := range b.Verses {
		if v.Book == bookName {
			result = append(result, b.Verses[i])
		}
	}

	return result
}

// If results aren't being returned as expected, check capitalization!
func (b Bible) FilterByBookAndChapter(bookName string, chapterNumber int) []Verse {
	result := make([]Verse, 0)

	for i, v := range b.Verses {
		if v.Book == bookName && v.Chapter == chapterNumber {
			result = append(result, b.Verses[i])
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
		log.Fatal(err)
	}
	n, err := strconv.ParseInt(ss[2], 10, 0)
	number = int(n)
	if err != nil {
		log.Fatal(err)
	}
	content = ss[3]

	return book, chapter, number, content
}

func loadLinesFromFile(fileLocation string) (lines []string) {
	f, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)

	lines = make([]string, 0)
	for fScanner.Scan() {
		newLine := fScanner.Text()

		lines = append(lines, newLine)
	}
	if err := fScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

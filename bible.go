package main

import (
	"log"
	"strconv"
	"strings"
)

type bible struct {
	translation string
	verses      []verse
}

type verse struct {
	book    string
	chapter int
	number  int
	content string
}

func NewBible(translation string) bible {
	emptyVerses := make([]verse, 0, 36000)
	return bible{translation, emptyVerses}
}

func (b bible) Load(content []string) (bible, error) {
	for _, str := range content {
		bo, ch, nu, co := parseLine(str)

		v := verse{
			book:    bo,
			chapter: ch,
			number:  nu,
			content: co,
		}

		b.verses = append(b.verses, v)
	}

	return b, nil
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

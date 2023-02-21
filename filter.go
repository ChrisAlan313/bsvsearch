package main

import "github.com/ChrisAlan313/bible"

type Filter struct{}

func (f *Filter) FilterByBook(verses []bible.Verse, book string) []bible.Verse {
	result := make([]bible.Verse, 0)

	for i, v := range verses {
		if v.Book == book {
			result = append(result, verses[i])
		}
	}

	return result
}

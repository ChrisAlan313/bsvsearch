package main

type Filter struct{}

func (f *Filter) FilterByBook(verses []Verse, book string) []Verse {
	result := make([]Verse, 0)

	for i, v := range verses {
		if v.Book == book {
			result = append(result, verses[i])
		}
	}

	return result
}

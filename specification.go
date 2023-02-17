package main

type Specification interface {
	IsSatisfied(v Verse) bool
}

type BookSpecification struct {
	book string
}

type ChapterSpecification struct {
	chapter int
}

type BetterFilter struct{}

func (b BookSpecification) IsSatisfied(v Verse) bool {
	return v.book == b.book
}

func (c ChapterSpecification) IsSatisfied(v Verse) bool {
	return v.chapter == c.chapter
}

func (f *BetterFilter) Filter(verses []Verse, spec Specification) []Verse {
	result := make([]Verse, 0)
	for i, v := range verses {
		if spec.IsSatisfied(v) {
			result = append(result, verses[i])
		}
	}
	return result
}

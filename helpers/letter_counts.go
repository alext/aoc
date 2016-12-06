package helpers

type LetterCounts [26]struct {
	Letter rune
	Count  uint
}

func (l *LetterCounts) Count(letter rune) {
	if 'a' > letter || letter > 'z' {
		panic("letter out of range a-z")
	}
	l[letter-'a'].Letter = letter
	l[letter-'a'].Count += 1
}

func (l *LetterCounts) Len() int      { return len(l) }
func (l *LetterCounts) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l *LetterCounts) Less(i, j int) bool {
	if l[i].Count == l[j].Count {
		return l[i].Letter <= l[j].Letter
	}
	return l[i].Count > l[j].Count
}

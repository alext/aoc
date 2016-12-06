package helpers

type LetterCounts [26]struct {
	Letter rune
	Count  uint
}

func (l *LetterCounts) Len() int      { return len(l) }
func (l *LetterCounts) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l *LetterCounts) Less(i, j int) bool {
	if l[i].Count == l[j].Count {
		return l[i].Letter <= l[j].Letter
	}
	return l[i].Count > l[j].Count
}

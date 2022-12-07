package pkg

type Queue []rune

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(character rune) {
	*q = append(*q, character)
}

func (q *Queue) Dequeue() (rune, bool) {
	if q.IsEmpty() {
		return '_', false
	}
	firstElement := (*q)[0]

	*q = (*q)[1:]
	return firstElement, true
}

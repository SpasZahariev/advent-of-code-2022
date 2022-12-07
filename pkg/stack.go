package pkg

// Golang does not have a stack in the standard library so I am implementing my own
type Stack []string

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // appending the new element to the end (top) of the list
}

// Push multiple items to the stack
func (s *Stack) PushAll(newElements []string) {
	*s = append(*s, newElements...) // ... expands the list into all its items (append accepts an unlimited number of inputs)
}

// Remove and return the top element of the stack
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	topIndex := len(*s) - 1
	element := (*s)[topIndex]
	*s = (*s)[:topIndex] // Removes the top element from the stack
	return element, true
}

// Gives the last x elements in the order they appear in the list
func (s *Stack) PopMultiple(numberOfElements int) ([]string, bool) {
	if s.IsEmpty() || len(*s) < numberOfElements || numberOfElements == 0 {
		return nil, false
	}

	newFinalIndex := len(*s) - numberOfElements
	elements := (*s)[newFinalIndex:]
	*s = (*s)[:newFinalIndex] // Removes the top element from the stack
	return elements, true
}

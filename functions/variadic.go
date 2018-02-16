package functions

//TopNumberFromList finds the max number in numbers
func TopNumberFromList(numbers ...int) int {
	best := numbers[0]

	for _, i := range numbers {
		if i > best {
			best = i
		}
	}
	return best
}

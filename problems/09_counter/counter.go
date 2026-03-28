package counter

func NewCounter() func() int {
	var count int

	return func() int {
		count++
		return count
	}
}

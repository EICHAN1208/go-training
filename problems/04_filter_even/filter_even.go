package filter_even

func FilterEven(nums []int) []int {
	s := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			s = append(s, num)
		}
	}
	return s
}

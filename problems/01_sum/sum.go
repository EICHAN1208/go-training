package sum

func Sum(nums []int) int {
	sum := 0
	for num := range nums {
		sum += num
	}
	return sum
}

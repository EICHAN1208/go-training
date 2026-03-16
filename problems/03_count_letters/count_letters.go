package count_letters

func CountLetters(s string) map[string]int {
	m := make(map[string]int)
	for _, r := range s {
		m[string(r)]++
	}
	return m
}

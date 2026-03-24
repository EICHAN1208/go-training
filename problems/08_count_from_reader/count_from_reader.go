package count_from_reader

import "io"

func CountFromReader(r io.Reader) map[string]int {
	m := make(map[string]int)
	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)

		for _, b := range buf[:n] {
			if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
				m[string(b)]++
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
	}

	return m
}

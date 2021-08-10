package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return -1, 	errors.New("sequences have different lengths")
	}

	for i, c := range b {
		if uint8(c) != a[i] {
			distance ++
		}
	}

	return distance, nil
}

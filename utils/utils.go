package utils

func CreateReversed[T any](s []T) []T {
	length := len(s)
	reversed := make([]T, length)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = s[j]
	}

	return reversed
}

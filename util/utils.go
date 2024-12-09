package util

func Combinations[T any](list []T) [][2]T {
	comb := [][2]T{}
	plen := len(list)

	for i := 0; i < plen-1; i++ {
		for j := i + 1; j < plen; j++ {
			comb = append(comb, [2]T{list[i], list[j]})
		}
	}

	return comb
}

func Contains[T comparable](s []T, v T) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}

func InRange(a, min, max int) bool {
	return min < a && a < max
}

func Abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func Absdiff(x, y int) int {
	return Abs(x - y)
}

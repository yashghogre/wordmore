package packages

import "math/rand"

func Zip(slice1 []rune, slice2 []rune) [][2]rune {
	length := min(len(slice1), len(slice2))
	result := make([][2]rune, length)

	for i := 0; i < length; i++ {
		result[i] = [2]rune{slice1[i], slice2[i]}
	}

	return result
}

func RtoI(key rune) int {
	var value int
	var set = map[rune]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26, '#': 0,
	}

	value = set[key]
	return value
}

func ItoR(key int) rune {
	var value rune
	var set = map[int]rune{
		1: 'a', 2: 'b', 3: 'c', 4: 'd', 5: 'e', 6: 'f', 7: 'g', 8: 'h', 9: 'i', 10: 'j', 11: 'k', 12: 'l', 13: 'm', 14: 'n', 15: 'o', 16: 'p', 17: 'q', 18: 'r', 19: 's', 20: 't', 21: 'u', 22: 'v', 23: 'w', 24: 'x', 25: 'y', 26: 'z', 0: '#',
	}

	value = set[key]
	return value
}

func MultinomialSample(n int, probs []float32) int {
	result := make([]int, len(probs))
	var idx int
	for i := 0; i < n; i++ {
		r := rand.Float32()
		for j, p := range probs {
			r -= p
			if r < 0 {
				result[j]++
				break
			}
		}
	}

	for i := 0; i < len(result); i++ {
		if result[i] == 1 {
			idx = i
			break
		}
	}
	return idx
}

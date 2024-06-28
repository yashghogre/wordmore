package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wordmore/packages"
)

func main() {
	var words_arr []string

	file, err := os.Open("words.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		words_arr = append(words_arr, word)
	}

	// fmt.Println("This is the first word: ", words_arr[0])
	// fmt.Println("This is the length of array: ", len(words_arr))

	// b := make(map[string]int)

	var N [27][27]float32
	for i := 0; i < 27; i++ {
		for j := 0; j < 27; j++ {
			N[i][j] = 1
		}
	}

	var t_char rune = '#'
	for _, word := range words_arr {

		wordRunes := []rune(word)
		processedWord := append([]rune{t_char}, wordRunes...)
		processedWord = append(processedWord, t_char)

		zipped := packages.Zip(processedWord, processedWord[1:])

		for _, ch := range zipped {
			// b[string(ch2)] += 1
			ix1 := packages.RtoI(ch[0])
			ix2 := packages.RtoI(ch[1])

			N[ix1][ix2] += 1
		}
	}

	var Q [27]float32
	for i := 0; i < 27; i++ {
		Q[i] = 0
	}

	for i := 0; i < 27; i++ {
		for j := 0; j < 27; j++ {
			Q[i] += N[i][j]
		}
	}

	var P [27][27]float32
	for i := 0; i < 27; i++ {
		for j := 0; j < 27; j++ {
			P[i][j] = N[i][j] / Q[i]
		}
	}

	for i := range 10 {
		var out []rune
		var ix int = 0

		for {
			// var p = P[ix]
			ix = packages.MultinomialSample(1, P[ix][:])
			if ix == 0 {
				break
			}
			out = append(out, packages.ItoR(ix))
		}
		fmt.Printf("%d) %c\n", i, out)
	}

	// value := packages.RtoI('#')

	// fmt.Println(P[0])
	// fmt.Println(value)
}

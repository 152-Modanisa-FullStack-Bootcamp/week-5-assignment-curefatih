package assignment

import (
	"math"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	if sum < x {
		return sum, true
	}

	return sum, false
}

func CeilNumber(f float64) float64 {
	unit := 0.25
	return math.Ceil(f/unit) * unit
}

func AlphabetSoup(s string) string {

	var result = []rune(s)
	// sort the string with buble sort
	// we can use any sort algo here
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return string(result)
}

func StringMask(s string, n uint) string {

	if len(s) == 0 {
		return "*"
	}

	var result = ""

	if n >= uint(len(s)) {
		n = 0
	}

	for i := 0; i < len(s); i++ {
		if i >= int(n) {
			result += "*"
		} else {
			result += string(s[i])
		}
	}

	return result
}

func WordSplit(arr [2]string) string {

	input := arr[0]
	words := strings.Split(arr[1], ",")

	var result []string

	// loop over input string
	for i := 0; i < len(input); i++ {

		shouldBreak := true
		// loop over words
		for j := 0; j < len(words); j++ {
			// check if the word is in the input string
			if i+len(words[j]) <= len(input) && input[i:len(words[j])+i] == words[j] {
				// if yes, add the word to result
				result = append(result, words[j])
				i += len(words[j]) - 1
				shouldBreak = false
				break
			}
		}

		if shouldBreak {
			return "not possible"
		}
	}

	if len(result) == 0 {
		return "not possible"
	}

	// join the result
	return strings.Join(result, ",")
}

func VariadicSet(i ...interface{}) []interface{} {

	// set for checking duplicate
	set := make(map[interface{}]struct{}, len(i))
	var result []interface{}

	for _, v := range i {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

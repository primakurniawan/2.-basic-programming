package main

import "fmt"

func palindrome(input string) bool {

	// your code here
	for i, char := range input {
		// cast the rune to a string for printing

		if string(char) != string(input[len(input)-(i+1)]) {
			return false
		}
		// if math.Ceil(len(input)/2) == i  {
		// 	return true
		// }
	}

	return true

}

func main() {
	fmt.Println(palindrome("katakx"))
}

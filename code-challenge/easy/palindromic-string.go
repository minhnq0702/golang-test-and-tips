// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/palindrome-check-2/
package easy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	// ! GO 1.22
	// lst_s := strings.Split(s, "")
	// slices.Reverse(lst_s)
	// s2 := strings.Join(lst_s, "")

	s2 := ""
	for _, c := range s {
		s2 = string(c) + s2
	}

	return s == s2
}

func PalindromicString() {
	buf := bufio.NewReader(os.Stdin)
	line, err := buf.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	line = strings.TrimRight(line, "\n")
	if isPalindrome(line) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

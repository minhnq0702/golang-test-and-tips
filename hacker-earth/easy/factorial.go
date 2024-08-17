// https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/find-factorial/
package easy

import "fmt"

// Get factorial number of n
func compute_factorial(n int) int64 {
	if n == 0 || n == 1 {
		return int64(1)
	}
	return int64(n) * compute_factorial(n-1)
}

func Factorial(n int) {
	fmt.Println(compute_factorial(n))
}

/*
https://www.hackerearth.com/practice/algorithms/searching/linear-search/practice-problems/algorithm/find-mex-62916c25/
Find minimal excluded
*/
package easy

import (
	"fmt"
	"strconv"
)

func findMex(arr []int) []string {
	res := []string{}
	existed := map[int]bool{}
	mex := 0
	for _, n := range arr {
		existed[n] = true
		for {
			if _, ok := existed[mex]; !ok {
				res = append(res, strconv.Itoa(mex))
				break
			}
			mex++
		}
	}
	return res
}

func TestMex() {
	res := findMex([]int{1, 0, 5, 5, 3})
	fmt.Println(res)

	// buffer := bufio.NewReader(os.Stdin)
	// n, _ := buffer.ReadString('\n')
	// str, _ := buffer.ReadString('\n')
	// arr := make([]int, 10, 10)
	// for _, c := range strings.Split(strings.TrimSpace(str), " ") {
	// 	i, _ := strconv.ParseInt(c, 10, 0)
	// 	arr = append(arr, int(i))
	// }
}

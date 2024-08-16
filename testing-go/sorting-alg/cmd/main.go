package main

import (
	"fmt"
	sortingalg "sorting-alg"
	testcase "sorting-alg/test-case"
	"time"
)

func PrepareTest() []testcase.SortArray {
	case1 := testcase.SortArray{2, 2, 5, 123, 657, 8, -1, 99, -293, 12}
	case2 := testcase.SortArray{10, 80, 30, 90, 40, 50, 70}

	return []testcase.SortArray{
		case1,
		case2,
		testcase.CreateRandomSlice(100),
		testcase.CreateRandomSlice(1000),
		testcase.CreateRandomSlice(10000),     // 10K
		testcase.CreateRandomSlice(100000),    // 100K
		testcase.CreateRandomSlice(1000000),   // 1M
		testcase.CreateRandomSlice(10000000),  // 10M
		testcase.CreateRandomSlice(100000000), // 100M
		testcase.CreateRandomSlice(200000000), // 200M
		// testcase.CreateRandomSlice(1000000000), // 1B
	}
}

func main() {
	for _, testCase := range PrepareTest() {
		start := time.Now()

		// sortingalg.LibSort(testCase)
		// fmt.Printf("Complete sort %d in %f\n", len(testCase), time.Since(start).Seconds())

		// start := time.Now()
		// result := sortingalg.BubbleSort(append(testcase.SortArray{}, testCase...))
		// fmt.Printf("Complete sortingalg.BubbleSort() %f in %d\n", time.Since(start).Seconds(), len(result))

		sortingalg.QuickSort(testCase, 0, int64(len(testCase)-1))
		fmt.Printf("Complete sort %d in %f\n", len(testCase), time.Since(start).Seconds())

		fmt.Println()
	}
}

package sortingalg

import (
	"sort"
	testcase "sorting-alg/test-case"
)

func LibSort(input testcase.SortArray) {
	// test := ([]int64)(input)
	sort.Sort(input)
}

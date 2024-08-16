package sortingalg

import (
	testcase "sorting-alg/test-case"
)

func QuickSortPartition(arr testcase.SortArray, low, hight int64) int64 {
	// fmt.Printf("partition %v  at pos===> %d -> %d \n", arr, low, hight)
	pi := arr[hight]
	pos := low
	for i := low; i <= hight; i++ {
		if arr[i] < pi {
			arr[pos], arr[i] = arr[i], arr[pos]
			// test := &arr[i]
			pos++
		}
	}
	arr[pos], arr[hight] = arr[hight], arr[pos]
	return pos
}

func QuickSort(input testcase.SortArray, low, hight int64) testcase.SortArray {
	if low < hight {
		pos := QuickSortPartition(input, low, hight)
		// lower bound
		QuickSort(input, low, pos-1)
		// upper bound
		QuickSort(input, pos+1, hight)
	}
	return testcase.SortArray{}
}

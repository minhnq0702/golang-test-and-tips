package sortingalg

import testcase "sorting-alg/test-case"

/*
* Complete sort 0.000001 in 10
* Complete sort 0.000014 in 100
* Complete sort 0.000650 in 1000
* Complete sort 0.048538 in 10000
* Complete sort 9.718690 in 100000
* Complete sort 1073.840738 in 1000000
 */
func BubbleSort(input testcase.SortArray) testcase.SortArray {
	var first int64
	var second int64
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			first = input[i]
			second = input[j]
			if second < first {
				tmp := first
				input[i] = second
				input[j] = tmp
			}
		}
	}
	return input
}

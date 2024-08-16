package testcase

import (
	"time"

	"math/rand"
)

type SortArray []int64

func (x SortArray) Len() int           { return len(x) }
func (x SortArray) Less(i, j int) bool { return x[i] < x[j] }
func (x SortArray) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func CreateRandomSlice(n int64) SortArray {
	randomer := rand.New(rand.NewSource(time.Now().Unix()))
	res := make([]int64, n)
	for i := range res {
		res[i] = randomer.Int63()
	}

	return SortArray(res)
}

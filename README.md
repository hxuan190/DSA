# DSA
Neither Minimum nor Maximum
// import (
// 	"math"
// )

// func findNonMinOrMax(nums []int) int {
// 	max, min := math.MinInt64, math.MinInt64
// 	if len(nums) == 2 {
// 		return -1
// 	}
// 	for _, v := range nums {
// 		if v == max || v == min {
// 			continue
// 		}

// 		switch {
// 		case v > max:
// 			max, min = v, max
// 		case v > min:
// 			min = v
// 		}
// 	}
//     if min == math.MinInt64 {
//         return -1
//     }
//     return min
// }
func findNonMinOrMax(nums []int) int {
if len(nums)==2 || len(nums)==1{
return -1
}
sort.Ints(nums)
return nums[1]
}


Third Maximum Number

// // built-in heap
// import(
//     "container/heap"
// )

// type MinHeap []int

// func (m MinHeap) Less(i, j int) bool {
//     return m[i] < m[j]
// }

// func (m MinHeap) Len() int {
//     return len(m)
// }

// func (m MinHeap) Swap(i, j int) {
//     m[i], m[j] = m[j], m[i]
// }

// func (m *MinHeap) Push(n interface{}) {
//     *m = append(*m, n.(int))
// }

// func (m *MinHeap) Pop() interface{} {
//     temp := (*m)[len(*m)-1]
//     *m = (*m)[:len(*m)-1]
//     return temp
// }
// func thirdMax(nums []int) int {
//     mh := &MinHeap{}
//     heap.Init(mh)
//     var seen = make(map[int]bool)
//     for _,num := range nums {
//         if _,exists := seen[num];exists {
//             continue
//         }
//         seen[num] = true
//         heap.Push(mh,num)
//         if mh.Len() > 3 {
//             heap.Pop(mh)
//         }
//     }
//     if mh.Len() == 3 {
//         return heap.Pop(mh).(int)
//     }
//     for mh.Len() >1 {
//         heap.Pop(mh)
//     }

//     return heap.Pop(mh).(int)
// }
import "math"

func thirdMax(nums []int) int {
max, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v == max || v == second || v == third {
			continue
		}
		switch {
		case v > max:
			max,second,third = v,max,second

		case v > second:
			second ,third = v,second

		case v > third:
			third = v
		}
	}
	if third == math.MinInt64 {
		return max
	}
	return third
}


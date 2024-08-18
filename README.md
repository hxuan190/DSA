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

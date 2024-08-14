/*
*
https://leetcode.com/problems/two-sum/description/
*/
package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	var resultNums []int
	// nums: []int{3,2,4},
	// 		target: 6,

	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				resultNums = []int{
					i,j,
				}
			}
		}
	}
	return resultNums
}


func TestMain(t *testing.T)  {
	// var nums = []int{3,2,4}
	// var target = 6

	type argumentParams struct {
		nums []int
		target int
	}

	test := []struct{
		name string
		argument argumentParams
		want []int
	}{
		{"Example 1", argumentParams{
			nums: []int{2,7,11,15},
			target: 9,
		}, []int{0,1}},
		{"Example 2", argumentParams{
			nums: []int{3,2,4},
			target: 6,
		}, []int{1,2}},
		{"Example 3", argumentParams{
			nums: []int{3,3},
			target: 6,
		}, []int{0,1}},
		{"Example 4", argumentParams{
			nums: []int{1, 3, 7, 9, 2},
			target: 11,
		}, []int{3,4}},
		{"Example 5", argumentParams{
			nums: []int{1, 3, 7, 9, 2},
			target: 25,
		}, []int(nil)},
		{"Example 6", argumentParams{
			nums: []int{3,2,3},
			target: 6,
		}, []int{0,2}},
		{"Example 7", argumentParams{
			nums: []int{},
			target: 1,
		}, []int(nil)},
		{"Example 8", argumentParams{
			nums: []int{5},
			target: 5,
		}, []int(nil)},
		{"Example 9", argumentParams{
			nums: []int{1,6},
			target: 7,
		}, []int{0,1}},
	}

	for _, v := range test {
		t.Run(v.name,func(t *testing.T) {
			result := twoSum(v.argument.nums,v.argument.target)
			assert.Equal(t,v.want, result)
		})
	}
}
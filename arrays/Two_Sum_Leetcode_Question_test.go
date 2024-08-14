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

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] + nums[i+1] == target {
			resultNums = []int{
				i,i+1,
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
	}

	for _, v := range test {
		t.Run(v.name,func(t *testing.T) {
			result := twoSum(v.argument.nums,v.argument.target)
			assert.Equal(t,v.want, result)
		})
	}
}
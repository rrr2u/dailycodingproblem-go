package day243

import "testing"

var testcases = []struct {
	nums []int
	k    int
	sum  int
}{
	{[]int{5, 1, 2, 7, 3, 4}, 3, 8},
	{[]int{3, 1, 1}, 3, 3},
	{[]int{10, 1, 1}, 4, 0},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 11},
	{[]int{5, 1, 2, 7, 3, 4}, 2, 14},
	{[]int{5, 1, 2, 7, 3, 4}, 1, 22},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 10, 99},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 8, 99},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 7, 99},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 6, 111},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 5, 143},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 4, 161},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 3, 203},
	{[]int{89, 72, 92, 51, 34, 26, 51, 17, 32, 99}, 2, 304},
}

func TestMinimizePartitionMaximumSum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if sum := MinimizePartitionMaximumSum(tc.nums, tc.k); sum != tc.sum {
			t.Errorf("Expected %v, got %v", tc.sum, sum)
		}
	}
}

func BenchmarkMinimizePartitionMaximumSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimizePartitionMaximumSum(tc.nums, tc.k)
		}
	}
}

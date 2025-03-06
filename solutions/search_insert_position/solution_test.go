package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	// nums := []int{1, 3, 5, 6}
	nums := []int{1, 3, 5}
	target := 2
	expected := 1

	res := searchInsert(nums, target)

	if expected != res {
		t.Fatalf("Failed: expected: %d recieved: %d", expected, res)
	}

	t.Logf("Success: expected: %d recieved: %d", expected, res)
}

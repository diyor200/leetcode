package rotate_list

import "testing"

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 14

	expected := []int{5, 6, 1, 2, 3, 4}
	res := rotate(nums, k)

	if len(expected) != len(res) {
		t.Fatalf("Failed: \nexpected: %v\nreceived: %v", res, expected)
	}

	for i := range res {
		if res[i] != expected[i] {
			t.Fatalf("Failed: \nexpected: %v\nreceived: %v", res, expected)
		}
	}

	t.Log("Success")
}

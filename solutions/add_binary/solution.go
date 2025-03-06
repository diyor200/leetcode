package add_binary

// problem url: https://leetcode.com/problems/add-binary/description/
func AddBinary(a string, b string) string {
	var (
		carry      string
		difference int
		res        string
	)
	if len(a) < len(b) {
		a, b = b, a
	}

	difference = len(a) - len(b)
	for i := 0; i < difference; i++ {
		b = "0" + b
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == "1" {
				res = "1" + res
			} else {
				res = "0" + res
				carry = "1"
			}
		} else if a[i] == '1' && b[i] == '0' {
			if carry == "1" {
				res = "0" + res
			} else {
				res = "1" + res
			}
		} else if a[i] == '0' && b[i] == '1' {
			if carry == "1" {
				res = "0" + res
			} else {
				res = "1" + res
			}
		} else {
			if carry == "1" {
				res = "1" + res
				carry = "0"
			} else {
				res = "0" + res
			}
		}
	}

	if carry == "1" {
		res = "1" + res
	}

	return res
}

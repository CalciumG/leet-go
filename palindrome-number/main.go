package main

func isPalindrome(x int) bool {
	return helper(x, x)
}

func helper(x int, y int) bool {
	// check number isn't negative or ends with 0
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	if y < 0 {
		return false
	}
	if y == 0 {
		return x == 0
	}

	if helper(x/10, y/10) {
		return x%10 == y%10
	}
	return false
}

func main() {
	x := 121
	result := isPalindrome(x)
	println(result)
}

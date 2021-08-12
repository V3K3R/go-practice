package palindromenumber 

func isPalindrome(x int) bool {
	
	z := 0
	y := x
	for y > 0 {
		z = z * 10 + y % 10
		y /= 10
	}
	return x == z 
}
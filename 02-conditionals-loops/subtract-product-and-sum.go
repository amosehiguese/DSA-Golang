package sub

func SubtractProductAndSum(n int) int {
	var product int = 1
	var sum = 0
	for n > 0 {
		rem := n % 10
		product *= rem
		sum += rem
		n /= 10
	}
	result := product - sum
	return result
}

package leetcode_goversion

func isMonotonic(A []int) bool {
	var incre = true
	var decre = true
	for i := 0; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			decre = false
		}
	}
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			incre = false
		}
	}
	return incre || decre
}

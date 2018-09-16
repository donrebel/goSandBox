package datastructures

// ArSum ...
func ArSum(ar []int) int {
	s := 0
	for _, x := range ar {
		s += x
	}
	return s
}

func BinarySearch(ar []int, k int) bool {
	var mid int
	low := 0
	high := len(ar) - 1
	for low <= high {
		mid = low + (high-low)/2
		if ar[mid] == k {
			return true
		}
		if ar[mid] < k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func BinarySearchGoStyle(ar []int, k int) (int, bool) {
	var mid int
	low := 0
	high := len(ar) - 1
	for low <= high {
		mid = low + (high-low)/2
		if ar[mid] == k {
			return mid, true
		}
		if ar[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

func RotateArray(ar []int, k int) []int {
	if len(ar) > k && k > 0 {
		rotateSubArray(ar, 0, k-1)
		rotateSubArray(ar, k, len(ar)-1)
		rotateSubArray(ar, 0, len(ar)-1)
	}
	return ar
}
func rotateSubArray(ar []int, x int, y int) {
	for x < y {
		ar[x], ar[y] = ar[y], ar[x]
		x++
		y--
	}
}

func MaxContiguousSum(ar []int) int {
	localMax := 0
	globalMax := 0
	for _, x := range ar {
		localMax += x
		if localMax >= globalMax {
			globalMax = localMax
		}
		if localMax < 0 {
			localMax = 0
		}
	}
	return globalMax
}

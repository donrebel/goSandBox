package algorithms

func BinarySearch(s []int, x int) bool {
	minIdx := 0
	maxIdx := len(s)
	var midIdx int
	for minIdx < maxIdx {
		midIdx = (minIdx + maxIdx) / 2
		// midIdx = minIdx + (maxIdx-minIdx)/2
		if s[midIdx] == x {
			return true
		}
		if s[midIdx] < x {
			minIdx = midIdx
		} else {
			maxIdx = midIdx
		}
	}
	return false
}

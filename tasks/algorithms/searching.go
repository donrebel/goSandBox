package algorithms

func BinarySearch(s []int, x int) bool {
	lowIdx := 0
	highIdx := len(s) - 1
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if s[midIdx] == x {
			return true
		}
		if s[midIdx] > x {
			highIdx = midIdx - 1
		} else {
			lowIdx = midIdx + 1
		}
	}
	return false
}

func SeparateEvenAndOddNumbers(s []int) {
	lowIdx := 0
	highIdx := len(s) - 1
	for lowIdx < highIdx {
		if s[lowIdx]%2 == 0 {
			lowIdx++
		} else if s[highIdx]%2 != 0 {
			highIdx--
		} else {
			s[lowIdx], s[highIdx] = s[highIdx], s[lowIdx]
		}
	}
}

func GetMedian(s []int) float32 {
	size := len(s)
	if size%2 != 0 {
		return float32(QuickSelect(s, size/2))
	} else {
		return (float32(QuickSelect(s, size/2-1)) + float32(QuickSelect(s, size/2))) / 2
	}
}

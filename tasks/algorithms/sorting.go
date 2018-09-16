package algorithms

//More ...
func More(a int, b int) bool {
	return a > b
}

//Less ...
func Less(a int, b int) bool {
	return a < b
}

func swap(arr []int, lowIdx int, highIdx int) {
	arr[lowIdx], arr[highIdx] = arr[highIdx], arr[lowIdx]
}

//BubbleSort ...
func BubbleSort(s []int, comp func(int, int) bool) {
	size := len(s)
	swap := 1
	for i := 0; i < size-1 && swap == 1; i++ {
		swap = 0
		for j := 1; j < size-i; j++ {
			if comp(s[j-1], s[j]) {
				swap = 1
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

// InsertionSort ...
func InsertionSort(s []int, comp func(int, int) bool) {
	size := len(s)
	var i, j, temp int
	for i = 1; i < size; i++ {
		temp = s[i]
		for j = i; j > 0 && comp(s[j-1], temp); j-- {
			s[j] = s[j-1]
		}
		s[j] = temp
	}
}

// SelectionSort ...
func SelectionSort(s []int, comp func(int, int) bool) {
	size := len(s)
	var maxIdx int
	for i := 0; i < size; i++ {
		maxIdx = 0
		for j := 0; j < size-i; j++ {
			if s[j] > s[maxIdx] {
				maxIdx = j
			}
		}
		s[maxIdx], s[size-i-1] = s[size-i-1], s[maxIdx]
	}
}

// MergeSort ...
func MergeSort(s []int, comp func(int, int) bool) {
	size := len(s)
	tempArray := make([]int, size)
	slice(s, tempArray, 0, size-1, comp)
}

func slice(arr []int, tempArray []int, lowIdx int, highIdx int, comp func(int, int) bool) {
	if lowIdx >= highIdx {
		return
	}
	midIdx := (lowIdx + highIdx) / 2
	slice(arr, tempArray, lowIdx, midIdx, comp)
	slice(arr, tempArray, midIdx+1, highIdx, comp)
	merge(arr, tempArray, lowIdx, midIdx, highIdx, comp)
}

func merge(arr []int, tempArray []int, lowIdx int, midIdx int, highIdx int, comp func(int, int) bool) {
	lowerStart := lowIdx
	lowerStop := midIdx
	upperStart := midIdx + 1
	upperStop := highIdx
	cnt := lowIdx
	for lowerStart <= lowerStop && upperStart <= upperStop {
		if comp(arr[lowerStart], arr[upperStart]) == false {
			tempArray[cnt] = arr[lowerStart]
			lowerStart++
		} else {
			tempArray[cnt] = arr[upperStart]
			upperStart++
		}
		cnt++
	}
	for lowerStart <= lowerStop {
		tempArray[cnt] = arr[lowerStart]
		lowerStart++
		cnt++
	}
	for upperStart <= upperStop {
		tempArray[cnt] = arr[upperStart]
		upperStart++
		cnt++
	}
	for i := lowIdx; i <= highIdx; i++ {
		arr[i] = tempArray[i]
	}
}

// QuickSort ...
func QuickSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1, comp)
}

func quickSortUtil(arr []int, lowIdx int, highIdx int, comp func(int, int) bool) {
	if highIdx <= lowIdx {
		return
	}
	pivot := arr[lowIdx]
	start := lowIdx
	stop := highIdx
	for lowIdx < highIdx {
		for comp(arr[lowIdx], pivot) == false && lowIdx < highIdx {
			lowIdx++
		}
		for comp(arr[highIdx], pivot) == true && lowIdx <= highIdx {
			highIdx--
		}
		if lowIdx < highIdx {
			swap(arr, highIdx, lowIdx)
		}
	}
	swap(arr, highIdx, start)
	quickSortUtil(arr, start, highIdx-1, comp)
	quickSortUtil(arr, highIdx+1, stop, comp)
}

// QuickSelect algorithm is used to find the element, which will be at Kth position
// when the list will be sorted without actually sorting the whole list.
func QuickSelect(arr []int, key int) int {
	size := len(arr)
	quickSelectUtil(arr, 0, size-1, key)
	return arr[key]
}

func quickSelectUtil(arr []int, lower int, upper int, key int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper
	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, lower, upper)
		}
	}
	swap(arr, start, upper)
	if start <= key && key <= upper {
		quickSelectUtil(arr, start, upper, key)
	} else {
		quickSelectUtil(arr, upper+1, stop, key)
	}
}

// BucketSort ...
func BucketSort(s []int, lowerRangeBound int, upperRangeBound int) {
	rng := upperRangeBound - lowerRangeBound
	tempArray := make([]int, rng)
	for i := 0; i < rng; i++ {
		tempArray[i] = 0
	}
	for i := 0; i < len(s); i++ {
		tempArray[s[i]-lowerRangeBound]++
	}
	cnt := 0
	for i := 0; i < rng; i++ {
		for ; tempArray[i] > 0; tempArray[i]-- {
			s[cnt] = i + lowerRangeBound
			cnt++
		}
	}
}

func insertionSortS(s []string) {
	var temp string
	var j int
	size := len(s)
	for i := 1; i < size; i++ {
		temp = s[i]
		for j = i; j > 0 && s[j-1] > temp; j-- {
			s[j] = s[j-1]
		}
		s[j] = temp
	}
}

func GeneralizedBucketSort(s []string, lowerRangeBound byte, upperRangeBound byte) {
	rng := int(upperRangeBound - lowerRangeBound)
	tempArray := make([][]string, rng)
	for i := 0; i < rng; i++ {
		tempArray[i] = make([]string, 0)
	}
	var idx int
	for i := 0; i < len(s); i++ {
		idx = int(s[i][0] - lowerRangeBound)
		tempArray[idx] = append(tempArray[idx], s[i])
		insertionSortS(tempArray[idx])
	}
	idx = 0
	for i := 0; i < rng; i++ {
		for ; len(tempArray[i]) > 0; tempArray[i] = tempArray[i][1:] {
			s[idx] = tempArray[i][0]
			idx++
		}
	}
}

// HeapSort ...
func HeapSort(a []int) {
	h := ds.NewHeap(a, true)
	for i := 0; i < len(a); i++ {
		a[i], _ = h.Remove()
	}
}

package algorithms

//Fibonachi ...
func Fibonachi(start int, end int) []int {
	res := make([]int, 0)
	a := 0
	b := 1
	for a < end {
		if a >= start {
			res = append(res, a)
		}
		a, b = b, a+b
	}
	return res
}

package datastructures

//More ...
func More(a int, b int) bool {
	return a > b
}

//Less ...
func Less(a int, b int) bool {
	return a < b
}

//Slice for conversion multiple outputs to choosen one
//
// func prnt() {
// 	fmt.Println(slice(ab)[1])
// }
//
// func ab(a, b int) (int, bool) {
// 	if b < 0 {
// 		return a, true
// 	}
// 	return a, false
// }
func Slice(args ...interface{}) []interface{} {
	return args
}

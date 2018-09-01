package tasks

import (
	"fmt"
)

func TowersOfHanoi(n int) {
	tohUtil(n, "A", "C", "B")
}

func tohUtil(n int, from string, to string, temp string) {
	if n < 1 {
		return
	}
	tohUtil(n-1, from, temp, to)
	fmt.Println("Move disk: ", n, " ", from, " -> ", to)
	tohUtil(n-1, temp, to, from)
}

// Great Common Divisor
func GCD(m int, n int) int {
	if n > m {
		return GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n, m%n)
}

func Factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * Factorial(n-1)
}

//Given N, find N'th number in Fibonacci series
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func BinarySearchRc(ar []int, k int) (int, bool) {
	if len(ar) > 0 {
		return binSearch(ar, k, 0, len(ar)-1)
	} else {
		return -1, false
	}
}

func binSearch(ar []int, k int, low int, high int) (int, bool) {
	mid := low + (high-low)/2
	if ar[mid] == k {
		return mid, true
	}
	if low == high {
		return -1, false
	}
	if ar[mid] < k {
		low = mid + 1
		return binSearch(ar, k, low, high)
	} else {
		high = mid - 1
		return binSearch(ar, k, low, high)
	}
}

func Permutation(data []int, i int, length int) {
	if length == i {
		PrintSlice(data)
		return
	}
	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func PrintSlice(data []int) {
	fmt.Println(data)
}

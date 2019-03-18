package main

import "fmt"

// O(1) > O(n) > O(logn) > O(n^2) > O(2^n)

func main() {
	src := []int{9, 2, 3, 4, 6, 8, 10, 23, 4, 43, 5}
	findTheBigOne(src)
	fmt.Println(splitTwo(src))
}

func findTheBigOne(src []int) {
	fmt.Println("bubble sort")

	bigOne := 0

	fmt.Println(src)

	for j := 0; j < len(src); j++ {
		for i := 0; i < len(src)-j; i++ {
			if src[bigOne] < src[i] {
				bigOne = i
			}
			if i == len(src)-j-1 {
				src[bigOne], src[len(src)-j-1] = src[len(src)-j-1], src[bigOne]
				fmt.Println(src)
			}
		}
	}

	// n + (n-1) + (n-2) + ... + 1 = (n + 1) * n / 2 = O(n^2)
}

// 递归：
// 1. 确定递归方程
// 2. 确定终止条件

// 分为n个只有两个数的小组，分别找出对应的大的和小的，大的放一起，小的放一起，变成两个数组
// 一直计算下去，直到，大的只有一个数，小的也只有一个数
func splitTwo(src []int) (sorted []int) {

	if len(src) == 1 {
		return src
	}

	if len(src) == 2 {
		if src[0] < src[1] {
			return src
		} else {
			src[0], src[1] = src[1], src[0]
			return src
		}
	}

	bigOne := 0

	for i := 0; i < len(src)/2; i++ {
		for j := 0; j < len(src)-i; j++ {
			if src[j] > src[bigOne] {
				bigOne = j
			}
			if j == len(src)-1 {
				src[bigOne], src[len(src)-i-1] = src[len(src)-i-1], src[bigOne]
			}
		}
	}

	fmt.Println("small", src[:len(src)-len(src)/2], "big", src[len(src)-len(src)/2:])

	return append(splitTwo(src[:len(src)-len(src)/2]), splitTwo(src[len(src)-len(src)/2:])...)

	// 递归深度为：log2n
	// 1 * (n + 1) * n + 2 * (n/2 + 1) * n/2 + ... + n/4 * (2 + 1) * 2 + n/2 * 1
	// T(n) = T(n/2) + O(n^2)
	//      = T(n/4) + O(n^2/4) + O(n^2)
	//      = O(n^2) + ... + O(n^2)
	//      = O(log2n * n^2)
}

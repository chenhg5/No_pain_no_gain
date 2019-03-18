package main

import "fmt"

func main() {

	// 1.求数组中的最大数
	// 2.1+2+3+...+n
	// 3.求n个整数的积
	// 4.求n个整数的平均值
	// 5.求n个自然数的最大公约数与最小公倍数
	// 6.有一对雌雄兔，每两个月就繁殖雌雄各一对兔子.问n个月后共有多少对兔子
	// 7.已知：数列1,1,2,4,7,13,24,44,...求数列的第n项

	// 已知有三根针分别用A, B, C表示，在A中从上到下依次放n个从小到大的盘子，现要求把所有的盘子
	// 从A针全部移到B针，移动规则是：可以使用C临时存放盘子，每次只能移动一块盘子，而且每根针上
	// 不能出现大盘压小盘，找出移动次数最小的方案.

	// 楼梯有n阶台阶，上楼可以一步上1阶，也可以一步上2阶，编一程序计算共有多少种不同的走法.

	// 递归与迭代

	src := []int{9, 2, 3, 4, 6, 8, 10, 100, 23, 4, 43, 5}
	fmt.Println(getTheMaxOne(src, 0))
	fmt.Println(cumulativeAdd(10))

}

func getTheMaxOne(src []int, max int) int {
	if len(src) == 1 {
		if src[0] > max {
			return src[0]
		} else {
			return max
		}
	}

	if src[0] > max {
		return getTheMaxOne(src[1:], src[0])
	} else {
		return getTheMaxOne(src[1:], max)
	}
}

func cumulativeAdd(num int) int {
	if num == 1 {
		return 1
	}
	return num + cumulativeAdd(num - 1)
}

func cumulativeMultiply(num int) int {
	if num == 1 {
		return 1
	}
	return num * cumulativeAdd(num - 1)
}
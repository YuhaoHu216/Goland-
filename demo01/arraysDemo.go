package main

import "fmt"

func main() {
	var a [5]int
	// 不初始化,默认零值,int的零值就是0
	fmt.Println("emp:", a)
	// 可将数组中的值拿出来单独操作
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// 使用内置len函数获取数组的长度
	fmt.Println("len:", len(a))
	// 创建时可初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println("2d: ", twoD)
}

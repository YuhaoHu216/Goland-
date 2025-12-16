package main

import "fmt"

func main() {
	i := 1
	// 类似于其他语言的while
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	// 正常的for循环
	for j := 10; j > 0; j-- {
		fmt.Println(j)
	}
	// 死循环
	for {
		fmt.Println("loop")
		break
	}

}

package main

import "fmt"

func main() {
	i := 8
	if i%2 == 0 {
		fmt.Println("i is even")
	} else {
		fmt.Println("i is odd")
	}
	// 条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}

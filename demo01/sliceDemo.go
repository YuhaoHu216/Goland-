package main

import "fmt"

func main() {
	// 使用内建的方法 make。这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// 操纵和数组类似
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	// slice可复制,但是必须相同长度
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 这里包含s[2],s[3],s[4]三个元素
	l := s[2:5]
	fmt.Println("sl1:", l)
	// 从s[0]包含到s[5]
	l = s[:5]
	fmt.Println("sl2:", l)
	// 从s[2]到最后一个值
	l = s[2:]
	fmt.Println("sl3:", l)
	// 初始化时赋值
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// 和组成多维数组结构,内部slice长度可以不同,和多维数组不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

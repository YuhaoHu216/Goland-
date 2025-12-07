package main

import "fmt"

/*
*
author by hyh
*/
func main() {
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2025-12-07"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	fmt.Printf(url, stockcode, enddate)
}

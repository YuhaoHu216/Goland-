package main

import "fmt"

/*
常规的通过通道发送和接收数据是阻塞的。
然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
*/
// 通俗讲就是select会马上检查所有case,如果有可以执行的就执行一个,都不能执行就走default的逻辑,不会等待
func main() {
	// 都是无缓冲的channel,当前程序中没有任何goroutine向其发送数据,所有接收和发送都会阻塞
	messages := make(chan string)
	signals := make(chan bool)

	// 这里是一个非阻塞 接收 的例子。如果在 messages 中存在值，然后 select 将这个值带入 <-messages case中。
	// 如果不存在，就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞 发送 的实现方法和上面一样。
	// 如果没有default就会一直卡(阻塞)在messages <- msg
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	// 这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作。
	// 这里如果msg := <-messages和sig := <-signals都可执行的话就会随机选择一个执行
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

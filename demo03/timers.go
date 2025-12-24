package main

import "time"
import "fmt"

// 我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。Go 的内置 定时器 和 打点器 特性让这些很容易实现。
func main() {

	// 基础定时器(阻塞等待)
	// 创建一个定时器timer1,两秒后触发,内部有一个channel
	timer1 := time.NewTimer(time.Second * 2)
	// <-timer.C 阻塞当前的goroutine(main) 直到两秒后定时器到期 定时器到期时runtime向timer.C发送一个time.time
	// main goroutine被唤醒 相当于time.Sleep(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 可取消的定时器
	// 这段代码实际执行顺序 创建timer2 启动goroutine但是没等到1s 立刻调用timer2.Stop() stop2等于true
	// 定时器被停止,timer2.C永远不会再收到值,goroutine被阻塞 程序结束goroutine被回收
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

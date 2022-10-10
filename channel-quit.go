// @Title  channel-quit
// @Description  使用无缓冲通道阻塞主线，等待goroutine结束
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-10 15:15
package main

import (
	"fmt"
)

// @title    main
// @description   使用无缓冲通道阻塞主线，等待goroutine结束
// @auth      MGAronya（张健）             2022-10-10 15:15
// @param     void
// @return    void
func main() {
	ch, quit := make(chan int), make(chan int)
	go func() {
		// TODO 添加数据
		ch <- 8
		// TODO 发送完成信号
		quit <- 1
	}()
	for isQuit := false; !isQuit; {
		// TODO 监视通道ch的数据输出
		select {
		case v := <-ch:
			fmt.Printf("received %d from ch", v)
		case <-quit:
			isQuit = true
		}
	}
}

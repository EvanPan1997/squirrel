package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// 创建一个整数类型的通道
	ch := make(chan int)
	// 启动一个goroutine，向通道发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // 关闭通道
	}()

	// 使用select语句从通道接收数据
	for v := range ch {
		fmt.Println(v)
		select {
		case value, ok := <-ch:
			if ok {
				fmt.Println("接收到数据：", value)
			} else {
				fmt.Println("通道已关闭")
				break
			}
		default:
			fmt.Println("没有数据可接收")
			time.Sleep(time.Second)
		}
	}
}

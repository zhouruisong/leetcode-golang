package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	l := rate.NewLimiter(2, 5)
	ctx := context.Background()
	start := time.Now()
	// 要处理二十个事件
	for i := 0; i < 20; i++ {
		l.Wait(ctx)
		// dosomething
	}
	fmt.Println(time.Since(start)) // output: 7.501262697s （初始桶内5个和每秒2个token）
}

package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

var mp sync.Map

type TimerInfo struct {
	ResetCh chan int
	Cancel  context.CancelFunc
}

type dosomething func()

func startTimer(id int, timeSecond time.Duration, f dosomething) {
	ctx, cancel := context.WithCancel(context.Background())
	resetCh := make(chan int)

	tinfo := TimerInfo{
		ResetCh: resetCh,
		Cancel:  cancel,
	}
	mp.Store(id, tinfo)

	//log.Infof("timeSecond:%v", timeSecond.Seconds())
	go func(ctx context.Context, timeSecond time.Duration, resetCh chan int) {
		defer func() {
			if v := recover(); v != nil {
				fmt.Errorf("panic,err:%+v,uuid:%v\n", v, debug.Stack())
			}
		}()

		fmt.Printf("====start timer====\n")
		//执行一次,多次执行需要执行reset
		timer := time.NewTimer(timeSecond)

		go func() {
			time.Sleep(8)
			timer.Stop()

			fmt.Printf("====stop timer====\n")
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("timeSecond:%v stop idleTimer\n", timeSecond)
				return
			case <-resetCh:
				timer.Reset(timeSecond)
			case <-timer.C:
				fmt.Printf("timeSecond:%v do f\n", timeSecond)
				f()
				fmt.Printf("goroutine exit\n")
				return
			}
			fmt.Printf("===end timer====\n")
			return
		}
	}(ctx, timeSecond, resetCh)
}

func dofmt() {
	fmt.Printf("=====do====\n")
}
func main() {
	go startTimer(1, 10*time.Second, dofmt)
	for true {
		time.Sleep(1 * time.Second)
		fmt.Printf("===1====\n")
	}
}

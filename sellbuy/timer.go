package main

import (
	"fmt"
	"time"
)

//定时启动
func runAtTime(ts time.Time) {
	//计算时差
	sl := ts.Sub(time.Now())
	fmt.Println("begin to sleep ", sl)
	time.Sleep(sl) //第一次通过睡眠触发
	getMatureAution()
	tk := time.NewTicker(24 * 60 * 60 * time.Second) //定义定时器
	//tk := time.NewTicker(1 * 1 * 30 * time.Second) //定义定时器
	for {
		tm := <-tk.C //阻塞等待tk信号
		fmt.Println(tm)
		getMatureAution()
	}

}

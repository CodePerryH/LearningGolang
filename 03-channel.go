package main

import (
	"fmt"
	"time"
)
/*
管道
1.channel 是一种数据类型
2.无缓存channel具有同步2个go线程的作用，channel的写入和读取 具有阻塞线程的能力
3.有缓存channel 不会阻塞goroutine（但是当channel已经满了，再向里面写数据，就会阻塞）
*/
func main() {
	//定义一个无缓存channel
	c := make(chan  int)
	 go func() {

	 	///defer的作用 是将 后面的执行语句 压到当前函数栈的栈底 当函数调用栈出栈到栈底再执行
	 	defer  fmt.Println("go routine 结束")

	 	fmt.Println("goroutine 正在运行")
	 	//将666 发给c
	 	c <- 666
	 }()
	//从c中接受数据，并赋值给num
	num := <-c
	fmt.Println("num =",num)
	fmt.Println("main goroutie 结束...")

	haveSaveChannel()
}

///有缓存channe 演示

func haveSaveChannel(){
	c := make(chan  int ,3) //带有3个容量的缓存的channel

	fmt.Println("len c = ",len(c),"cap(c)",cap(c))

	go func() {
		defer  fmt.Println("子go结束")

		i := 0
		for  {
			i++
			c <- i
			fmt.Println("子go程正在运行，发送的元素 =",i,"len(c) = ",len(c),"cap(c) =",cap(c))
			if i >= 3 {
				break
			}
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3 ; i ++{
		 num := <-c
		 fmt.Println("取出channel的缓存依次是num = ",num)
	}

	fmt.Println("main 结束")
}

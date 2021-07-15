package main

import (
	"fmt"
	"runtime"
	"time"
)

//协程调度器
//GMP  G:goroutine协程 M:thrend线程 P:processor处理器
//调度器的设计策略
//1.复用线程 （work stealing 机制、hand off机制） 2.利用并行（GOMAXPROCS 限定p的个数 = cpu核数/2） 3.抢占 4.全局G队列
func main() {

	///1.创建一个go线程 去执行newTask()
	//go newTask()
	//
	//i := 0
	//for  {
	//	i++
	//	fmt.Printf("main Goroutine : 1 = %d\n",i)
	//
	//	time.Sleep(1 * time.Second)
	//
	//}

	///2.用go创建一个承载一个形参为空，返回值为空的函数
	go func(){
		defer  fmt.Println("A.defer")
		func(){
			defer  fmt.Println("b.defer")
			fmt.Println("b")

			//这句话表示跳出这个2个函数
			runtime.Goexit()
		}()
		fmt.Println("A")
	}()

	///3.go创建一个有参数的函数
	go func(a int,b int) bool {
		fmt.Println("a = ",a,"b =",b)
		return  true
	}(20,10)
}

func newTask(){
	i := 0
	for  {
		i++
		fmt.Printf("new Goroutine : 1 = %d\n",i)

		time.Sleep(1 * time.Second)

	}
}
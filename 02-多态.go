package main

import "fmt"

//interface 本质是一个指针 , 感觉和协议差不多的样子
/*
多态 ： 子类以不同的方式展示同一条指令
1.有一个父类（接口类）
2.有子类（实现了父类的全部接口）
3.父类类型的变量（指针）指向（引用）子类的数据变量
*/
type Animal interface {
	sleep()
	GetColor() string
	GetType() string
}
//struct
type Cat struct {
	color string

}

func (this Cat) sleep() {
	panic("implement me")
}

func  (this Cat)Sleep()  {
	fmt.Println("cat is sleep")
}
func  (this Cat)GetColor() string  {
	return  this.color
}
func  (this Cat)GetType() string{
	return  "cat"
}
//stuct
type Dag struct {
	color string
}
func  (this Dag) sleep()  {
	fmt.Println("Dag is sleep")
}
func  (this Dag)GetColor() string  {
	return  this.color
}
func  (this Dag)GetType() string{
	return  "Dag"
}
func main() {
	cat := Cat{color: "red"}
	isWhoSleep(&cat)

	dag := Dag{ "block"}
	isWhoSleep(&dag)

}

///
func isWhoSleep(animal Animal) {
	animal.sleep()
}
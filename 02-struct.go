package main

import "fmt"
/*
1.go 中的类 属性 方法 首字母大写 表示外包可以访问，小写表示只能在本包中使用

2. go语言的继承不存在 公有 私有 属性 继承子类的
*/

// Hero /*
type Hero struct {
	Name string
	Ad int
}

//给结构体绑定方法
func (this *Hero) GetName()  string {
	fmt.Println("name -=",this.Name)
	return this.Name
}
func  (this *Hero) SetName(nawName string)  {
	this.Name = nawName
}
func (this *Hero) eat(){
	fmt.Println("吃饭了")
}

/*这里表示superman 继承于Hero*/
type SuperMan struct {
	Hero
	lecel int
}
///重写父类
func (this *SuperMan) eat(){
	fmt.Println("SuperMan吃饭了")

}
func (this *SuperMan) run(){
	fmt.Println("SuperMan奔跑吧")

}

func main() {
	//创建
	hero := Hero{Name: "xiaowang",Ad: 100}

	fmt.Println("hero is =",hero)

	hero.SetName("ni da ye ")

	fmt.Println("修改后 =",hero.GetName(),"2. = ",hero.Name)

	hero.eat()

	///继承演示
	//superMan := SuperMan{{Name: "man",Ad: 70},99}
	var superMan SuperMan
	superMan.Name = "man"
	superMan.Ad = 19
	superMan.lecel = 100


	superMan.eat()

	superMan.run()


}



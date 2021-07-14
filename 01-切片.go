package main

import (
	"fmt"
	"strings"
)

func main() {
	///切片是go中的一个数据结构 和数组的区别是 数组容量固定了 ，而切片可以宽容
	//切片有3个成员变量(low,hight,cap),low为初始值，cap为容量
	//切片做函数参数 传递的是引用，数组是值传递
	//切片的创建

	//1.自动推到
	s1 := []int{1,2,3,4,5}
	s1 = append(s1,888)
	fmt.Println("s1 = ",s1)

	//2.指定长度 和容量
	//make([]int,长度，容量)

	s2 := make([] int,5,10)
	fmt.Println("s2 = ",s2)

	//3.初始化指定长度的切片，长度 = 容量
	//make([]int,长度)
	s3 := make([]int,7)
	fmt.Println("s3 = ",s3)

	///切片的遍历
	strSlicn := []string{"s","","c","f","","t","c","s","f","t","t"}

	fmt.Println("out = ",noEmpty(strSlicn))

	fmt.Println("out = ",noSame(strSlicn))

	///切片的copy
	//copy(目标位置切片,源切片)
	copydata := [...]int{5,6,7,8,9}

	s11 := copydata[1:5] //{6,7,8,9}
	s22 := copydata[0:2] //{5,6}

	fmt.Println("s11 = ",s11,"s22 = ",s22)

	copy(s11,s22)

	fmt.Println("s22 = ",s11)

    //map 与 切片的练习
    juzi := "I  love you  go lang  , I world"

    mp := stringUncdShowNumber(juzi)

    fmt.Println("字符串单词出现的次数=\n",mp)

}

func noEmpty(data []string) []string {
	//在strSlicn上截取一个长度为0的切片 等价于make（【】string，0）

	out := data[:0]
	for _,str := range data{
		if str != "" {
			out = append(out,str)
		}
	}
	return  out
}
func noEmpty2(data []string) []string {

	i := 0
	for _,str := range data{
		if str != "" {
			data[i] = str
			i++
		}
	}
	return  data[:i]
}
///去出切片的重复元素
func noSame(data []string) []string {
	noSm := data[:0]
	for _,str := range data{
		bl := false
		for _,str2 := range noSm {
			if str2 == str  {
				bl = true
				break
			}
		}
		if !bl {
			noSm = append(noSm,str)
		}

	}
	return noSm
}

///计算字符串单词出现的次数
func stringUncdShowNumber(str string) map[string]int{
	mp := make(map[string]int)
	///拆分字符串 strings.Fields() 按空格 返回的是一个string切片
	for _,str1 := range strings.Fields(str) {
		///判断当前map中是否包含str1这个key
		if _,ok := mp[str1];ok {
			mp[str1] += 1
		} else  {
			mp[str1] = 1
		}
	}
	return  mp
}
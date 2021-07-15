package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type resume struct {
	name string `info"name" doc:"我的名字"`
	sex string `info"sex"`
}

func findTag(str interface{})  {
	t := reflect.TypeOf(str).Elem()
	for i := 0;i < t.NumField(); i++ {
		tageInfo := t.Field(i).Tag.Get("info")
		tagedoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info = ",tageInfo,"doc = ",tagedoc)

	}
}
func main() {
	var re resume
	//通过reflect 库的方法得到结构体设置的标签
	findTag(&re)

	//json转

	movie := Movie{"喜剧之王","2020",30,[]string{"x","y"}}
	//编码的过程 结构体 -》 json
	jsonstr,err := json.Marshal(movie)
	if err != nil{
		fmt.Println("json marshal error",err)
		return
	}
	fmt.Println("jsonstr =",jsonstr)

	//json -》 结构体
	mymovie := Movie{}
	err = json.Unmarshal(jsonstr,&mymovie)
	if err != nil {
		fmt.Println("出差了")
		return
	}
	fmt.Println("得到的结构体 = ",mymovie)

}


	/*1.列子 结构体转json*/
	type Movie struct {
		Title string `json:"title"`
		Year string `json:"year"`
		Price int `json:"rmb"`
		Actors []string `json:"actors"`
	}
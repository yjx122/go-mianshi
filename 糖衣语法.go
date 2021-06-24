package main

import "fmt"

func sugar(values...string){
	for _,v:=range values{
		fmt.Println("v:",v)
	}
}
func main() {
	//测试语法糖
	sugar("a","b","c")//可传递多个字符串string
}

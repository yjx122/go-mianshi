package main

import "fmt"
type Behe interface{
	Eat()string
}
type Dog struct{
name string
age int
}
type Cat struct{
	name string
	age int
}
func Dog1() {
	/*//方法1
		var dog Dog
		dog.name="xiaogou"
		dog.age=1
	fmt.Println(dog)*/
	/*//方法2
	  dogg:=Dog{name:"xioagou",age:1,}
	  	fmt.Println(dogg)
	  }*/
	//方法3
	dog := new(Dog)
	dog.name = "xiaogou"
	dog.age = 1
	cat:= new(Cat)
	cat.name = "xiaogoue"
	cat.age = 2
}
func (d *Dog)Eat()string{
	fmt.Println("yumy yumy")
	return "kt"
}
func (d *Cat)Eat()string{
	fmt.Println("yumy yumy1")
	return "dt"
}

func main() {
   Dog1()
	//测试结构定义变量
	var b Behe
	b=new(Dog)
	b.Eat()
	b=new(Cat)
	b.Eat()

}

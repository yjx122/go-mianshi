package main

import "fmt"

func sslice(){
	slice:=make([]string,3,4)
	slice[0]="kldog"
	slice[1]="cat"
	fmt.Println(slice)
}
/*func panic1(){
	defer func(){
		message:=recover()
		fmt.Println("panic message:",message)
	}()//匿名函数
	panic(1)

}*/
func panic1(){
	defer coverpainc()
	panic(1)

}
func coverpainc(){
	message:=recover()
	fmt.Println("panic message:",message)
}


func main() {
	panic1()
	sslice()

}

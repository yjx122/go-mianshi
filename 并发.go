package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var chanInt chan int=make(chan int,10)//定义一个chan类型为int
var timeout chan bool=make(chan bool,10)//定义一个chan类型为int
var WG sync.WaitGroup
func loop(){
	for i:=0;i<11;i++{
		fmt.Println(i)
	}
}
func Send(){
	time.Sleep(time.Second*1)
	chanInt<-1
	time.Sleep(time.Second*1)
	chanInt<-2
	time.Sleep(time.Second*1)
	chanInt<-3
	time.Sleep(time.Second*1)
	timeout<-true
}
func Receive(){

	time.Sleep(time.Second*1)
	num:=<-chanInt
	fmt.Println(num)
	time.Sleep(time.Second*1)
	num1:=<-chanInt
	time.Sleep(time.Second*1)
	fmt.Println(num1)
	num2:=<-chanInt
	time.Sleep(time.Second*1)
	fmt.Println(num2)


	/*for {
		select{
		case num:=<-chanInt:
		fmt.Println(num)
		case <-timeout:
			fmt.Println("timeout")
		}
	}
	*/

}
//读取数据
func Read(){
	for i:=0;i<3;i++{
		WG.Add(1)

	}
}
//写入数据
func Wrinte(){
	for i:=0;i<3;i++{
		WG.Done()
		fmt.Println(i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go loop()
	go loop()
	time.Sleep(1*time.Second)
	Send()
	//Receive()
	go Read()
	go Wrinte()
	WG.Wait()
	fmt.Println("wancheng")
}
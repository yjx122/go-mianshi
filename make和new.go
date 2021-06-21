package main

import "fmt"

func sslice(){
	slice:=make([]string,3,4)
	slice[0]="kldog"
	slice[1]="cat"
	fmt.Println(slice)
}
func mmap(){
	va:=make(map[int]string)
	vaa:=new(map[int]string)

	va[0]="cat"
	va[100]="dog"
	delete(va,0)
	fmt.Println(va,vaa)
}
func cchan(){
	mchan:=make(chan int)
	 close(mchan)


}

func sppendd(){
	s1:=make([]int,2)
	s1[0]=1
	s1[1]=2
	s2:=make([]int,3)
	s2[0]=3
	s2[1]=4
	s2[2]=5
//s4:=append(s1,3)
copy(s1,s2)

fmt.Println(s1)

}
func main() {
sslice()
mmap()
cchan()
	sppendd()
}

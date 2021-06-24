package main

import (
	"encoding/json"
	"fmt"
)

//序列化结构体
type Sever struct{
SeverName string  `json:"name"`
SeverIp string `json:"ip"`
Severpprt int 		`json:"port"`
}

func Serialize(){
	sever:=new(Sever)//返回传入类型的指针地址
	sever.SeverName="Demo-for-json"
	sever.SeverIp="127.0.0.1"
	sever.Severpprt=8080
	fmt.Println(sever)
	b,err:=json.Marshal(sever)
	if err!=nil{
		fmt.Println("Marshal err:",err.Error())
		return
	}
fmt.Println("Marshal json:",string(b))//将json字节数组转化为string
}
//序列化Map
func SerialMap(){
	sever:=make(map[string]interface{})
	sever["SeverName"]="Demo"
	sever["SeverIp"]="127"
	sever["Severpprt"]=9090
	b,err:=json.Marshal(sever)
	if err!=nil{
		fmt.Println("Marshal err:",err.Error())
		return
	}
	fmt.Println("Marshal json:",string(b))//将json字节数组转化为string

}
//反序列化结构体
func DeserializeStruct(){
	jsonString:=`{"name":"Demo-for-json","ip":"127.0.0.1","port":8080}`
server:=new(Sever)
err:=json.Unmarshal([]byte(jsonString),&server)
	if err!=nil{
		fmt.Println("Marshal err:",err.Error())
		return
	}
	fmt.Println(server)
}
//反序列化Map
func DesreialzeMap(){
	jsonString:=`{"name":"Demo-for-json","ip":"127.0.0.1","port":8080}`
	server:=make(map[string]interface{})
	err:=json.Unmarshal([]byte(jsonString),&server)//为什么要取地址，读取server的地址，并修改它
	if err!=nil{
		fmt.Println("Marshal err:",err.Error())
		return
	}
	fmt.Println(server)
}
func main() {
	//Serialize()
	//SerialMap()
	DeserializeStruct()
	DesreialzeMap()

}

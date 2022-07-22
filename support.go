package main

//定义接口
type Support interface {
	SetNext(next Support) //参数不确定，所以这里使用接口
	HandleTrouble(trouble *Trouble) bool
}

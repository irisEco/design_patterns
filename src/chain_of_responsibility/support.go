package chain_of_responsibility

import (
	"fmt"
	"strconv"
)

//定义接口
type Next interface {
	SetNext(next Next)
	HandleTrouble(trouble *Trouble)
	Done(trouble *Trouble)
	Fail(trouble *Trouble)
	Resolve(trouble *Trouble) bool
}
type S struct{}

func Resolve(t *Trouble) bool {
	return false
}
func Done(t *Trouble) {
	fmt.Println(strconv.Itoa(t.Number) + "is resolved by" + strconv.Itoa(t.Number) + ".")
}

func Fail(t *Trouble) {
	fmt.Println(strconv.Itoa(t.Number) + "cannot be resolved.")
}

//1.已经解决或者失败
//2.未解决且有下一个方法，则丢给下一个方法直至1
//此处可用泛型动态绑定函数；目前版本为1.17暂无好的解决方法
func (s *S) HandleTrouble(t *Trouble) {}

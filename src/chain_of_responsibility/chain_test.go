package chain_of_responsibility

import (
	"fmt"
	"testing"
)

func Test_HandleTrouble_In_Chain(t *testing.T) {
	//这里没有实际的顺序关系，在实际应用中应该有顺序关系
	spe := &SpecialSupport{Name: "A", Number: 1}
	lim := &LimitSupport{Name: "B", Limit: 124}
	odd := &OddSupport{Name: "B"}
	//设置下一个责任节点
	spe.SetNext(lim)
	//设置lim下一个责任节点
	lim.SetNext(odd)

	for i := 0; i < 100; i += 33 {
		event := &Trouble{Number: i}
		fmt.Println(event.ToString())
		spe.HandleTrouble(event)
	}
}

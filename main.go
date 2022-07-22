package main

import (
	"design_patterns/src/chain_of_responsibility"
	"fmt"
)

//制作support的职责链，制造问题并测试程序行为

func main() {
	//这里没有实际的顺序关系，在实际应用中应该有顺序关系
	spe := &chain_of_responsibility.SpecialSupport{Name: "A", Number: 1}
	lim := &chain_of_responsibility.LimitSipport{Name: "B", Limit: 124}
	odd := &chain_of_responsibility.OddSupport{Name: "B"}
	//设置下一个责任节点
	spe.SetNext(lim)
	//设置lim下一个责任节点
	lim.SetNext(odd)

	for i := 0; i < 100; i += 33 {
		event := &chain_of_responsibility.Trouble{Number: i}
		fmt.Println(event.ToString())
		spe.HandleTrouble(event)
	}

}

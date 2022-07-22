package main

import (
	chain "design_patterns/src/chain_of_responsibility"
	"fmt"
	"math/rand"
)

//制作Next的职责链，制造问题并测试程序行为

func main() {
	//这里没有实际的顺序关系，在实际应用中应该有顺序关系
	spe := &chain.SpecialSupport{Name: "SpecialSupport", Number: rand.Intn(180)}
	lim := &chain.LimitSupport{Name: "LimitSupport", Limit: rand.Intn(180)}
	odd := &chain.OddSupport{Name: "OddSupport"}
	ons := &chain.NoSupport{Name: "NoSupport"}
	println(lim.Limit)

	//设置下一个责任节点
	spe.SetNext(lim)
	//设置lim下一个责任节点
	lim.SetNext(odd)
	odd.SetNext(ons)
	odd.SetNext(ons)
	for i := rand.Intn(rand.Intn(150)); i <= 100; i += rand.Intn(10) {
		event := &chain.Trouble{Number: i}
		fmt.Println(event.ToString())
		spe.HandleTrouble(event)
	}

}

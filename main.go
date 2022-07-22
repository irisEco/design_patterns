package main

import "fmt"

//制作support的职责链，制造问题并测试程序行为

func main() {

	spe := &SpecialSupport{name: "A", number: 1}
	lim := &LimitSipport{name: "B", limit: 124}
	odd := &OddSupport{name: "B"}
	//设置下一个责任节点
	spe.SetNext(lim)
	//设置下一个责任节点
	lim.SetNext(odd)

	for i := 0; i < 100; i += 33 {
		event := &Trouble{number: i}
		fmt.Println(event.toString())
		spe.HandleTrouble(event)
	}

}

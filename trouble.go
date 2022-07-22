package main

import "strconv"

//表示发生问题的类
type Trouble struct {
	number int //问题编号
}

func (t *Trouble) newTrouble(num int) {
	t.number = num
}
func (t *Trouble) getNumber() int {
	return t.number
}
func (t *Trouble) toString() string {
	return "[Trouble: " + strconv.Itoa(t.number) + "]"
}

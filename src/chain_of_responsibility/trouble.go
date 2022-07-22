package chain_of_responsibility

import "strconv"

//表示发生问题的类
type Trouble struct {
	Number int //问题编号
}

func (t *Trouble) NewTrouble(num int) {
	t.Number = num
}
func (t *Trouble) GetNumber() int {
	return t.Number
}
func (t *Trouble) ToString() string {
	return "[Trouble: " + strconv.Itoa(t.Number) + "]"
}

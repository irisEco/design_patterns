package chain_of_responsibility

import "fmt"

//用来解决问题的具体类（仅解决编号小于指定编号的问题）
type LimitSipport struct {
	Name  string
	Limit int
	Support
}

func (r *LimitSipport) HandleTrouble(t *Trouble) bool {
	if t.GetNumber() < r.Limit {
		fmt.Println("LimitSipport handler is succes")
		return true
	} else {
		fmt.Println("LimitSipport handler is fail")
		r.Support.HandleTrouble(t)
		return false
	}
}

func (s *LimitSipport) SetNext(next Support) {
	s.Support = next
}

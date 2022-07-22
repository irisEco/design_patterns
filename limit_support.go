package main

import "fmt"

//用来解决问题的具体类（仅解决编号小于指定编号的问题）
type LimitSipport struct {
	name  string
	limit int
	Support
}

func (r *LimitSipport) HandleTrouble(t *Trouble) bool {
	if t.getNumber() < r.limit {
		fmt.Println("LimitSipport handler is succes")
		return true
	} else {
		fmt.Println("LimitSipport handler is fail")
		return false
	}
}

func (s *LimitSipport) SetNext(next Support) {
	s.Support = next
}

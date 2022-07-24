package chain_of_responsibility

import (
	"fmt"
	"strconv"
)

//用来解决问题的具体类（仅解决编号小于指定编号的问题）
type LimitSupport struct {
	Name  string
	Limit int
	Next
}

func (l *LimitSupport) Resolve(t *Trouble) bool {
	if t.GetNumber() < l.Limit {
		//fmt.Println("LimitSupport handler is success")
		return true
	} else {
		//fmt.Println("LimitSupport handler is fail")
		return false
	}
}

func (l *LimitSupport) SetNext(next Next) {
	l.Next = next
}

func (l *LimitSupport) Done(t *Trouble) {
	fmt.Println("Success: " + strconv.Itoa(t.Number) + " is resolved by " + string(l.Name) + ".")
}

func (l *LimitSupport) Fail(t *Trouble) {
	fmt.Println("Fail: " + strconv.Itoa(t.Number) + " cannot be resolved by " + string(l.Name))
}

func (s *LimitSupport) HandleTrouble(t *Trouble) {
	if s.Resolve(t) {
		s.Done(t)
	} else if s.Next != nil {
		s.Next.HandleTrouble(t)
	} else {
		s.Fail(t)
	}
}

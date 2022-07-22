package chain_of_responsibility

import (
	"fmt"
	"strconv"
)

//用来解决问题的具体类（永远不处理问题）
type NoSupport struct {
	Name string
	Next
}

// 不处理问题
func (n *NoSupport) Resolve(t *Trouble) bool {
	return false
}
func (n *NoSupport) SetNext(next Next) {
	n.Next = next
}
func (n *NoSupport) Done(t *Trouble) {
	fmt.Println("Success: " + strconv.Itoa(t.Number) + " is resolved by " + string(n.Name) + ".")
}

func (n *NoSupport) Fail(t *Trouble) {
	fmt.Println("Fail: " + strconv.Itoa(t.Number) + " cannot be resolved by " + string(n.Name))
}
func (s *NoSupport) HandleTrouble(t *Trouble) {
	if s.Resolve(t) {
		s.Done(t)
	} else if s.Next != nil {
		s.Next.HandleTrouble(t)
	} else {
		s.Fail(t)
	}
}

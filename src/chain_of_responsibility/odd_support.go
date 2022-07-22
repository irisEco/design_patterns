package chain_of_responsibility

import (
	"fmt"
	"strconv"
)

//用来解决问题的具体类（仅解决奇数编号的问题）
type OddSupport struct {
	Next
	Name string
}

func (o *OddSupport) Resolve(t *Trouble) bool {
	if t.GetNumber()%2 == 1 {
		return true
	} else {
		return false
	}
}
func (o *OddSupport) SetNext(next Next) {
	o.Next = next
}
func (o *OddSupport) Done(t *Trouble) {
	fmt.Println("Success: " + strconv.Itoa(t.Number) + " is resolved by " + string(o.Name) + ".")
}

func (o *OddSupport) Fail(t *Trouble) {
	fmt.Println("Fail: " + strconv.Itoa(t.Number) + " cannot be resolved by " + string(o.Name))
}
func (s *OddSupport) HandleTrouble(t *Trouble) {
	if s.Resolve(t) {
		s.Done(t)
	} else if s.Next != nil {
		s.Next.HandleTrouble(t)
	} else {
		s.Fail(t)
	}
}

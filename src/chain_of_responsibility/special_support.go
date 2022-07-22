package chain_of_responsibility

import (
	"fmt"
	"strconv"
)

//用来解决问题的具体类（仅解决指定编号的问题）
type SpecialSupport struct {
	Name   string
	Number int
	Next
}

func (s *SpecialSupport) Resolve(t *Trouble) bool {
	fmt.Println("handler get Number:", t.GetNumber(), s.Number)
	if s.Number == t.GetNumber() {
		return true
	} else {
		return false
	}
}
func (s *SpecialSupport) SetNext(next Next) {
	s.Next = next
}
func (s *SpecialSupport) Done(t *Trouble) {
	fmt.Println("Succeed: " + strconv.Itoa(t.Number) + " is resolved by " + string(s.Name) + ".")
}

func (s *SpecialSupport) Fail(t *Trouble) {
	fmt.Println("Fail: " + strconv.Itoa(t.Number) + " cannot be resolved by " + string(s.Name))
}
func (s *SpecialSupport) HandleTrouble(t *Trouble) {
	if s.Resolve(t) {
		s.Done(t)
	} else if s.Next != nil {
		s.Next.HandleTrouble(t)
	} else {
		s.Fail(t)
	}
}

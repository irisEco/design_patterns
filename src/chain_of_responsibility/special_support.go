package chain_of_responsibility

import "fmt"

//用来解决问题的具体类（仅解决指定编号的问题）
type SpecialSupport struct {
	Name   string
	Number int
	Support
}

func (s *SpecialSupport) HandleTrouble(t *Trouble) bool {
	fmt.Println("handler get Number:", t.GetNumber())
	fmt.Println(s.Number)
	if s.Number == t.GetNumber() {
		s.Support.HandleTrouble(t)
		fmt.Println("special_support handler is succes")
		return true
	} else {
		fmt.Println("special_support handler is fail")
		s.Support.HandleTrouble(t)
		return false
	}
}
func (s *SpecialSupport) SetNext(next Support) {
	s.Support = next
}

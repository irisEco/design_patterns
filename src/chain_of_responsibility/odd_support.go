package chain_of_responsibility

//用来解决问题的具体类（仅解决奇数编号的问题）
type OddSupport struct {
	Support
	Name string
}

func (o *OddSupport) HandleTrouble(t *Trouble) bool {
	if t.GetNumber()%2 == 1 {
		return true
	} else {
		return false
	}
}
func (s *OddSupport) SetNext(next Support) {
	s.Support = next
}

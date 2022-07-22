package main

//用来解决问题的具体类（永远不处理问题）
type NoSupport struct {
	Support
}

// 不处理问题
func (n *NoSupport) HandleTrouble(t *Trouble) bool {
	n.Support.HandleTrouble(t)
	return false
}
func (s *NoSupport) SetNext(next Support) {
	s.Support = next
}

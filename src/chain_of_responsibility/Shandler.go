package chain_of_responsibility

//暂时废弃类
////
//type SHandler interface {
//	resolve(t *Trouble)
//	done(t *Trouble)
//	fail(t *Trouble)
//	//handler(t *Trouble)
//}
//
//////用来解决问题的抽象类
////type Support struct {
////	name string
////	next *Support
////}
////
//////显示字符串
////func (s *Support) toString() string {
////	return "[" + s.name + "]"
////}
////
////func (s *Support) setName(name string) {
////	s.name = name
////}
////
//////设置要推卸的对象
////func (s *Support) setNext(next Support) {
////	s.next = &next
////}
////
////func resolve(t *Trouble) bool {
////	return false
////}
////func done(t *Trouble) {
////	fmt.Println(string(t.number) + "is resolved by" + string(t.number) + ".")
////}
////
////func fail(t *Trouble) {
////	fmt.Println(string(t.number) + "cannot be resolved.")
////}
////
//////1.已经解决或者失败
//////2.未解决且有下一个方法，则丢给下一个方法直至1
////func (s *Support) handler(t *Trouble) {
////	if resolve(t) {
////		done(t)
////	} else if s.next != nil {
////		s.next.handler(t)
////	} else {
////		fail(t)
////	}
////}

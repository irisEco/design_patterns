package chain_of_responsibility

import "testing"

func Test_HandleTrouble_In_Chain(t *testing.T) {
	oba := &SpecialSupport{name: "A", number: 123}
	obb := &LimitSipport{name: "B", limit: 124}
	oba.SetNext(obb)

	event := &Trouble{number: 1}
	oba.HandleTrouble(event)

	event = &Trouble{number: 2}
	oba.HandleTrouble(event)
}

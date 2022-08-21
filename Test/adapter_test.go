package Test

import "testing"

func TestAdapter(t *testing.T) {
	RunComputer(NewMacAdapter())
	var c Computer = NewMacAdapter()
	c.Buy()
	c.Run()
}

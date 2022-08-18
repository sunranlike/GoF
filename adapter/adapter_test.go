package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var i IBuyComputer = NewWinAdapter()
	i.BuyComputer(1, 2, "i5")

	var i2 IBuyComputer = NewMacAdapter()
	i2.BuyComputer(1, 2, "i5")
}

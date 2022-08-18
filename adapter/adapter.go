package adapter

import "fmt"

type IBuyComputer interface {
	BuyComputer(price int, count int, cpu string)
}
type WinComputer struct {
}

func (WinComputer) BuyWinComputer(price int, count int, cpu string) {
	fmt.Println("购买了win电脑", price, count, cpu)
}

type MacComputer struct {
}

func (MacComputer) BuyMacComputer(price int) {
	fmt.Println("购买了mac电脑", price)
}

type WinAdapter struct {
	computer WinComputer
}

func (w *WinAdapter) BuyComputer(price int, count int, cpu string) {
	w.computer.BuyWinComputer(price, count, cpu)
}

type MacAdapter struct {
	computer MacComputer
}

func (m *MacAdapter) BuyComputer(price int, count int, cpu string) {
	m.computer.BuyMacComputer(price)
}
func NewMacAdapter() *MacAdapter {
	return &MacAdapter{
		computer: MacComputer{},
	}
}
func NewWinAdapter() *WinAdapter {
	return &WinAdapter{
		computer: WinComputer{}}
}

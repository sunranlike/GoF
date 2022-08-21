package Test

import "fmt"

type Mac struct {
	cpu   string
	mem   int
	macOs string
}

func (m *Mac) BuyMacComputer(cpu, macOs string, mem int) error {
	fmt.Println("Buy a mac computer")
	m.cpu = cpu
	m.macOs = macOs
	m.mem = mem
	return nil
}
func (m *Mac) RunMacComputer() error {
	fmt.Println("Running a mac computer")
	return nil
}

type MacAdapter struct {
	Mac
}

func (m *MacAdapter) Buy() error {
	return m.BuyMacComputer("m1", "macOs10", 8)
}
func (m *MacAdapter) Run() error {
	return m.RunMacComputer()
}

func NewMacAdapter() Computer {
	return &MacAdapter{}
}

type Win struct {
	cpu     string
	mem     int
	windows string
}

func (w *Win) BuyWinComputer(cpu, windows string, mem int) error {
	fmt.Println("Buy a Win computer")
	w.cpu = cpu
	w.windows = windows
	w.mem = mem
	return nil
}
func (w *Win) RunWinComputer() error {
	fmt.Println("Running a Win computer")
	return nil
}

type Computer interface {
	Run() error
	Buy() error
}

func RunComputer(c Computer) {
	c.Buy()
	c.Run()
}

package Decorator

import "fmt"

type Book struct{}

func (b Book) Read() {
	fmt.Println("read a book:")
}

type IRead interface {
	Read()
}
type ENBook struct {
	language string
	b        Book
}

func (e ENBook) Read() {
	e.b.Read() //
	fmt.Print("it is a " + e.language + " book")
}
func NewEnbook(l string, b Book) ENBook {
	enbook := ENBook{
		language: l,
		b:        b,
	}
	return enbook

}

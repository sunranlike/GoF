package Decorator

import "testing"

func TestName(t *testing.T) {
	var book Book
	book.Read()
	NewEnbook("en", book).Read()
}

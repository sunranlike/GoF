package Chain

import "testing"

func TestChain(t *testing.T) {
	h := Newhospital()
	d := Newdoctor()
	m := Newmedical()
	h.SetNext(d)
	d.SetNext(m)

	h.Execute()
}

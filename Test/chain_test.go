package Test

import "testing"

func TestChain(t *testing.T) {
	//h := &Hospital{}
	//h.SetNextDepartment(&Doctor{})
	//h.Execute()
	//
	d := &Doctor{}
	d.SetNextDepartment(&Hospital{})
	d.Execute()
}

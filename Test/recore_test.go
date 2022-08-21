package Test

import "testing"

func TestRecord(t *testing.T) {
	var c = &Myctx{}
	//Test(c)
	R := Record(Test, c)
	R(c)

}

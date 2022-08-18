package observer

import "testing"

func TestName(t *testing.T) {

	subj1 := &Subject{}
	subj1.Add(ObserverTemp1{})
	subj1.Add(ObserverTemp2{})
	subj1.Notify()

}

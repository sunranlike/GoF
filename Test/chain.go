package Test

import "fmt"

//type Department struct {
//	nextDepartment *IDepartment
//}

type IDepartment interface {
	SetNextDepartment(nextDepartment IDepartment)
	Execute()
}

//Hospital 部门1：医院
type Hospital struct {
	IDepartment
}

func (h *Hospital) SetNextDepartment(nextDepartment IDepartment) {
	h.IDepartment = nextDepartment
}
func (h *Hospital) Execute() {
	fmt.Println("go to hospital")
	if h.IDepartment == nil {
		return
	}
	h.IDepartment.Execute()

	return
}

//部门2：医生
type Doctor struct{ IDepartment }

func (d *Doctor) SetNextDepartment(nextDepartment IDepartment) {
	d.IDepartment = nextDepartment
}
func (d *Doctor) Execute() {

	fmt.Println("go to doctor")
	if d.IDepartment == nil {
		return
	}
	d.IDepartment.Execute()
	return
}

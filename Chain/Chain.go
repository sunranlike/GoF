package Chain

import (
	"fmt"
)

type Subject struct {
	name string
}
type Department interface { //责任链模式的每个节点都叫做部门
	Execute()
	SetNext(department Department)
}

type hospital struct {
	nextDepartment Department
}

func Newhospital() Department {
	return &hospital{}
}

func (h hospital) Execute() {
	fmt.Println("去了医院")
	h.nextDepartment.Execute()
}
func (h *hospital) SetNext(department Department) {
	h.nextDepartment = department
}

type doctor struct {
	nextDepartment Department
}

func Newdoctor() Department {
	return &doctor{}
}

func (d doctor) Execute() {
	fmt.Println("找了医生")
	d.nextDepartment.Execute()
}
func (d *doctor) SetNext(department Department) {
	d.nextDepartment = department
}

type medical struct {
	nextDepartment Department
}

func Newmedical() Department {
	return &medical{}
}

func (m medical) Execute() {
	fmt.Println("买了药")
	//m.nextDepartment.Execute()
}
func (m *medical) SetNext(department Department) {
	m.nextDepartment = department
}

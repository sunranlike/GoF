package observer

import "fmt"

type IObserver interface { //定义接口，观察者接口，需要有一个Update方法，这个方法在收到更改时会执行
	Update()
}
type Subject struct {
	observers []IObserver //观察者列表
}
type ISubject interface {
	Add(observer IObserver)
	Remove(observer IObserver)
	Notify()
}

func (s *Subject) Add(observer IObserver) {
	s.observers = append(s.observers, observer) //直接加到最后
}
func (s *Subject) Remove(observer IObserver) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}
func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update()
	}
}

type ObserverTemp1 struct {
}

func (o ObserverTemp1) Update() {
	fmt.Println("tihs is observerTemo1 Update")
}

type ObserverTemp2 struct {
}

func (o ObserverTemp2) Update() {
	fmt.Println("tihs is observerTemo2 Update")
}

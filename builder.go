package builder

//建造者有四个部分：产品，建造者，指挥者
//此例子中，产品是book，建造者是Writer，指挥者是boss
//流程就是建造者实现生产产品的流程，指挥者指挥建造者生产产品，最后返回产品

//book ：产品
type book struct {
}

//定义建造者的接口,建造者要实现这些接口

type WriteBook interface {
	paper() error
	print() error
	publish() error
	getBook() book
}

//定义实际的建造者！

type Writer struct{}

func (b Writer) paper() error   { return nil }
func (b Writer) print() error   { return nil }
func (b Writer) publish() error { return nil }
func (b Writer) getBook() book  { return book{} }

//指挥者

type Boss struct {
	Writer
}

func (a Boss) SetW(w Writer) {
	a.Writer = w
}

func (a Boss) LeaderGetBook() book {
	return a.Writer.getBook()

}
func (a Boss) Work() book {
	a.paper()
	a.print()
	a.publish()
	return a.LeaderGetBook()
}

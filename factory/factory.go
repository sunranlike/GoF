package factory

//简单工厂！
//简单工厂就是 一个 工厂函数，根据传入的值，返回一个实例对象，
//所以这个工厂函数的返回值就需要是一个

type Article interface { //定义一个接口 写入两个方法
	ReadArticle(content string) string
	WriteArticle(content string) string
}

//定义两个结构体实现Article接口
type ChineseArticle struct {
}

func (a ChineseArticle) ReadArticle(content string) string  { return "read chinese:" + content }
func (a ChineseArticle) WriteArticle(content string) string { return "write chinese:" + content }

type EnglishArticle struct {
}

func (e EnglishArticle) ReadArticle(content string) string  { return "read english:" + content }
func (e EnglishArticle) WriteArticle(content string) string { return "write english:" + content }

// SimpleFactory 实现简单工厂：根据传入的值返回对应的结构体，需要用到swithc
func SimpleFactory(name string) Article {
	switch name {
	case "cn":
		return ChineseArticle{}
	case "en":
		return EnglishArticle{}
	}
	return nil
}

//实现工厂方法：不用根据传入的值返回特定的结构体，也就是每个实例都需要一个结构体去返回
//也就是相比简单工厂，工厂方法需要给每个实例都新建一个结构体和创建方法，不是统一的一个方法！

type CreateChineseArticle struct{}

func (CreateChineseArticle) Create() Article { return ChineseArticle{} }

type ChineseEnglishArticle struct{}

func (ChineseEnglishArticle) Create() Article { return EnglishArticle{} }

//实现抽象工厂！

//定义工厂接口
type BookCompany interface {
	Create() Book
}

//定义产品
type Book interface {
	ReadBook() string
	WriteBook() string
}

//实现两个产品
type ChineseBooks struct{}

func (ChineseBooks) ReadBook() string  { return "读中文书" }
func (ChineseBooks) WriteBook() string { return "写中文书" }

type EnglishBooks struct{}

func (EnglishBooks) ReadBook() string  { return "read english book" }
func (EnglishBooks) WriteBook() string { return "write english book" }

//实现工厂
type CnbookCompany struct{}

func (CnbookCompany) Create() Book {
	return ChineseBooks{}
}

type EnbookCompany struct{}

func (EnbookCompany) Create() Book {
	return EnglishBooks{}
}

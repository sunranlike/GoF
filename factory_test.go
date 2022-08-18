package factory

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestSimpleFactory(t *testing.T) {
	//使用简单工厂
	read := SimpleFactory("cn").ReadArticle("平凡的世界")
	fmt.Println(read)
	write := SimpleFactory("en").WriteArticle("book")
	fmt.Println(write)
	assert.Equal(t, read, "read chinese:平凡的世界")
	assert.Equal(t, write, "write english:book")

	//工厂方法
	cnCreator := CreateChineseArticle{}
	cnArticle := cnCreator.Create()
	assert.Equal(t, cnArticle.ReadArticle("book"), "read chinese:book")

	//抽象工厂
	company := &CnbookCompany{}
	assert.Equal(t, company.Create().ReadBook(), "读中文书")

}

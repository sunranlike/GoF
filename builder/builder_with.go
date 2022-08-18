package builder

import (
	"sync"
)

//该例子实现建造者+单例懒汉+原型

//定义产品
type product struct {
	id    int
	price int
	name  string
}
type Prototype interface { // 定义prototype原型接口
	Clone() product
}
type BuilderInterface interface { // 定义builder接口
	SetId(id int) BuilderInterface
	SetPrice(price int) BuilderInterface
	SetName(name string) BuilderInterface
	Build() product
	Prototype
}

//定义一个builder
type builderImpl struct {
	product product
}

var once sync.Once
var b *builderImpl

func (b *builderImpl) SetId(id int) BuilderInterface {
	b.product.id = id
	return b
}
func (b *builderImpl) SetPrice(price int) BuilderInterface {
	b.product.price = price
	return b
}
func (b *builderImpl) SetName(name string) BuilderInterface {
	b.product.name = name
	return b
}
func NewBuilder() BuilderInterface { //单例模式，是一个懒汉模式，在调用的时候才返回
	once.Do(func() { b = &builderImpl{} })
	return b
}
func (b *builderImpl) Build() product {
	b.SetId(1).SetName("sunran").SetPrice(1)

	return b.product
}

//实现原型模式的接口
func (b *builderImpl) Clone() product {
	var temp product
	temp.id, temp.price, temp.name = b.product.id, b.product.price, b.product.name
	return temp
}

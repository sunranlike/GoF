package singleton

//以下是饿汉模式的实现：
import "sync"

//定义msgPool结构体
type messagePool struct {
	pool *sync.Pool //使用pool的结构体在实例化的时候需要赋值一个实例化函数
}

type Message struct {
	Count int
}

var msgPool = &messagePool{ //饿汉模式的关键就是在于直接实现
	// 如果消息池里没有消息，则新建一个Count值为0的Message实例
	pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }}, //实例化的时候要给pool的new字段赋值，赋值内容是一个函数

}

//获取消息池的方法
func Instance() *messagePool {
	return msgPool
}
func (msgPool *messagePool) Addmsg(msg *Message) {
	msgPool.pool.Put(msg)
}
func (msgPool *messagePool) Getmsg() *Message {
	return msgPool.pool.Get().(*Message) //从msgpool获取
}

//以下是懒汉模式的实现msgpool：

var once = &sync.Once{}

// 消息池单例，在首次调用时初始化
var msgPool2 *messagePool

// 全局唯一获取消息池pool到方法
func Instanc2() *messagePool {
	// 在匿名函数中实现初始化逻辑，Go语言保证只会调用一次
	once.Do(func() {
		msgPool2 = &messagePool{
			// 如果消息池里没有消息，则新建一个Count值为0的Message实例
			pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
		}
	})
	return msgPool2
}

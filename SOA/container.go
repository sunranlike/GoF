package SOA

import "sync"

type Container interface {
	Make(name string, forceNew bool) interface{} //根据name
	Bind(sp ServiceProvider)
}
type myContainer struct { //优先小写,只能New方法实例化对象.因为有map!
	lock     sync.RWMutex
	FuncMap  map[string]NewInstance //创建实例的方法map
	Instance map[string]interface{} //实例的map
}

func NewMyContainer() *myContainer { //实例化对象,map要使用make初始化
	return &myContainer{
		lock:     sync.RWMutex{},
		FuncMap:  make(map[string]NewInstance),
		Instance: make(map[string]interface{}),
	}
}
func (c *myContainer) Make(name string, forceNew bool) interface{} {
	c.lock.RLock()         //先加锁
	defer c.lock.RUnlock() //defer unlock

	fun := c.FuncMap[name] //从创建实例方法的FuncMap中获取创建实例方法

	if forceNew {
		//如果要强制用新的,
		ins := fun()
		c.Instance[name] = ins
		return ins
	}
	//如果没有强制用新的:

	if ins, ok := c.Instance[name]; ok { //去实例map中寻找,如果ok就代表存在!直接返回
		return ins
	}
	//如果不forceNew,在instance中也不存在这个实例,就要创建+存入map+返回
	instance := fun()           //执行实例化方法,获取的实例返回
	c.Instance[name] = instance //将实例存入mao中
	return instance
}
func (c *myContainer) Bind(sp ServiceProvider) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	c.FuncMap[sp.Name()] = sp.Register() //调用register方法,将创建实例的方法存入FuncMap中

}

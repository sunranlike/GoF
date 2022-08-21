package SOA

import (
	"testing"
)

func TestContainer(t *testing.T) {
	container := NewMyContainer()
	var test *TestServiceProvider
	test = &TestServiceProvider{}
	container.Bind(test)
	testinstace := container.Make(test.Name(), false).(TestServiceContract) //类型断言!很重要的语法特性
	testinstace.Hello(" sunran ")

}

//所以我们是如何实现SOA架构?
//简单来说,比起直接新建一个实例,我们添加了两层东西:
//1container接口
//2 Service provider接口
//3 service 的接口和类型断言
//4容器

//1container
//container接口提供了bind方法和make方法,
//bind方法将sp提供的实例化方法注册到容器中,
//make方法调用容器中的实例化方法,生产实例化对象

//2Service provider接口
//sp接口定义register方法和Name方法,
//register方法 返回实例化方法(返回值是一个函数类型变量,利用函数闭包)
//而Name则是返回服务名称

//3 我们要定义服务所提供的功能的接口,然后在外部通过类型断言,可以让我们的调用这个接口的方法
//为什么要这麻烦要断言?原因时我们的绑定是发生在运行时,只有在运行时我们才能真正的获取到实例,所以类型断言帮助我们对没在编译时的类型进行了判断
//从而让我们可以在写代码的时候就获取到运行时特性,这是一种弱耦合,我们不知道make返回的是个什么东西,
//但是通过类型断言,我们可以让左值符合我们的预期(如果不符合会bug)
//这是一个很重要的语法特性还有一个特性就是反射

//4容器,容器就是实现容器接口的一个实例,需要有两个map字段去存储这些实例化方法和实例.
//同时容器不要外部可以访问,最好定义一下New方法以实现一些初始化给外部访问.

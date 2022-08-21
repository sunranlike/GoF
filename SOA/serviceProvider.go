package SOA

import "fmt"

type NewInstance func() interface{} //类型别名,go里面的函数也是一等公民的思想
type ServiceProvider interface {
	Register() NewInstance //注册，用来实现Di，返回的是实例化函数，容器在bind会将这个函数的返回值存入容器
	Name() string          //用来获取实例的名字
}
type TestServiceProvider struct {
}

func (s *TestServiceProvider) Register() NewInstance { //返回闭包函数
	return func() interface{} {
		return &TestService{
			name: "test",
		}
	}
}
func (s *TestServiceProvider) Name() string { return "test" }

type TestServiceContract interface {
	Hello(string)
}
type TestService struct {
	name string //
}

func (t *TestService) Hello(name string) {
	fmt.Println("Hello" + name + "i am " + t.name)

}

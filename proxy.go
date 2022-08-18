package proxy

import (
	"fmt"
	"log"
	"time"
)

//声明被代理对象
type User struct{}

func (u *User) Login() (string, error) {
	fmt.Println("登录")
	return "", nil
}

type UserProxy struct {
	User
}

func (u *UserProxy) ProxyLogin() (string, error) {
	log.Println("代理启动")
	t := time.Now().Unix()
	log.Println(t)
	u.Login()
	fmt.Print("已登录,耗时为：")
	fmt.Println(time.Now().Unix() - t)
	return "", nil

}

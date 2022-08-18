package singleton

import "sync"

// 以下实现饿汉模式

type Singleton struct{}  //定义一个Singleton结构体，是用来被单例模式实现的
var singleton *Singleton //定义一个指向上面那个结构体的指针。

func init() { //init函数，在全局变量后面执行吗，main之前执行
	singleton = &Singleton{} //让全局变量指针直接指向了一个实例化的Singleton结构体
}

func GetInstance() *Singleton {
	return singleton //直接返回指针即可
}

//以下实现懒汉模式

var lazy *Singleton //定义一个全局变量指针

var lazyonce *sync.Once //定义once

func GetllazyInstance() *Singleton { //懒汉模式获取实例化对象
	if lazyonce == nil { //判断指针是否已经被指向结构体了了，如果已经指向了，不是nil说明
		lazyonce.Do(func() { //调用once.DO,只会执行一次！
			lazy = &Singleton{} //将lazy指针指向一个实例化的Singleton
		})
	}

	return lazy //最后返回这个指针，这个指针在调用getllazyInstance函数的时候会将这个指针指向一个结构体
}

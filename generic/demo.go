package generic

import "fmt"

/*
	范型demo https://blog.csdn.net/neweastsun/article/details/128275384
*/
// Number 必须为接口
type Number interface {
	int64 | float64
}

// 范型函数，范例1
func GenericSum[T Number](nums []T) T {
	var sum T // 范型变量
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 范型函数，范例2 【这种方式如果类型很多，代码可读性不好，所以定义一个Number接口代替】
func GenericSum1[T int64 | float64](nums []T) T {
	var sum T // 范型变量
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 通用类型
func GenericPrint[T any](nums []T) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

///////////////////////////////////////
//这里的 T 需要实现 interface

type Info interface {
	Name() string
}

type UserInfo struct {
}

func (t UserInfo) Name() string {
	return "UserInfo"
}

type ProductInfo struct {
}

func (t ProductInfo) Name() string {
	return "ProductInfo"
}

// 定义打印Name的范型函数:最简单的理解方式： Type T = Info,其实T就是Info的别名
func PrintInfo[T Info](info T) {
	fmt.Println(info.Name())
}

// 接口类型【范型】

type Writer[T any] interface {
	Write(data T)
}

type FileWriter[T any] struct {
}

func (t FileWriter[T]) Write(data T) {
	fmt.Println(data)
}

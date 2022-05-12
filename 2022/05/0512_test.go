package _5

import (
	"fmt"
	"testing"
)

// init() 函数是什么时候执行的？
func init() {
	fmt.Println("init1:", a)
}

// 一个包甚至是一个源文件可以有多个init() 函数，且执行顺序不保证
// init() 函数不能被其他函数调用
func init() {
	fmt.Println("init2:", b)
}

var a = 10

const b = 100

// 一句话总结： import –> const –> var –> init() –> main()
func TestA(t *testing.T) {
	fmt.Println("main:", a)
}

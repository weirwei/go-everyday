package _3

import (
	"testing"
)

/**
下面代码输出什么？

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
A. 1 1
B. 0 1
C. 1 0
D. 0 0
*/

// defer 修饰的匿名函数，只能更新具名返回值
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func Test0325(t *testing.T) {
	t.Log(increaseA())
	t.Log(increaseB())
}

package _3

import "testing"

/*
有下面 3 行代码：

// 32 位机器
1）var x int32 = 32.0
2）var y int = x
3）var z rune = x
它们是否能编译通过？为什么？

如果面试时问这道题，你需要想想面试官想考察你什么。
*/

/**
int 是一个独特的类型，它的大小跟CPU有关，如果CPU是32位，它就是32位，如果CPU是64位，它就是64位
rune 就是一个int32 的别名，一般用rune 表示字符，int32表示数字
*/
func Test0323(t *testing.T) {
	var x int32 = 32.0
	//var y int = x
	var z rune = x
	//_ = y
	_ = z
}

package _3

import (
	"fmt"
	"testing"
)

/**
通常，JS 面试，闭包应该是必考的题目。随着越来越多的语言对函数式范式的支持，闭包问题经常出现。在 Go 语言中也是如此。

这是 Go 语言爱好者周刊第 90 期的一道题目。以下代码输出什么？
A：Hi All；B：Hi go All；C：Hi；D：go All

这道题目答对的人蛮多的：60%。不管你是答对还是答错，如果最后再加一行代码：fmt.Println(a("All"))，它输出什么？想看看你是不是蒙对了。（提示：你可以输出 t 的地址，看看是什么情况。）
*/

/**
对闭包来说，函数在该语言中得是一等公民。一般来说，一个函数返回另外一个函数，这个被返回的函数可以引用外层函数的局部变量，这形成了一个闭包。
通常，闭包通过一个结构体来实现，它存储一个函数和一个关联的上下文环境。
但 Go 语言中，匿名函数就是一个闭包，它可以直接引用外部函数的局部变量，因为 Go 规范和 FAQ 都这么说了。
*/
func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		fmt.Println(&t)
		return t
	}
	return c
}

func Test0328(t *testing.T) {
	a := app()
	b := app()
	a("go")
	t.Log(b("All"))
	t.Log(a("All"))
}

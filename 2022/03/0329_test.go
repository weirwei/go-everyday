package _3

import (
	"fmt"
	"testing"
)

// 下面代码段输出什么？

/**
29 29 28。变量person是一个指针变量
1.person.age 此时是将28当做defer 参数的函数，会把28缓存在栈中，等到最后执行该defer语句的时候去除，即取出28
2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
3.很简单，闭包引用，输出 29；

又由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 29 28。
*/

type Person struct {
	age int
}

func Test0329(t *testing.T) {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

package _4

import (
	"fmt"
	"testing"
)

/**
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4

defer 定义的函数是延迟函数，会延迟执行
*/
func Test0412(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
}

/**
test 3 1
test 2 2
inc 2
test 1 3
main 2
*/
func Test0412Next(t *testing.T) {
	fmt.Println("main", inc())
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func inc() int {
	t := &test{num: 0}
	defer t.Inc(3).Inc(2).Inc(1)
	fmt.Println("inc", t.num)
	return t.num
}

type test struct {
	num int
}

func (t *test) Inc(flag int) *test {
	t.num++
	fmt.Println("test", flag, t.num)
	return t
}

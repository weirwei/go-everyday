package _4

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

/*

values   method receivers
--------------------------
T        (t T)
*T       (t T)或(t *T)

*/
func Test0415(t *testing.T) {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

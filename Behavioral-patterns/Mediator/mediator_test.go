package Mediator

import (
	"fmt"
	"testing"
)

// 中介者——对象行为型模式
func TestSum(t *testing.T) {
	fmt.Printf("%#v\n", Sum(One{}, Two{}))
	fmt.Printf("%d\n", Sum(1, 2))
	fmt.Printf("%d\n", Sum(One{}, 2))
}

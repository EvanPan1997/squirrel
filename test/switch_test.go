package test

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	var i = 1
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}
}

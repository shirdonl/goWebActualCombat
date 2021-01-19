package testexample

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	array := []int{6, 8, 10}
	ret := Min(array)
	fmt.Println(ret)
}

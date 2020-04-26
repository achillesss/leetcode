package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	var l = NewDList()
	l.Append(1).Append(2).Append(3).Append(4)
	fmt.Printf("list: %+v\n", l)
}

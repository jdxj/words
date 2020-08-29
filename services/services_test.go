package services

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []byte{1, 2, 3}
	fmt.Printf("%v\n", s[3:])
}

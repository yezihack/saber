package util

import (
	"fmt"
	"testing"
)

func TestIsExistLastSep(t *testing.T) {
	b1 := HasSuffix("aa/aab", "/")
	fmt.Printf("b1:%v\n", b1)

	b1 = HasSuffix("aa/aab/", "/")
	fmt.Printf("b1:%v\n", b1)
}

func TestAddSuffix(t *testing.T) {
	fmt.Println(AddSuffix("aa/bb", "/"))
	fmt.Println(AddSuffix("aa/bb/", "/"))
}

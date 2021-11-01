package webuery_test

import (
	"fmt"
	"testing"

	"github.com/func25/gofunc/webfunc/webuery"
)

func TestStringToInt64s(t *testing.T) {
	x := "1,2,3,4"
	arr, err := webuery.StringToInt64s(x, ",")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(arr)

	x = "1,2,3,,,4,"
	arr, err = webuery.StringToInt64s(x, ",")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(arr)

	x = "1,2,,"
	arr, err = webuery.StringToInt64s(x, ",")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(arr)

	x = "1,a,asd,3,"
	arr, err = webuery.StringToInt64s(x, ",")
	if err == nil {
		t.Error("error of course")
		return
	}
	fmt.Println(arr)
}

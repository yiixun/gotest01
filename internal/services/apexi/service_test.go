package apexi_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yiixun/gotest01/v2/pkg/service/fishing"
)

func TestInterface(t *testing.T) {
	var a = new(fishing.Fishing)

	fmt.Println(reflect.TypeOf(a))
}

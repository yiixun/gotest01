package apex_test

import (
	"fmt"
	"testing"

	"github.com/yiixun/gotest01/v2/pkg/apex"
)

// Test B01 depends on A01
type A01 struct {
	V string
}
type B01 struct {
	A01 *A01
	V   string
}

func NewA01() {
	a01 := A01{V: "Hi A01"}
	apex.Register("A01", &a01)
}

func NewB01() {
	a01 := apex.Get("A01").(*A01)
	b01 := B01{V: "Hi B01", A01: a01}

	apex.Register("B01", &b01)
}

func TestApex01(t *testing.T) {
	apex.Bootstrap()

	apex.Bootup(
		NewA01,
		NewB01,
	)

	apex.WaitUtilAllUp()

	b01 := apex.Get("B01").(*B01)
	fmt.Println(b01.V)
	fmt.Println(b01.A01.V)

}

package go_cc

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	res := If(false, "K.Y.").
		Elif(true, "L.L.").
		Else("K.S.")
	fmt.Println(res)

}

package test

import "testing"

func Test_Soma(t *testing.T) {
	total := Soma(15+15)

	if total != 30 {
		t.Error("result is invalid")
	}
}
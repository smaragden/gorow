package com

import "testing"

func Test_Int2bytes(t *testing.T) {
	res := _Int2bytes(5, 5)
	t.Logf("res: %v", res)
}

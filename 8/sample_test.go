package main

import (
	"testing"
)

func TestMultiple(t *testing.T) {

	var x, y, result = 2, 2, 4
	res := Multiple(x, y)

	if res != result {
		t.Errorf("%d!=%d", res, result)
	}

}

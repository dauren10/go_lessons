package main

import (
	"fmt"
	"os"
	"testing"
)

//тесты для всего пакета напрмиер маин

func TestMain(m *testing.M) {
	fmt.Println("testting package")
	res := m.Run()
	fmt.Println("teaar down")
	os.Exit(res)

}

func TestMultiple2(t *testing.T) {
	t.Parallel()
	t.Run("simple", func(t *testing.T) {

		var x, y, result = 2, 2, 4
		res := Multiple(x, y)

		if res != result {
			t.Errorf("%d!=%d", res, result)
		}
		t.Run("2", func(t *testing.T) {
			var x, y, result = 8, 8, 64
			res := Multiple(x, y)

			if res != result {
				t.Errorf("%d!=%d", res, result)
			}
		})
	})

	t.Run("medium", func(t *testing.T) {
		var x, y, result = 8, 8, 64
		res := Multiple(x, y)

		if res != result {
			t.Errorf("%d!=%d", res, result)
		}
	})

	t.Run("groupA", func(t *testing.T) {
		t.Run("11", func(t *testing.T) {
			var x, y, result = 8, 8, 64
			res := Multiple(x, y)

			if res != result {
				t.Errorf("%d!=%d", res, result)
			}
		})
		t.Run("12", func(t *testing.T) {
			var x, y, result = 8, 8, 64
			res := Multiple(x, y)

			if res != result {
				t.Errorf("%d!=%d", res, result)
			}
		})
	})

}

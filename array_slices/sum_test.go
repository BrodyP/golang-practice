package array_slices

import "testing"

func TestSum(t *testing.T) {

	t.Run("collections of 5 numbers", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		sum := Sum(array)
		expected := 15

		if expected != sum {
			t.Errorf("expected %d but got %d given %v", expected, sum, array)
		}

	})

	t.Run("collections of any numbers", func(t *testing.T) {
		array := []int{1, 2, 3}
		sum := Sum(array)
		expected := 6

		if expected != sum {
			t.Errorf("expected %d but got %d given %v", expected, sum, array)
		}

	})

}

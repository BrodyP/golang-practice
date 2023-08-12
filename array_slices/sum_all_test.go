package array_slices

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect %v but got %v ", want, got)
	}
}

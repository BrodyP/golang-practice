package structMethodInterface

import "testing"

func TestCalculateArea(t *testing.T) {

	t.Run("test calculate rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("expect %.2f got %.2f", want, got)
		}
	})

	t.Run("test calculate circle area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("expect %.2f got %.2f", want, got)
		}
	})
}

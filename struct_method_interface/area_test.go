package structMethodInterface

import "testing"

// func TestCalculateArea(t *testing.T) {

// 	t.Run("test calculate rectangle area", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Area()
// 		want := 100.0

// 		if got != want {
// 			t.Errorf("expect %.2f got %.2f", want, got)
// 		}
// 	})

// 	t.Run("test calculate circle area", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("expect %.2f got %.2f", want, got)
// 		}
// 	})
// }

//refactor 01
// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("expect %.2f but got %.2f", want, got)
// 		}
// 	}

// 	t.Run("calculate rectangle area", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		checkArea(t, rectangle, 100)
// 	})

// 	t.Run("calculate circle area", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// //refactor2 and add triangle
// func TestArea(t *testing.T) {

// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Rectangle{10.0, 10.0}, 100.0},
// 		{Circle{10.0}, 314.1592653589793},
// 		{Triangle{12.0, 6.0}, 36.0},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if tt.want != got {
// 			t.Errorf("expect %.2f got %.2f", tt.want, got)
// 		}
// 	}
// }

//last refactor and add test name

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

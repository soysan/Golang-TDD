package structs

import "testing"

func TestArea(t *testing.T) {
	// Table base test.
	areaTests := []struct {
		shape   Shape
		name    string
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name.
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %#v want %.2f", got, tt.hasArea)
			}
		})
	}

	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//  t.Helper()
	//  got := shape.Area()
	//  if got != want {
	//    t.Errorf("got %g want %g", got, want)
	//  }
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//  rectangle := Rectangle{12.0, 6.0}
	//  checkArea(t, rectangle, 72.0)
	//})
	//
	//t.Run("circle", func(t *testing.T) {
	//  circle := Circle{10}
	//  checkArea(t, circle, 314.1592653589793)
	//})
}

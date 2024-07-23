package shapes

import "testing"

func TestPerimeter(t *testing.T) {
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f expected %.2f", got, want)
		}
	}
	t.Run("will calculate RECT area", func(t *testing.T) {
		// t.Helper()
		rectangle := Rectangle{3.0, 6.0}
		checkArea(t, rectangle, 18)
	})
	t.Run("will calculate CIRCLE area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTable(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 4.0, Height: 6.0}, hasArea: 24},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 4, Height: 10}, hasArea: 59},
	}

	t.Run("test area calculations on shapes", func(t *testing.T) {

		for _, tt := range areaTests {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %.2f hasArea %.2f", tt.shape, got, tt.hasArea)
			}
		}
	})
}

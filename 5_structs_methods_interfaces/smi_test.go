package smi

import "testing"

func assertValuesAreEqual(t testing.TB, shape Shape, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Unexpected output given %#v- got %g, want %g", shape, got, want)
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	assertValuesAreEqual(t, shape, got, want)
}

func TestArea(t *testing.T) {
	t.Run("Calculates correct area for rectangles", func(t *testing.T) {
		rect := Rectangle{5.0, 10.0}
		got := rect.Area()
		want := 50.0
		assertValuesAreEqual(t, rect, got, want)
	})

	t.Run("Calculates correct area for circles", func(t *testing.T) {
		circle := Circle{100.0}
		got := circle.Area()
		want := 314.1592653589793
		assertValuesAreEqual(t, circle, got, want)
	})
}

func TestAreaUsingInterfaces(t *testing.T) {
	t.Run("Check area function for rectangles", func(t *testing.T) {
		rect := Rectangle{5.0, 10.0}
		checkArea(t, rect, 50.0)
	})

	t.Run("Check area function for circle", func(t *testing.T) {
		rect := Circle{100.0}
		checkArea(t, rect, 314.1592653589793)
	})
}

func TestAreaUsingTables(t *testing.T) {
	// We are:
	// first, declaring a struct
	// second, populating it
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{5.0, 10.0}, want: 50.0},
		{name: "Circle", shape: Circle{100.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{5.0, 10.0}, want: 25.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assertValuesAreEqual(t, tt.shape, got, tt.want)
		})
	}
}

package shapes

import "testing"


func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Shape is an interface for objects that can calculate their area
type Shape_Area interface {
	Area() float64
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape_Area
		hasArea  float64
	}{
		{name:"Rectangle",shape:  Rectangle{12, 6},hasArea: 72.0},
		{name:"Circle", shape: Circle{10},hasArea: 314.1592653589793},
		{name:"Triangle",shape:  Triangle{12, 6},hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}
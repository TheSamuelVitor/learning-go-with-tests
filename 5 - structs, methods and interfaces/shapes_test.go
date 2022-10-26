package structs

import "testing"

func ComparisonBetweenExpectedAndGot(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot %.2f\nwant %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0
	ComparisonBetweenExpectedAndGot(t, got, want)
}

func TestArea(t *testing.T) {

	// criando uma "struct anonima" que recebe dois "parametros"
	// o primeiro é o Shape e o want é o resultado esperado do teste
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72},
		{name: "circle", shape: Circle{Raio: 10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	// iterando no struct
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			// passando a forma e chamando a funcao area
			got := (tt.shape).Area()
			if got != tt.want {
				t.Errorf("%#v \ngot %g\n want %g", tt.shape, got, tt.want)
			}
		})
	}
}

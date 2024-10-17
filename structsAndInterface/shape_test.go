package structsandinterface

import(
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("testing rectangle Perimeter",func(t *testing.T) {
		rect:=Rectangle{10.0,15.0}
		got:=rect.Perimeter()
		want:=50.00
		if got!=want {
			t.Errorf("got %g but wanted %g",got,want)			
		}
	})

	t.Run("testing circle perimeter",func(t *testing.T) {
		circle:=Circle{10}
		got:=circle.Perimeter()
		want:=62.83185307179586

		if got!=want {
			t.Errorf("got %g but wanted %g",got,want)
		}
	})
	
}


func TestArea(t *testing.T) {
	check:=func(t testing.TB,shape Shape,want float64){
		t.Helper()
		got:=shape.Area()
		if got!=want {
			t.Errorf("got %g but wanted %g",got,want)
		}
	}

	t.Run("testing rectangle Area",func(t *testing.T) {
		rect:=Rectangle{10.0,15.0}
		want:=150.00
		check(t,rect,want)
	})

	t.Run("testing circle Area",func(t *testing.T) {
		circle:=Circle{10.0}
		want:=314.1592653589793
		check(t,circle,want)
	})

}
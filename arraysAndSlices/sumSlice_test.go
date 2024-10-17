package arraysandslices

import(
	"testing"
)

func TestSumSlice(t *testing.T) {
	number:=[]int{1,2,3,4}
	got:=SumSlice(number)
	expected:=10

	if got!= expected {
		t.Errorf("expected %d but got %d",expected,got)
	}
	
}
package integers

import (
	"testing"
)

func TestSum(t *testing.T){
	sum:=Sum(2,2)
	expected:=4

	if sum!=expected {
		t.Errorf("expected %d but got %d ",sum,expected)
	}
}

func BenchmarkSum(b *testing.B) {
	for i:=0;i<b.N;i++{
		Sum(i,1)
	}
	
}
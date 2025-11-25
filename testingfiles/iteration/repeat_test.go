package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected{
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

//benchmarking

func BenchmarkRepeat(b *testing.B){
	for b.Loop(){
		Repeat("a", 5)
	}
}
package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("b", 5)
	want := "bbbbb"

	if got != want {
		t.Errorf("got '%q' but want '%q'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	b.Run("without optimization", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Repeat("b", 5)
		}
	})
	
	b.Run("with optimization", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			RepeatOptimized("b", 5)
		}
	})

}

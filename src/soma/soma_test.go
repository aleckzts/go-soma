package main
import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(2, 2)
	if resultado != 4 {
		t.Errorf("soma(2, 2) = %d; deveria ser 4", resultado)
	}
}
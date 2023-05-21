package calculator_test

import (
	"github.com/mocamocaland/tdd-in-go-sample/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0

	// Act
	got := e.Add(x, y)

	// Assert
	if got != want {
		t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want:%.2f", 2.5, 3.5, got, 6.0)
	}
}

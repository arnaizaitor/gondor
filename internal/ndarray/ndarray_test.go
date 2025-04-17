package ndarray_test

import (
	"testing"

	"github.com/arnaizaitor/gondor/internal/ndarray"
)

func TestNewValidShape(t *testing.T) {
	a, err := ndarray.New(2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if a == nil {
		t.Fatal("expected NDArray, got nil")
	}

	expectedSize := 6
	if len(a.Shape()) != 2 || a.Shape()[0] != 2 || a.Shape()[1] != 3 {
		t.Errorf("unexpected shape: got %v", a.Shape())
	}

	if a.Size() != expectedSize {
		t.Errorf("expected size %d, got %d", expectedSize, a.Size())
	}
}

func TestNewInvalidShape(t *testing.T) {
	_, err := ndarray.New(0, 3)
	if err == nil {
		t.Error("expected error for zero dimension, got nil")
	}
}

func TestGetValid(t *testing.T) {
	a, err := ndarray.New(2, 2)
	if err != nil {
		t.Fatalf("error creating array: %v", err)
	}

	// manually set value using known index logic
	a.Set(42.0, 1, 1)
	val, err := a.Get(1, 1)
	if err != nil {
		t.Fatalf("unexpected error on Get: %v", err)
	}

	if val != 42.0 {
		t.Errorf("expected 42.0, got %f", val)
	}
}

func TestGetOutOfBounds(t *testing.T) {
	a, _ := ndarray.New(2, 2)

	_, err := a.Get(3, 0)
	if err == nil {
		t.Error("expected out-of-bounds error, got nil")
	}
}

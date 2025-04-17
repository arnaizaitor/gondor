// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║ ███╗   ██╗██████╗  █████╗ ██████╗ ██████╗  █████╗ ██╗   ██╗                        ║
// ║ ████╗  ██║██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔══██╗╚██╗ ██╔╝                        ║
// ║ ██╔██╗ ██║██║  ██║███████║██████╔╝██████╔╝███████║ ╚████╔╝                         ║
// ║ ██║╚██╗██║██║  ██║██╔══██║██╔══██╗██╔══██╗██╔══██║  ╚██╔╝                          ║
// ║ ██║ ╚████║██████╔╝██║  ██║██║  ██║██║  ██║██║  ██║   ██║                           ║
// ║ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝                           ║
// ║---------------------------------------------------------------                     ║
// ║ Gondor - N-Dimensional Array Engine for Go                                         ║
// ║ Inspired by the ancient power of NumPy, forged in the fire of Go                   ║
// ║ Built for numerical computing, indexing precision, and low-level control.          ║
// ║ --------------------------------------------------------------                     ║
// ║ Author: Aitor Arnaiz                                                               ║
// ║ Project: github.com/arnaizaitor/gondor                                             ║
// ║ License: TBD                                                                       ║
// ║ --------------------------------------------------------------                     ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝

package ndarray

import (
	"fmt"
)

// NDArray represents an n-dimensional array of float64 values.
type NDArray struct {
	data    []float64
	shape   []int
	strides []int
}

// New creates a new NDArray with the given shape.
// All values are initialized to zero.
func New(shape ...int) (*NDArray, error) {
	// TODO: Implement initialization logic
	return nil, nil
}

// Get returns the value at the given indices.
func (a *NDArray) Get(indices ...int) (float64, error) {
	// TODO: Implement index calculation and bounds checking
	return 0, nil
}

// Set sets the value at the given indices.
func (a *NDArray) Set(value float64, indices ...int) error {
	// TODO: Implement index calculation and bounds checking
	return nil
}

// Shape returns the shape of the array.
func (a *NDArray) Shape() []int {
	return a.shape
}

// Reshape reshapes the array to the new shape.
// It must contain the same number of elements.
func (a *NDArray) Reshape(newShape ...int) error {
	// TODO: Implement reshape logic
	return nil
}

// Zeros returns a new NDArray of zeros with the given shape.
func Zeros(shape ...int) (*NDArray, error) {
	return New(shape...)
}

// Ones returns a new NDArray filled with 1.0 with the given shape.
func Ones(shape ...int) (*NDArray, error) {
	// TODO: Allocate and fill with 1.0
	return nil, nil
}

// Full returns a new NDArray filled with a specified value.
func Full(value float64, shape ...int) (*NDArray, error) {
	// TODO: Allocate and fill with given value
	return nil, nil
}

// Size returns the total number of elements in the array.
func (a *NDArray) Size() int {
	// TODO: Compute product of shape
	return 0
}

// String provides a simple string representation.
func (a *NDArray) String() string {
	return fmt.Sprintf("NDArray(shape=%v, data=%v)", a.shape, a.data)
}

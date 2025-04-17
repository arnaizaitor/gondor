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

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   STRUCT: NDArray – The Core of the Engine                                         ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Inspired by NumPy's internals, this struct holds:                                ║
// ║                                                                                    ║
// ║     - `data []float64` : Flat memory holding the actual values                     ║
// ║     - `shape []int`    : Dimensions of the array (e.g., [3, 4])                    ║
// ║     - `strides []int`  : Jump distances to traverse dimensions                     ║
// ║                                                                                    ║
// ║   These three together allow fast, flexible, and memory-efficient                  ║
// ║   indexing and reshaping of multidimensional arrays.                               ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   DIAGRAM: An n-dimensional array built over flat memory                           ║
// ║   ──────────────────────────────────────────────────────                           ║
// ║                                                                                    ║
// ║  ┌────────────┐    ┌────────────┐    ┌────────────────────────┐                    ║
// ║  │  shape     │    │  strides   │    │         data           │                    ║
// ║  └────────────┘    └────────────┘    └────────────────────────┘                    ║
// ║      ↓                ↓                        ↓                                   ║
// ║  [3, 4]          [4, 1]         [a00, a01, a02, a03,                               ║
// ║                                 a10, a11, a12, a13,                                ║
// ║                                 a20, a21, a22, a23]                                ║
// ║                                                                                    ║
// ║ → shape[0] = 3 rows                                                                ║
// ║ → shape[1] = 4 columns                                                             ║
// ║                                                                                    ║
// ║ → strides[0] = 4  → jump 4 elements to go to next row                              ║
// ║ → strides[1] = 1  → jump 1 element to go to next column                            ║
// ║                                                                                    ║
// ║ Example: a[2][3]  → index = 2*4 + 3*1 = 11                                         ║
// ║                    → data[11] = a23                                                ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
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

	if len(indices) != len(a.shape) {
		return 0, fmt.Errorf("number of indices (%d) does not match array dimensions (%d)", len(indices), len(a.shape))
	}

	// Calculate the flat index from the multi-dimensional indices
	flatIndex := 0
	for i, index := range indices {
		if index < 0 || index >= a.shape[i] {
			return 0, fmt.Errorf("index %d out of bounds for axis %d with size %d", index, i, a.shape[i])
		}
		flatIndex += index * a.strides[i]
	}

	// Check if the flat index is within bounds
	if flatIndex < 0 || flatIndex >= len(a.data) {
		return 0, fmt.Errorf("flat index %d out of bounds for array of size %d", flatIndex, len(a.data))
	}

	// Return the value at the calculated index
	return a.data[flatIndex], nil
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

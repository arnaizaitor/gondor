// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║ ███╗   ██╗██████╗  █████╗ ██████╗ ██████╗  █████╗ ██╗   ██╗                        ║
// ║ ████╗  ██║██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔══██╗╚██╗ ██╔╝                        ║
// ║ ██╔██╗ ██║██║  ██║███████║██████╔╝██████╔╝███████║ ╚████╔╝                         ║
// ║ ██║╚██╗██║██║  ██║██╔══██║██╔══██╗██╔══██╗██╔══██║  ╚██╔╝                          ║
// ║ ██║ ╚████║██████╔╝██║  ██║██║  ██║██║  ██║██║  ██║   ██║                           ║
// ║ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝                           ║
// ║---------------------------------------------------------------                     ║
// ║ Gondor - N-Dimensional Array Engine written in Go                                  ║
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
// ║      [3, 4]           [4, 1]              [a00, a01, a02, a03,                     ║
// ║                                            a10, a11, a12, a13,                     ║
// ║                                            a20, a21, a22, a23]                     ║
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

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: New – Create a new NDArray                                                 ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Initializes an NDArray with the given shape, zero-filled by default.             ║
// ║                                                                                    ║
// ║   - Validates the shape dimensions                                                 ║
// ║   - Computes total size and allocates flat data slice                              ║
// ║   - Computes strides in row-major order                                            ║
// ║                                                                                    ║
// ║   Returns: (*NDArray, error)                                                       ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║                                                                                    ║
// ║   a, err := New(3, 4)                                                              ║
// ║   Shape:   [3, 4]                                                                  ║
// ║   Strides: [4, 1]                                                                  ║
// ║   Data:    [0.0, 0.0, ..., 0.0] (length 12)                                        ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func New(shape ...int) (*NDArray, error) {

	if len(shape) == 0 {
		return nil, fmt.Errorf("shape must have at least one dimension")
	}

	if len(shape) > 32 {
		return nil, fmt.Errorf("shape has too many dimensions (max 32)")
	}

	// Calculate the total number of elements
	totalSize := 1
	for _, dim := range shape {
		if dim <= 0 {
			return nil, fmt.Errorf("dimension size must be positive, got %d", dim)
		}

		totalSize *= dim
	}

	// Allocate the data slice
	data := make([]float64, totalSize)

	// Calculate strides (row-major order)
	strides := make([]int, len(shape))
	stride := 1
	for i := len(shape) - 1; i >= 0; i-- {
		strides[i] = stride
		stride *= shape[i]
	}

	// Return the constructed NDArray
	return &NDArray{
		data:    data,
		shape:   shape,
		strides: strides,
	}, nil
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Get – Read a value from the NDArray                                        ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Retrieves a value from a specific multidimensional index.                        ║
// ║                                                                                    ║
// ║   - Validates number of indices                                                    ║
// ║   - Bounds check for each axis                                                     ║
// ║   - Computes flat index and returns data[offset]                                   ║
// ║                                                                                    ║
// ║   Returns: (float64, error)                                                        ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║                                                                                    ║
// ║   val, _ := a.Get(2, 3) → reads a[2][3]                                            ║
// ║   offset = 2*4 + 3*1 = 11                                                          ║
// ║   returns data[11]                                                                 ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
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

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Set – Write a value into the NDArray                                       ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Writes a value at a specific multidimensional index.                             ║
// ║                                                                                    ║
// ║   - Checks dimensionality and bounds                                               ║
// ║   - Computes flat index using strides                                              ║
// ║   - Updates the data[offset] with the given value                                  ║
// ║                                                                                    ║
// ║   Returns: error (if index out of bounds)                                          ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║                                                                                    ║
// ║   a.Set(42.0, 1, 2) → sets element a[1][2]                                         ║
// ║   offset = 1*4 + 2*1 = 6                                                           ║
// ║   data[6] = 42.0                                                                   ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func (a *NDArray) Set(value float64, indices ...int) error {

	if len(indices) != len(a.shape) {
		return fmt.Errorf("number of indices (%d) does not match array dimensions (%d)", len(indices), len(a.shape))
	}

	// Bounds checking and index calculation
	offset := 0
	for i, idx := range indices {
		if idx < 0 || idx >= a.shape[i] {
			return fmt.Errorf("index %d out of bounds for axis %d (size %d)", idx, i, a.shape[i])
		}
		offset += idx * a.strides[i]
	}

	// Set the value
	a.data[offset] = value
	return nil
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Shape – Return the array dimensions                                        ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Returns the internal shape of the array.                                         ║
// ║                                                                                    ║
// ║   - Shape is returned as []int                                                     ║
// ║   - Caller should treat it as read-only                                            ║
// ║                                                                                    ║
// ║   Returns: []int                                                                   ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║                                                                                    ║
// ║   a.Shape() → [3, 4]                                                               ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func (a *NDArray) Shape() []int {
	return a.shape
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Reshape – Change the shape of the array                                    ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Alters the shape of the NDArray without changing the underlying data.            ║
// ║                                                                                    ║
// ║   - Validates that new shape has the same total size                               ║
// ║   - Recomputes `strides` for the new shape                                         ║
// ║   - No memory is reallocated                                                       ║
// ║                                                                                    ║
// ║   Returns: error                                                                   ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║   a.Shape() → [2, 6]                                                               ║
// ║   a.Reshape(3, 4)                                                                  ║
// ║   a.Shape() → [3, 4]                                                               ║
// ║   Total elements remain: 12                                                        ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func (a *NDArray) Reshape(newShape ...int) error {
	// TODO: Implement reshape logic
	return nil
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Zeros – Create an array filled with 0.0                                    ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Constructs a new NDArray with the specified shape and zero-filled data.          ║
// ║                                                                                    ║
// ║   - Internally calls `New(shape...)`                                               ║
// ║   - Convenience helper                                                             ║
// ║                                                                                    ║
// ║   Returns: (*NDArray, error)                                                       ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║   a, _ := Zeros(2, 2)                                                              ║
// ║   a.data → [0.0, 0.0, 0.0, 0.0]                                                    ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func Zeros(shape ...int) (*NDArray, error) {
	return New(shape...)
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Ones – Create an array filled with 1.0                                     ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Constructs a new NDArray of the given shape, filled with ones.                   ║
// ║                                                                                    ║
// ║   - Uses `New` and fills `data` with 1.0                                           ║
// ║                                                                                    ║
// ║   Returns: (*NDArray, error)                                                       ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║   a, _ := Ones(2, 2)                                                               ║
// ║   a.data → [1.0, 1.0, 1.0, 1.0]                                                    ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func Ones(shape ...int) (*NDArray, error) {
	// TODO: Allocate and fill with 1.0
	return nil, nil
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Full – Create an array filled with a specific value                        ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Builds a new NDArray of a given shape and fills it with a custom value.          ║
// ║                                                                                    ║
// ║   - Uses `New`, then manually fills `data` with the given value                    ║
// ║                                                                                    ║
// ║   Returns: (*NDArray, error)                                                       ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║   a, _ := Full(7.0, 2, 2)                                                          ║
// ║   a.data → [7.0, 7.0, 7.0, 7.0]                                                    ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func Full(value float64, shape ...int) (*NDArray, error) {
	// TODO: Allocate and fill with given value
	return nil, nil
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: Size – Return total number of elements                                     ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Returns the total number of elements in the NDArray.                             ║
// ║                                                                                    ║
// ║   - Computes product of dimensions in `shape`                                      ║
// ║                                                                                    ║
// ║   Returns: int                                                                     ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║   shape = [2, 3] → size = 2 * 3 = 6                                                ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func (a *NDArray) Size() int {

	return len(a.data)
}

// ╔════════════════════════════════════════════════════════════════════════════════════╗
// ║                                                                                    ║
// ║   FUNC: String – Text representation of NDArray                                    ║
// ║   ───────────────────────────────────────────────────────────────                  ║
// ║   Returns a basic human-readable string of shape and data.                         ║
// ║                                                                                    ║
// ║   - Implements `Stringer` interface                                                ║
// ║   - Useful for printing arrays via `fmt.Println()`                                 ║
// ║                                                                                    ║
// ║   Returns: string                                                                  ║
// ║                                                                                    ║
// ║────────────────────────────────────────────────────────────────────────────        ║
// ║   EXAMPLE:                                                                         ║
// ║   fmt.Println(a) → NDArray(shape=[2 2], data=[1 2 3 4])                            ║
// ║                                                                                    ║
// ╚════════════════════════════════════════════════════════════════════════════════════╝
func (a *NDArray) String() string {
	return fmt.Sprintf("NDArray(shape=%v, data=%v)", a.shape, a.data)
}

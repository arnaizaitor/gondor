# 📘 Go Language Wiki – Learning & Practical Tips

This document is a personal wiki built through my journey to learn and master Go (Golang). It compiles practical knowledge, idioms, design patterns, and best practices.

---

## 🔹 Functions & Methods

### 🧠 What does `func (a *Type) Method(...)` mean?

- `(a *Type)` is the **method receiver** (similar to `self` in Python).
- If it’s a **pointer (`*Type`)**, you can **modify the actual struct**.
- If it’s **non-pointer (`Type`)**, you’re working on a **copy**, so changes won’t persist.

**Use a pointer when:**
- You want to modify internal state.
- The struct is large and copying it would be inefficient.

```go
func (a *NDArray) Set(...)  // modifies the array
func (a NDArray) String()   // read-only, operates on a copy
```

---

## 🔹 Variadic Parameters (`...`)

Go allows functions to accept a variable number of arguments:

```go
func Get(indices ...int)
```

This means you can call it like `Get(0, 1, 2)`, and `indices` will be of type `[]int`.

---

## 🔹 Multiple Return Values

It’s idiomatic in Go to return multiple values, especially `(value, error)` pairs:

```go
func Get(indices ...int) (float64, error)
```

Example usage:

```go
val, err := a.Get(1, 2)
if err != nil {
    log.Fatal(err)
}
```

---

## 🔹 Packages and Modules

- A Go file starts with `package name`.
- Use `internal/` for internal packages.
- Use `pkg/` for public, reusable APIs.

---

## 🔹 Project Structure (Recommended)

- `cmd/` → for executables
- `internal/` → internal business logic
- `pkg/` → public API code
- `go.mod` → module and dependencies
- `go.sum` → dependency checksums

---

## 🔹 Error Handling

- Go doesn’t use `try/catch`.
- Idiomatic pattern: return `error` as second value.

```go
result, err := DoSomething()
if err != nil {
    return nil, fmt.Errorf("operation failed: %w", err)
}
```

---

## 🔹 Formatting & Linting

- `go fmt` → format code
- `go vet` → static analysis
- `golangci-lint` → powerful multi-linter

---

## 🔹 Pointers vs Values: When to Use Them

### Should you use a pointer for read-only methods like `Get`?

Even if you're not modifying the struct, **using a pointer is often better**:

✅ Reasons to use `*NDArray` in `Get`:
- Prevents unnecessary copying of large structs.
- Keeps method call syntax consistent (all methods use pointer receiver).
- Lets you evolve the implementation (e.g., add caching) without breaking the API.

**Only use value receiver (`NDArray`) when:**
- The method is extremely lightweight.
- The struct is small and copying is trivial.
- You want to guarantee immutability.

---

## 🔹 Other Useful Concepts

### 🔸 Slices vs Arrays

- `[]int` is a **slice** (dynamic length)
- `[3]int` is a **fixed-size array**

### 🔸 Structs

```go
type NDArray struct {
    data    []float64
    shape   []int
    strides []int
}
```

---

## 🧬 The `NDArray` Struct Explained

```
╔════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║   STRUCT: NDArray – The Core of the Engine                                 ║
║   ────────────────────────────────────────────────                         ║
║   Inspired by NumPy's internals, this struct holds:                        ║
║                                                                            ║
║     - `data []float64` : Flat memory holding the actual values             ║
║     - `shape []int`    : Dimensions of the array (e.g., [3, 4])            ║
║     - `strides []int`  : Jump distances to traverse dimensions             ║
║                                                                            ║
║   These three together allow fast, flexible, and memory-efficient          ║
║   indexing and reshaping of multidimensional arrays.                       ║
║                                                                            ║
╚════════════════════════════════════════════════════════════════════════════╝
```

### 📐 What are Strides?

In a flat `data` slice, you need to know how to **translate multidimensional indices**
into a single linear offset. That’s what `strides` are for.

**Definition**:
> `strides[i]` = how many elements to skip to move along axis `i`

### 🧪 Example:

For a shape `[3, 4]`, your 2D data layout is:

```
Row 0: a00 a01 a02 a03
Row 1: a10 a11 a12 a13
Row 2: a20 a21 a22 a23
```

This flattens to:

```go
data := []float64{
  a00, a01, a02, a03,
  a10, a11, a12, a13,
  a20, a21, a22, a23,
}
```

And the strides would be:

```go
strides := []int{4, 1}
```

So to access `a[2][3]`, compute:

```go
index := 2*4 + 3*1 = 11
value := data[11]
```

---

### ⚡ Why Use Strides?

- 🔄 Efficient `reshape()` and `transpose()` without copying memory
- 📦 Enables slices, views, and broadcasting
- ⚙️ One array → multiple virtual representations

### 💡 Access Formula (general N-D):

```go
index := 0
for i, coord := range indices {
    index += coord * strides[i]
}
return data[index]
```

---

🛡️ **Design Principle**: Separate the *view* (shape + strides) from the *data* (flat memory).
This gives you immense flexibility with zero-cost abstractions.


---

✍️ **Tip:** Keep expanding this wiki with notes, tricks, and patterns as you develop!

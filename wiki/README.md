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

## 🔧 Understanding `make` in Go

`make` is a built-in Go function used to initialize **composite types**: slices, maps, and channels. Unlike `new`, which allocates memory but doesn't initialize, `make` returns a **ready-to-use value** of the correct type.

---

### 🔹 Syntax

```go
make(type, length[, capacity])
```

| Type       | Use `make`? | Example                         |
|------------|-------------|---------------------------------|
| Slice      | ✅ Yes       | `make([]int, 5)`                |
| Map        | ✅ Yes       | `make(map[string]int)`          |
| Channel    | ✅ Yes       | `make(chan int, 3)`             |
| Struct     | ❌ No        | use `new(MyStruct)` or literal  |
| Array      | ❌ No        | use `[N]T{}` or `var arr [N]T`  |

---

### 📌 Slices

```go
s := make([]int, 5)
```

- Creates a slice of 5 integers, all zero-initialized.
- Length and capacity are both 5.

```go
s := make([]int, 2, 10) // len=2, cap=10
```

Allows the slice to grow efficiently up to 10 elements.

---

### 📌 Maps

```go
m := make(map[string]int)
m["score"] = 100
```

- Required before writing to a map.
- Avoids nil map panics.

---

### 📌 Channels

```go
ch := make(chan string, 3)
```

- Creates a buffered channel with capacity 3.
- Can be used immediately in goroutines.

---

### 🧠 `make` vs `new`

| Feature        | `make`                     | `new`                       |
|----------------|----------------------------|-----------------------------|
| Allocates?     | ✅ Yes                     | ✅ Yes                      |
| Initializes?   | ✅ Yes (for slices, maps)  | ❌ No (just zero value)     |
| Returns        | Value                      | Pointer                     |
| Use with       | Slices, Maps, Channels     | Structs, Arrays             |

---

💡 **Pro Tip**:
Use `make` whenever you need a working slice, map, or channel — it avoids `nil` references and is the idiomatic way in Go.

---

## 🧠 Pointers in Go: When to Use `*`, `&`, or Neither

Go uses pointers to reference memory addresses, giving you low-level control without the complexity of manual memory management. Here's a simple breakdown to clarify when and why you'd use `*`, `&`, or neither.

---

### 🔹 `*` (Pointer declaration / dereference)

#### 1. As a **type** (declare a pointer)
```go
var p *int // p is a pointer to an int
```

#### 2. To **dereference** (access the value)
```go
fmt.Println(*p) // Get the value at the memory address p
```

---

### 🔹 `&` (Address-of operator)

Use `&` to get the memory address of a value.

```go
x := 10
p := &x // p points to x
```

---

## 🔄 When to use each

| Use Case                             | Symbol     | Example                           | Why?                                  |
|-------------------------------------|------------|-----------------------------------|----------------------------------------|
| You want a pointer to a value       | `&`        | `p := &x`                         | To avoid copying or to mutate original |
| You want to dereference a pointer   | `*`        | `val := *p`                       | To access the pointed value            |
| You want to declare a pointer var   | `*` in type| `var p *MyStruct`                | To store an address                    |
| You want to define method receiver  | `*`        | `func (s *Struct) Do()`          | To allow mutation                      |
| You don’t need to mutate or share   | None       | `func (s Struct) Copy()`         | Value copy is fine                     |

---

## 🔧 In method receivers

```go
type NDArray struct {
    data []float64
}

func (a NDArray) Copy() NDArray {
    return a // works on a copy
}

func (a *NDArray) Set(i int, val float64) {
    a.data[i] = val // modifies the original
}
```

🔸 Use a **pointer receiver** (`*NDArray`) when:
- The method **modifies** the struct
- You want to avoid **copying** on every call (performance)

---

## ⚠️ Beware!

- `*` is used for **dereferencing** and also in **type declarations**
- `&` is used to get a variable’s **memory address**
- You usually combine both: `*p := &x`

---

## 🧠 Summary

| Situation                        | Use |
|----------------------------------|-----|
| Access original (not a copy)     | `&` |
| Store address of a value         | `*` in type |
| Read from a pointer              | `*` before variable |
| Write with modification intent   | pointer receiver (`*Type`) |
| Just pass by value (copy)        | no symbol needed |

---

✍️ **Tip:** Keep expanding this wiki with notes, tricks, and patterns as you develop!

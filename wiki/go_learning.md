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

✍️ **Tip:** Keep expanding this wiki with notes, tricks, and patterns as you develop!

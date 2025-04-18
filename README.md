# Gondor - Scientific Computing Structures in Go

![Tests](https://github.com/arnaizaitor/gondor/actions/workflows/test.yml/badge.svg)
[![codecov](https://codecov.io/gh/arnaizaitor/gondor/branch/master/graph/badge.svg)](https://codecov.io/gh/aitoringenieria/gondor)


<p align="center">
  <img src="./static/gondor_banner_3.png" alt="Bionicle Banner" width="1000"/>
</p>

**Gondor** is a from-scratch reimplementation of [NumPy](https://numpy.org/) in [Go](https://go.dev/), built for learning and experimentation.

> “A personal project to deeply learn Go by recreating the inner workings of a numerical computing library.”

---

## 🚀 What is this?

Gondor is an experimental library that replicates core features of NumPy:

- `NDArray` structure for multi-dimensional data
- Vectorized operations (`add`, `multiply`, `dot`, `sum`, etc.)
- Shape manipulation (`reshape`, `transpose`)
- Stride-based indexing
- (Planned) Support for broadcasting, generic types, and more

All implemented **from scratch, with no external dependencies**, to gain a true understanding of numerical array internals.

---

## ❗ Disclaimer

This project is **not meant to be a production-ready alternative to NumPy**.
It is a learning tool for anyone who wants to:

- Master Go through a real-world, technical challenge
- Understand numerical data structures at a low level
- Explore the design of scientific computing APIs
- Have fun building something from the ground up

---

## 📦 Current Project Structure

```
gondor/
│   .gitignore
│   go.mod
│   Makefile
│   README.md
│
├───cmd
│   └───gondor
│           main.go              # Usage examples
│
├───internal
│   └───ndarray                  # Core multidimensional array logic
│           ndarray.go
│           ndarray_test.go
│           ops.go
│           shape.go
│           utils.go
│
├───static
│       gondor_banner.png
│       gondor_banner_2.png
│       gondor_banner_3.png
│
└───wiki
        README.md                # Documentation and examples
```

---

## 👋 Want to Learn Go by Building NumPy?

Welcome to **Gondor**.

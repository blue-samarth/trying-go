# The Go Programming Language Guide

Welcome to this comprehensive Go programming guide. This document serves as a detailed reference for Go programmers of all skill levels, covering everything from basic language structures to advanced concurrency patterns.

## Table of Contents

1. [Introduction to Go](#introduction-to-go)
2. [Language Fundamentals](#language-fundamentals)
   - [Basic Structure](#basic-structure)
   - [Variables and Types](#variables-and-types)
   - [Control Flow](#control-flow)
3. [Data Structures](#data-structures)
   - [Arrays](#arrays)
   - [Slices](#slices)
   - [Maps](#maps)
4. [Functions](#functions)
   - [Basic Functions](#basic-functions)
   - [Multiple Return Values](#multiple-return-values)
   - [Variadic Functions](#variadic-functions)
   - [Closures](#closures)
5. [Structs and Pointers](#structs-and-pointers)
   - [Structs](#structs)
   - [Pointers](#pointers)
   - [Constructor Pattern](#constructor-pattern)
6. [Interfaces](#interfaces)
   - [Interface Basics](#interface-basics)
   - [Real-World Example: Payment Systems](#real-world-example-payment-systems)
7. [Generic Data Structures](#generic-data-structures)
   - [Stack Implementation](#stack-implementation)
   - [Queue Implementation](#queue-implementation)
8. [Concurrency](#concurrency)
   - [Goroutines](#goroutines)
   - [Synchronization](#synchronization)
9. [Setup and Environment](#setup-and-environment)

## Introduction to Go

Go (also called Golang) is a statically typed, compiled language designed at Google. It provides numerous benefits that make it an excellent choice for modern development:

- **Compilation Speed**: Go compiles directly to machine code, resulting in extremely fast builds and startup times.
- **Concurrency Model**: Go's concurrency model with goroutines provides Python-like coroutines but with much better parallelism capabilities across multiple CPU cores.
- **Simplicity**: Go emphasizes simplicity and readability, with a small but powerful standard library.
- **Static Typing**: Go is statically typed, which helps catch errors at compile time rather than runtime.

## Language Fundamentals

### Basic Structure

Every Go program starts execution from the `main` function in the `main` package:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

### Variables and Types

Go supports various ways to declare variables:

```go
// Variable declaration
var name string = "Gopher"
age := 5  // Short declaration, type inferred

// Basic types
var (
    integer int = 42
    float float64 = 3.14
    text string = "Hello"
    boolean bool = true
)
```

### Control Flow

Go has unique syntax characteristics for control flow structures:

#### Loop Syntax

Go only has `for` loops (no `while` or `do-while`), but they can be used in multiple ways:

```go
// While-like Loop
var i int = 0
for i <= 5 {
    fmt.Println(i)
    i += 1
}

// Traditional C-style Loop
for i := 5; i >= 0; i-- {
    fmt.Println(i)
}

// Range-based Loop
for i := range 5 {
    fmt.Println(i)  // Prints 0, 1, 2, 3, 4
}
```

#### Conditional Statements

In Go, an `else` or `else if` statement must begin on the same line as the closing curly brace of the previous block:

```go
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}
```

#### Complex Example with Conditionals

```go
for i := range 10 {
    if i%2 == 0 {
        fmt.Println(i)  // Print even numbers
    } else if i & 1 == 1 && i/5 == 0 {
        // Note: else if must start on the same line as the closing brace
        fmt.Printf("You've got yourself a special number %d \n", i)
    } else if i & 1 == 1 || i/5 == 0 {
        fmt.Printf("You've got an even special number wow %d \n", i)
    } else {
        // Type conversion example for arithmetic with different types
        // fmt.Println(i * 2.5)  // This is not allowed because i is an int and 2.5 a float
        fmt.Println(float64(i) * 2.5)  // Type conversion makes this legal
    }
}
```

Important notes:
- Bitwise operations (`&`, `|`) with 1 to check odd/even: `i & 1 == 1` checks if number is odd
- Unlike some languages, Go requires explicit comparison with 1, not just `i & 1`
- Type conversion is required for mixed-type arithmetic: `float64(i) * 2.5` is valid

## Data Structures

### Arrays

Arrays in Go have a fixed size that cannot change during runtime:

```go
vals := [10]float64{1, 2.3, 4}
```

This creates a fixed-size array of 10 elements. We only provide values for the first three positions (1, 2.3, and 4), and Go automatically fills the remaining seven positions with zeros.

Key characteristics of arrays:
- Fixed size that cannot change during runtime
- Memory efficiency due to contiguous allocation
- Constant time access to any element by index
- Pass-by-value semantics (copying the entire array when passed to functions)

### Slices

Slices are flexible, dynamic-length views into arrays. They are one of the most commonly used data structures in Go.

#### Nil Slice

```go
var nums []int
```

This declares a slice without initializing it. It creates a nil slice with no backing array. Both its length and capacity are zero. When checking `cap(nums)` and `len(nums)`, both will return 0.

#### Creating with make()

```go
var nums2 = make([]int, 2, 5)
```

This uses the `make()` function to create a slice with:
- Initial length of 2 (contains two zeros initially)
- Capacity of 5 (can grow to 5 elements without reallocation)

The slice has a backing array of size 5, but only the first 2 elements are currently accessible. When printed, it will show `[0 0]`.

#### Empty Slice Literal

```go
nums3 := []float64{}
```

This creates an empty slice using literal syntax. Unlike the nil slice, this has a backing array of length 0. It's initialized but contains no elements.

#### Key Differences Between Arrays and Slices

**Arrays**:
- Fixed size that cannot change
- Type includes the size: `[10]float64`
- Passed by value (copying the entire array)

**Slices**:
- Dynamic size that can grow
- Reference type (points to a backing array)
- Has both length and capacity properties
- More flexible and commonly used in Go

### Maps

Maps in Go are unordered collections of key-value pairs, similar to dictionaries in Python or hash tables in other languages:

```go
mapy := make(map[string]int) // map[keyType]valueType
```

Keys are unique and must be of a type that is comparable (e.g., string, int, bool). Values can be of any type.

#### Adding and Accessing Elements

```go
mapy["a"] = 1
fmt.Println(mapy["a"]) // Output: 1
```

If a key exists, it returns its value. If not, it returns the zero value of the value type (e.g., 0 for int).

#### Deleting Elements

```go
delete(mapy, "a")
```

This removes the key-value pair with key "a". No error is raised if the key doesn't exist.

#### Map Literal Initialization

```go
secondMap := map[string]int{
    "x": 42,
    "y": 99,
}
```

This instantiates and populates a map in one line, often used when the values are known at compile time.

#### Clearing a Map

```go
clear(secondMap)
```

This removes all key-value pairs. The map remains allocated; length becomes 0. Note that this was introduced in Go 1.21 and is not available in earlier versions.

#### Checking if a Key Exists

```go
val, ok := mapy["b"]
if ok {
    fmt.Println("Key exists with value:", val)
}
```

The `ok` variable is a boolean indicating whether the key was found. This pattern is preferred over checking for zero value when a key may be absent.

#### Key Characteristics of Maps

- **Unordered**: Iteration order is not guaranteed.
- **Reference Type**: Passed by reference. Modifications affect the original.
- **Zero Value**: A nil map behaves like an empty map on read, but writing to it causes a panic.
- **Automatic Growth**: Maps grow automatically; capacity is not explicitly defined.

## Functions

### Basic Functions

Functions in Go start with the `func` keyword, followed by the function name, parameters, and return type:

```go
func add(a int, b int) int {
    return a + b
}
```

Parameters are declared as `name type`, and the return type follows the parameter list. Go supports only named returns or return values at the end — no keyword-based returns like in Python.

### Multiple Return Values

Go supports multiple return values, a core feature used heavily in error handling and unpacking:

```go
func getLanguages() (string, string) {
    return "Python", "Go"
}

// Usage with tuple unpacking
lang1, lang2 := getLanguages()
```

### Variadic Functions

Functions that accept a variable number of arguments:

```go
func summation(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Call like:
result := summation(1, 2, 3)
```

The `...` syntax allows passing zero or more arguments of the specified type. Internally, it's treated as a slice.

### Closures

Closures are functions defined within other functions that capture variables from their surrounding lexical scope:

```go
func outer() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

Each call to `outer()` returns a new closure with its own `sum` state. Closures are useful for:
- Creating stateful functions
- Implementing memoization and caching
- Event handlers, callbacks, and functional constructs

## Structs and Pointers

### Structs

Structs are value types that group related fields together:

```go
type order struct {
    id        string
    price     float64
    quantity  int
    createdAt time.Time
}
```

Think of structs as Go's version of classes — but without inheritance or traditional polymorphism. You can instantiate structs in several ways:

```go
order1 := order{}                             // default zero values
order2 := order{id: "123", price: 10.0, quantity: 2}
order2.createdAt = time.Now()
```

Use dot (.) notation to access or modify fields.

### Pointers

Pointers store the memory address of another variable:

```go
func changeNum(num *int) {
    *num = 10
}

// Usage
num := 100
changeNum(&num)
fmt.Println(num) // 10
```

Key pointer concepts:
- `*int` → pointer to an integer
- `*num` → dereference: access the value at the address
- `&num` → reference: get the address of num

Pointers are useful for:
- Avoiding copying large structs
- Modifying original data from within functions
- Enabling memory-efficient manipulation

### Constructor Pattern

Go doesn't have constructors like OOP languages. Instead, use custom initializer functions:

```go
func NewOrder(id string, price float64, quantity int) *order {
    return &order{id: id, price: price, quantity: quantity}
}
```

This returns a pointer to a new order instance, keeping initialization logic clean and reusable.

## Interfaces

### Interface Basics

In Go, an interface is a type that defines a set of method signatures. Any type (struct) that implements all the methods defined in the interface is said to satisfy that interface — without needing to explicitly declare it.

```go
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}

// Now Dog implements Speaker
// You can pass a Dog to a function that takes a Speaker:
func makeSpeak(s Speaker) {
    s.Speak()
}
```

Interfaces are useful for:
- **Decoupling**: Separating logic from implementation
- **Plug-and-play Design**: Swapping components easily
- **Testing**: Making it easy to mock implementations during unit testing
- **Scalability**: Adding new types without changing existing code

### Real-World Example: Payment Systems

Here's a practical example of interfaces for a payment processing system:

```go
// 1. Define Two Payment Providers
type razorpay struct{}
func (r *razorpay) processPayment(amount float64) {
    fmt.Println("Processing payment of", amount, "through Razorpay")
}

type stripe struct{}
func (s *stripe) processPayment(amount float64) {
    fmt.Println("Processing payment of", amount, "through Stripe")
}

// 2. Define an Interface to Standardize Them
type paymenter interface {
    processPayment(amount float64)
}

// 3. Build a Generic Payment Handler Using the Interface
type payment struct {
    gateway paymenter
}

func (p *payment) processPayment(amount float64) {
    p.gateway.processPayment(amount)
    fmt.Println("Processing payment of", amount, "through Payment Gateway")
}

// 4. Use It in main()
func main() {
    stripePaymentGateway := &stripe{}
    newPayment := payment{
        gateway: stripePaymentGateway,
    }
    newPayment.processPayment(100.0)
}
```

This results in:
```
Processing payment of 100 through Stripe
Processing payment of 100 through Payment Gateway
```

The magic here is that `gateway` can be swapped between Stripe, Razorpay, or any other future payment gateway — and this code will still work without modification.

## Generic Data Structures

Go 1.18+ introduced generics, allowing for type-safe, reusable data structures.

### Stack Implementation

```go
type Stack[T any] struct {
    data []T
}

func (s *Stack[T]) Push(value T) {
    s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if s.IsEmpty() {
        return zero, errors.New("stack is empty")
    }
    last := len(s.data) - 1
    val := s.data[last]
    s.data = s.data[:last]
    return val, nil
}

func (s *Stack[T]) Peek() (T, error) {
    var zero T
    if s.IsEmpty() {
        return zero, errors.New("stack is empty")
    }
    return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) Len() int {
    return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.data) == 0
}

func (s *Stack[T]) Clear() {
    s.data = nil
}
```

Key Go syntax concepts used:
- Generics: `type Stack[T any]`
- Method receiver: `func (s *Stack[T])` — pointer to the type
- Multiple return values: `func Pop() (T, error)`
- Zero value of T: `var zero T`
- Error handling: `return zero, errors.New(...)`
- Built-in functions: `len(...)`, `append(...)`

### Queue Implementation

```go
type Queue[T any] struct {
    data []T
}

func (q *Queue[T]) Push(value T) {
    q.data = append(q.data, value)
}

func (q *Queue[T]) Pop() (T, error) {
    var zero T
    if q.IsEmpty() {
        return zero, errors.New("Queue is empty")
    }
    popped := q.data[0]
    q.data[0] = zero // Optional: prevent memory leaks
    q.data = q.data[1:]
    return popped, nil
}

func (q *Queue[T]) Peek() (T, error) {
    var zero T
    if q.IsEmpty() {
        return zero, errors.New("Queue is empty")
    }
    return q.data[0], nil
}

func (q *Queue[T]) Len() int {
    return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
    return len(q.data) == 0
}

func (q *Queue[T]) Clear() {
    q.data = nil
}
```

For production use, you might want to make the queue thread-safe by wrapping it with a mutex:

```go
type Queue[T any] struct {
    mu   sync.Mutex
    data []T
}

// Each method then wraps access like:
func (q *Queue[T]) Push(value T) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.data = append(q.data, value)
}
```

## Concurrency

### Goroutines

Goroutines are lightweight threads managed by the Go runtime:

```go
type Number interface {
    ~int | ~float64
}

func tasks[T Number](id T) {
    fmt.Printf("Task %v started\n", id)
}

func main() {
    for i := 0; i < 11; i++ {
        go tasks(i) // Start a new goroutine for each task

        go func(i int) { 
            fmt.Printf("Task %v ended\n", i) 
        }(i) // Start a new goroutine using closure
    }
    
    time.Sleep(1 * time.Second) // Wait for goroutines to finish (not a good practice)
}
```

Key points:
- `go tasks(i)` creates a goroutine for the `tasks` function
- Each iteration of the loop launches a new concurrent task that executes independently
- The anonymous goroutine captures the value of `i` and prints a message concurrently

### Synchronization

In the example above, `time.Sleep(1 * time.Second)` is used to wait for goroutines to finish. However, this is not a good practice for production code. Instead, use proper synchronization mechanisms like `sync.WaitGroup`:

```go
func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 11; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            tasks(i)
        }(i)
    }
    
    wg.Wait() // Properly wait for all goroutines to complete
}
```
Got it — here's the properly **consistent continuation** of your Markdown documentation, matching the style and formatting of the snippet you provided:

### WaitGroups

WaitGroups are used to wait for a collection of goroutines to finish executing. This ensures all concurrent tasks complete before moving on.

```go
type Number interface {
    ~int | ~float64
}

func tasks2[T Number](id T, w *sync.WaitGroup) {
    defer w.Done() // Decrement the WaitGroup counter when the goroutine completes
    fmt.Printf("Task %v started in a WaitGroup\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 11; i++ {
        go tasks(i) // Unsynchronized goroutine

        go func(i int) {
            fmt.Printf("Task %v ended\n", i)
        }(i)

        wg.Add(1)
        go tasks2(i, &wg) // Start goroutine with WaitGroup tracking
    }

    wg.Wait() // Block until all goroutines call Done()
    fmt.Println("All tasks started")
}
```

Key points:

* `sync.WaitGroup` is a struct used to synchronize goroutines.
* `wg.Add(1)` increments the counter before starting a goroutine.
* `defer wg.Done()` decrements the counter when the goroutine finishes.
* `wg.Wait()` blocks until the counter is zero, meaning all goroutines have completed.
* The WaitGroup must be passed by reference (`&wg`) so all goroutines share the same counter.

> ✅ Prefer `sync.WaitGroup` over `time.Sleep` for precise and reliable goroutine synchronization.


## Concurrency with Channels and WaitGroups in Go

This document illustrates how to use **channels** and **WaitGroups** in Go to manage concurrent tasks efficiently and avoid common pitfalls like deadlocks.

---

### 🔁 Overview

This example covers:

* Creating **unbuffered** and **buffered** channels
* Launching **goroutines** to process data concurrently
* Using **sync.WaitGroup** to wait for all tasks to complete
* **Closing channels** properly to avoid deadlocks

---

### 🔌 Channels: Unbuffered vs Buffered

```go
numChan1 := make(chan any)        // Unbuffered channel
numChan2 := make(chan any, 5)     // Buffered channel with capacity 5
```

* **Unbuffered Channel (`numChan1`)**:

  * Blocks the sender until the receiver reads the value
  * Risk of deadlock if a value is sent without an active receiver

* **Buffered Channel (`numChan2`)**:

  * Allows multiple values to be sent up to buffer limit without blocking
  * Useful for temporary storage before processing

---

### ⚙️ WaitGroup for Synchronization

```go
var wg sync.WaitGroup
wg.Add(2) // We have 2 goroutines
```

* `Add(n)` specifies the number of goroutines to wait for
* Each goroutine calls `defer wg.Done()` when finished
* `wg.Wait()` blocks the main function until all goroutines complete

---

### 🧵 Goroutines with Channels

```go
func processNum(numChan chan any, wg *sync.WaitGroup) {
    defer wg.Done()
    for num := range numChan {
        fmt.Println("Processing number:", num)
    }
}
```

This function runs as a goroutine. It reads values from the channel in a loop. The `range` stops when the channel is **closed**.

---

### 🚀 Main Function Execution

```go
go processNum(numChan1, &wg)
go processNum2(numChan2, &wg)

numChan1 <- 42
numChan1 <- 3.14
numChan2 <- 9.81
numChan1 <- "Hello"
numChan2 <- "World"
numChan1 <- true
numChan2 <- false

close(numChan1)
close(numChan2)

wg.Wait()
```

* Values of different types are sent using `chan any`
* Channels are **closed** after all values are sent
* `wg.Wait()` ensures main doesn't exit prematurely

---

### ⚠️ Common Pitfalls

| Issue           | Cause                                               | Solution                                          |
| --------------- | --------------------------------------------------- | ------------------------------------------------- |
| Deadlock        | Sending to an unbuffered channel without a receiver | Launch a goroutine first or use buffered channels |
| Runtime Panic   | Sending to a closed channel                         | Always close after all sends                      |
| Stuck Goroutine | Not closing channel                                 | Close channels when done sending                  |

---

### ✅ Key Takeaways

* Use **WaitGroup** to manage goroutine completion
* Close **channels** to avoid blocking consumers
* Prefer **buffered channels** if producers might outpace consumers
* `chan any` provides flexibility, but use with care (type assertions may be needed in real applications)

---

## Setup and Environment

To permanently add Go to your PATH, add these lines to your profile file:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

A common mistake when setting up Go is making temporary environment changes. The above ensures Go is always available in your environment.

---

This documentation aims to provide a comprehensive guide to Go programming. For more detailed information, refer to the official Go documentation at [golang.org](https://golang.org).
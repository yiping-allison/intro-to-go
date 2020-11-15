# Introduction

Welcome to a simple introduction to Go!

The goal of this repository is to show
how Go is different from traditional object-oriented programming languages and highlight
some of the cooler features in Go.

## What is Go?

Go is a programming language developed by Google. It's designed to be syntatically simpler than C
while having almost the same computational power.

Here are some other facts:

* Statically-typed
* Compiled language
* Memory safety (garbage collector)
* First-class concurrency support (multi-threading)
* **Go 2 is in development and brings new features such as Generics**

## Go Types, Syntax and Control Flow

### Basic Types

* bool
* string
* int
  * int 8
  * int 32
  * int 64
* uint
  * uint 8
  * uint 32
  * uint 64
* byte
* rune
* float
  * float 32
  * float 64
* complex
  * complex 32
  * complex 64

You can convert types in Go like you would in Python:

```python
# Python type conversion
num = 10
num_float = float(num)
```

```go
// Go type conversion
num := 10
num_float := float64(num)
```

:bulb: `:=` in Go is a **variable assignment expression**. It is equivalent to the following:

```go
var num int
num = 10
```

On the backend, Go does this through type inference (not covered in this intro). Read more about
[type inference here](https://tour.golang.org/basics/14).

Constants are prefixed with the `const` keyword. Constants can be of strings, bools, or numeric.

```go
const greeting = "Hello World!"
```

### Function Syntax

A typical Go function looks like this:

```go
func greeting(msg string, name string) string {
    return fmt.Sprintf("Hello %s! %s", name, msg)
}
```

:bulb: **Several things to note here :bulb:

* The return type is specified after the function name and parameters
* Parameter types must be given right after the parameter name: `msg string` & `name string`
* If we have multiple parameters of the same type, we can actually shorten it to: `(msg, name string)`

Functions can return multiple values. You will se functions with multiple return values a lot in Go.
It especially commonly used for error handling.

```go
func returnMultVals(x, y, z int) (int, int, int) {
    return x, y, z
}

func createUser(name string, age int) (*User, error) {
    if goodVals(name, age) {
        return &User{Name: name, Age: age}, nil
    }
    return nil, fmt.Errorf("invalid user error")
}
```

### Control Flow

#### Loops

Go only has one loop type: **the For loop**. The **for** loop can be used both iteratively and as a **while** loop.

```go
// simple indexed for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while loop
for {
    // this loop will go on forever
    if myCondition {
        // you can stop it by including a break condition
        break
    }
}
```

#### `If-Else` Statements



## Go and Object Oriented Programming

### Object-Oriented Programming

**What is OOP?**

OOP is a programming style which organizes software around data objects.

i.e. objects contains both the data and code and can modify its own data

There are 4 tenets of OOP:

1. Encapsulation
2. Polymorphism
3. Inheritance
4. Abstraction

We'll take a look at each of the tenets and see how these concepts are implemented in-depth.

## Go and Memory-Related Operations

Go has pointers. Pointers in Go, however, are simpler and much easier to work with than in
traditional systems-level programming languages (C/C++/Rust).

Go also has a native garbage collecter, so you don't have to worry about memory deletion.

:exclamation: NOTE :exclamation:

Go by default comes with a garbage collector, but it's also possible to disable this feature if you
want to write hand-optimized Go code.

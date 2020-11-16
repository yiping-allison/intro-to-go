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

:bulb: **Several things to note here** :bulb:

* The return type is specified after the function name and parameters
* Parameter types must be given right after the parameter name: `msg string` & `name string`
* If we have multiple parameters of the same type, we can actually shorten it to: `(msg, name string)`

Functions can return multiple values. You will see functions with multiple return values a lot in Go.
It is especially common for error handling.

```go
// generic multi-value return
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

Go only has one loop type: **the For loop**. The **for** loop can be used iteratively with an indexer and as a **while** loop.

```go
// simple indexed for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while loop
for {
    // this loop will go on forever
    if myCondition {
        // you can exit the loop by including a break condition
        break
    }
}
```

#### `If-Else` Statements

If/Else statements are similar to other programming languages'.

```go
if x < 10 {
    fmt.Println("Less than 10")
} else {
    fmt.Println("Greater or equal to 10")
}
```

Go also has a switch statement for larger if/else blocks.

```go
x := 10
switch {
case x < 5:
    fmt.Println("Less than 5")
case x > 5:
    fmt.Println("Greater than 5")
default:
    fmt.Println("Equal to 5")
}
```

Unlike C/C++, Go does not require break statements in the switch block.

## Go Essentials

### Maps

```go
// create our map which maps strings to an integer
myMap := make(map[string]int)
// set some values
myMap["one"] = 1
myMap["two"] = 2

// iterate over all our map values and print the key and its value
for k, v := range myMap {
    fmt.Printf("Key: %s, Value: %d\n", k, v)
}

// check if our map has the key: three
if _, ok := myMap["three"]; ok {
    fmt.Println("Our map has 3")
} else {
    fmt.Println("Our map does not have 3")
}
```

:bulb: Like Python, Go has the `_` symbol to denote that we will not be using a particular value.
:bulb: Maps have a validation bool - _ok_. **ok** will be set true if our condition returns true, else false

### Slices

Slices are dynamic arrays in Go. They're used much more often than arrays.

```go
// slice that defines a set size
primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
fmt.Println(s)

// unset size
even := []int{2, 4, 6, 8, 10}
// appending to slices
append(even, 12)
```

### Structs

Structs are like objects in traditional OOP languages. They represent a collection of fields.

```go
type Book struct {
    Author string
    Title string
    Price int
}
```

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

Go has pointers.

```go
type Book struct {
    Author string
    Title string
    Price int
}

func main() {
    book := &Book{
        Author: "J.K. Rowling",
        Title: "Harry Potter and the Deathly Hallows",
        Price: 20,
    }

    // Go will automatically dereferences struct fields
    fmt.Println(book.Author)
    // when you print the object, you can see it's a pointer to a struct
    fmt.Println(book)
}
```

Go also has a native garbage collecter, so you don't have to worry about memory deletion.

:exclamation: NOTE :exclamation:

Go by default comes with a garbage collector, but it's also possible to disable this feature if you
want to write hand-optimized Go code.

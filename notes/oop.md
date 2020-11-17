# Object Oriented Programming

## Encapsulation

Go acheives encapuslation by exported and unexported functions. Encapsulation is implemented on
the package level. For example, it's possible to hide methods, structs, and constants within
a specific package by using lowercase leading characters. Uppercase leading characters signify that
the method/variable/struct is publicly usable by methods from other packages.

```go
// only available to current package
func notExported() bool {
    return false
}

// available to all packages
func ThisIsExported() bool {
    return true
}
```

Notice that this is not possible in Python as all methods and classes are by default public.
You can usually signify private methods and files by prepending the method or filename with a `_`.

For example, the equivalent of a private method in python would be:

```python
def _my_private_method() -> bool:
    return True
```

Most IDEs (including JetBrain) will automatically remove prepended `_` methods from the module autocomplete list.

## Abstraction

Abstraction and encapsulation goes pretty much hand-in-hand. Encapsulation - most of the time - helps implement abstraction.

For example, suppose we have a package which handles database operations. Before we send the request to the database,
we need to sanitize user inputs or check connection settings, etc.

Most of the time, a third-party user doesn't need to know the nitty gritty details, so in Go, these are implemented as
unexported methods.

```go
package database

func New() *Database {
    return &Database{}
}

func (d *Database) Get(input string) *Entry {
    validate(input)
    return d.get(input)
}

func (d *Database) get(input string) *Entry {
    return d.sql(input)
}

func validate() bool {
    if some_condition {
        return true
    }
    return false
}
```

## ~Inheritance??~ Composition

Unlike traditional OOP languages, Go does not have inheritance. Object relationships are handled by composition.
In other words, objects can have loosely coupled relationships with each other. Loose coupling is also commonly known
as `is a` relationship in system design.

```go
type Car struct {
    wheel Wheel
    window Window
}

type Wheel struct {...}

type Window struct {...}
```

From the example above, we can see that a Car object is comprised of both a `wheel` and a `window`. Both the wheel and the
window are two separate objects which when grouped together, help make a car.

Notice how this type of relationship doesn't try to connect the direct relationship between a car and a wheel or window.

This type of embedding is known as **struct embedding** in Go. Go also supports embedding for interfaces.

```go
type Calc interface {
    SimpleArith
    MultDiv
}

type SimpleArith interface {
    add(x, y int) int
    subtract(x, y int) int
}

type MultDiv interface {
    multiply(x, y int) int
    divide(x, y int) int
}
```

It's generally recommended to keep interfaces small and add them to larger structures when needed. This makes
it easier to maintain large and complex systems.

## Polymorphism

Go supports polymorphism through receiver types. As mentioned before, Go objects are represented as structs. We can use
these structs and define a method which applies solely to that object.

```go
type Book struct {
    name string
    title string
    author string
}

func (b Book) getName() string {
    return b.name
}

type Cat struct {
    name string
    owner string
}

func (c Cat) getName() string {
    return c.name
}
```

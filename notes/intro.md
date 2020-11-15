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

We'll take a look at each of the tenets and see how these concepts are implemented in Go.

## Go and Memory-Related Operations

Go has pointers. Pointers in Go, however, are simpler and much easier to work with than in
traditional systems-level programming languages (C/C++/Rust).

Go also has a native garbage collecter, so you don't have to worry about memory deletion.

:exclamation: NOTE :exclamation:

Go by default comes with a garbage collector, but it's also possible to disable this feature if you
want to write hand-optimized Go code.

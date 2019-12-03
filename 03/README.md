# Ducking errors

## Objectives
1. Learn about method receivers
1. Learn about interfaces
1. Learn about errors

## Talk about the weather / Get settled (5 minutes)

.. 

## Dive into interfaces and errors (15 minutes)

But first..!

### The method receiver

- Adding behavior to `data`

```go
type SuperHero struct{
    Catchphrase string
}

func (s SuperHero) Speak() string {
    return s.Catchphrase
}
```

.. which is syntax sugar for

```go
func Speak(s SuperHero) string {
    return s.Catchphrase
}
```

- Deciding when to use a method receiver or a parameter depends on the situation:
    - Does the type need to satisfy an interface? Receiver.
    - Does the type really add any value? No? Parameter.
        - Package name gives the syntax of a type, without a type. (e.g. `mypackage.MyMethod(myparameter)`)

- NOTE: Avoid stutter. `superhero.NewSuperHero()` vs `superhero.New()`

### The interface

- Common approaches to naming interfaces:
    - If single method interface, name the interface the name of the method with `er` appended to it.
    - Otherwise, an active noun describing what it can do.

..more naming while we're at it:
- `struct`: A plain ol' boring noun (Person, Document, ...)
- `interface`: An active noun (Reader, Processor, ...)
- `function`: A verb (Read, Process, Sync, ...)

```go
type Speaker interface {
    Speak() string
}
```

- Definition includes method name, parameters, and return type
- Any type that satisfies the interface can be used. Any.. type. From anywhere.

```go
type SuperHero struct{
    Catchphrase string
}

func (s SuperHero) Speak() string {
    return s.Catchphrase
}
```

```go
type Dog struct {}
func (d Dog) Speak() string {
    return "woof"
}
```

```go
func Speak(s Speaker) string {
    return s.Speak()
}
```

- Commonly referred to as _duck typing_. Technically, it's `struct` typing as it's done at compile time. But you'll hear both.

- As the _Go_ saying goes.. *accept interfaces, return structs*.
    - common exception: Factories

**KEY TAKEAWAY: DO NOT START WITH INTERFACES**
- Introduce interfaces as you need them. 
- Introduce them in the file where you need them. Not near the type.
- Typically create interfaces when using other packages. Packages _generally_ don't expose interfaces.

### The error

- Errors are just types that satisfy the `error` interface.
```go
type error interface {
    Error() string
}
```

So when you see..
```go
func MyMethod() (int, error) {
    ...
}
```

`error` is actually an _interface_.

## Problem (25 minutes)

1. Write a function that uses [https://github.com/jpreese/slowimage](https://github.com/jpreese/slowimage) to download an image. 

    - Your function should return an **error** if the file that was downloaded is a blank string.
    - Your function should return the **name of the file** that was downloaded (if not blank).

1. Print the result of the function to the console.

1. Write a unit test to make sure that your function behaves as expected.

### Stuck?

**Getting started**
- [The previous lessons](https://github.com/jpreese/go-learn) contain a lot of information on how to create a new project, and create test files.

- You _will_ need to create a `go.mod` file

**The program**

- Your program will need to create a new `Image` provided by the `slowimage` library. Feel free to use whatever Filename in the struct when creating it.

- To import the `slowimage` library, include `"github.com/jpreese/slowimage"` in your import statement block. Similar to `fmt`.

- To access the members of the package (after importing it), use the `slowimage.` prefix.

- **Your function will most likely need to take a parameter.**

- Your function should have a return type of `(string, error)`

- To return an error, you can use `errors.New()`. 
    - e.g. `return "", errors.New("your error message")`

**Testing**

- **You'll want to create a new type**. Using `Image` is not recommended.

- Your test should make sure that the function that you created, returns the correct filename.

## Review and Discussion (15 minutes)

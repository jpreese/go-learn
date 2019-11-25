# Lets get slicing

## Objectives
1. Learn about Go memory allocation
1. Learn about the `array` type
1. Learn about the `slice` type

## Talk about the weather / Get settled (5 minutes)

.. 

## Introducing the Slice (15 minutes)

### Go memory allocation

- Go does not have a CLR or JVM. It runs on your machine.
- Deterministic memory allocation and access.
- *Example: The order that variables are declared in Go, matter (if you care).*

```go
// allocate 24 bytes
type MyStruct struct {
	firstBool  bool
	firstInt   int64
	secondBool bool
}
```

```go
// allocate 16 bytes
type MyStruct struct {
	firstBool  bool
	secondBool bool
	firstInt   int64
}
```

- This showcases what Go calls *alignment*. 
- Determistic memory allocation for the target platform and faster runtime performance (at the cost of memory... technically).

### The `array`

- Arrays and their size are a type

```go
func MyFuncTakingAnArray(myArray [42]int)
```

- Limited use. "Don't use an array unless you know what you are doing." *Example: You need to specifically need to plan out memory locations.*

### The `slice`

```go
func MyFuncTakingASlice(mySlice []int)
```

- Slices are an abstraction over an array. A window of an array.
- References an `array` under the hood. Multiple slices can reference the same array!
- Fast to copy. A struct containing `Len`, `Cap`, and a pointer to the underlying `array`.

- Need to add some elements? Use `append()` to add elements to a slice.

```go
mySlice = append(mySlice, "somevalue") // assume slice of string
```

*Fun fact: In Go, `string` is a read-only slice of `byte`*

## Problem (30 minutes)

Write a function that takes a `slice` of integers and returns how many times the number `100` is found in the slice.

Print the result to the console, and prove your functionality works with a unit test.

### Steps

*REMINDER: If you're saving your work to a GitHub repo, don't forget to put your source in your repositories folder!*

1. Create a module. `go mod init yourmodulehere`
1. Create a `.go` file for your `main` function and counting function

*NOTE: To write this program, you'll need to use `range`. Usage can be found here https://tour.golang.org/moretypes/16*. `range` is used to iterate over collections like `array`, `slice`, and `map`.

### Stuck?

- For getting started, you can refer to the [first lesson](../01/README.md). This goes over into more detail on how to start a new Go project and set up tests.

- You'll need to create a new module, `.go` file to put your code into, and a new `_test.go` file to put your tests.

- The solution and full project set up for the previous lesson can be found [here](../01/solution).

**The program**

- Your function should take a `slice` of `int` and return an `int`.
- Use `range` to check if the current value is `100` and if so, count it.
- Return the total count.

**Testing**
- Your test should have a `slice` of `int` that you create with some numbers, some of which should be `100`.

- Your test should have a variable called `expected` of type `int` which is the number of times you expect the number 100 to appear in the list you passed in.

- Your test should compare if `expected` is equal to `actual` (the number returned from your function).

- Remember that `t.Errorf()` should be used to report a failed test.

## Review and Discussion (10 minutes)

TBD

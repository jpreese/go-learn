# Lets get slicing

## Objectives
1. Learn about Go memory allocation
1. Learn about the `array` type
1. Learn about the `slice` type

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
- Determistic memory allocation for the target platform and faster runtime performance (at the cost of memory allocation.. technically).
- Driving reason behind lack of other primatives (linkedlist, graphs, etc).

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

*Fun fact: In Go, `string` is a (read-only) slice of `byte`*

## Problem (35 minutes)


## Review and Discussion (10 minutes)

TBD

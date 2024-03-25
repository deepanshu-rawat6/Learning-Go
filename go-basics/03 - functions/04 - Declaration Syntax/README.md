# Declaration Syntax

Developers often wonder why the declaration syntax in Go is different from the tradition established in the C family of languages.

## C-Style syntax

The C language describes types with an expression including the name to be declared, and states what type that expression will have.

```go
int y;
```

The code above declares `y` as an `int`. In general, the type goes on the left and the expression on the right.

Interestingly, the creators of the Go language agreed that the C-style of declaring types in signatures gets confusing really fast - take a look at this nightmare.

```go
int (*fp)(int (*ff)(int x, int y), int b)
```

## Go-style syntax

Go's declarations are clear, you just read them left to right, just like you would in English.

```go
x int
p *int
a [3]int
```

It's nice for more complex signatures, it makes them easier to read.

```go
f func(func(int,int) int, int) int
```

## REFERENCE

The [following post on the Go blog](https://blog.golang.org/declaration-syntax) is a great resource for further reading on declaration syntax.

## Questions:

### What are we talking about when we discuss 'declaration syntax'?

* The style of language used to create new variables, types, functions, etc...
* The ever-important question of tabs vs spaces
* The decision about whether to use camelCase or snake_case
* The argument over guard clauses vs if-else statements

`The style of language used to create new variables, types, functions, etc...`

### Which language's declaration syntax reads like English from left-to-right

* Go
* C

`Go`

### What is 'f func(func(int,int) int, int) int'?

* A function named 'f' that takes an int as the argument and returns an int
* A function named 'f' that takes a function as the argument and returns an int
* A function named 'f' that takes a function and an int as arguments and returns an int
* A function named 'f' that takes a function and an int as arguments and returns a function

`A function named 'f' that takes a function and an int as arguments and returns an int`
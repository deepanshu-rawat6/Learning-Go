# COMPILING EXPLAINED

Computers don't know how to do anything unless we as programmers tell them what to do.

Unfortunately, computers don't understand human language. They don't even understand uncompiled computer programs.

For example, the code:

```go
package main

import "fmt"

func main() {
  fmt.Println("hello world")
}
```

means nothing to a computer.

## COMPUTERS NEED MACHINE CODE

A computer's [CPU](https://en.wikipedia.org/wiki/Central_processing_unit) only understands its own _instruction set_, which we call "machine code".

Instructions are basic math operations like addition, subtraction, multiplication, and the ability to save data temporarily.

For example, an [ARM processor](https://en.wikipedia.org/wiki/ARM_architecture) uses the ADD instruction when supplied with the number `0100` in binary.

## GO, C, RUST, ETC.

Go, C, and Rust are all languages where the code is first converted to machine code by the compiler before it's executed.

![compilation](https://i.imgur.com/uqnMubj.png)

## Questions

### Do computer processors understand English instructions like 'open the browser'?

`No`
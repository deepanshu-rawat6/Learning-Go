# GO PROGRAMS ARE EASY ON MEMORY

Go programs are fairly lightweight. Each program includes a small amount of "extra" code that's included in the executable binary. This extra code is called the [Go Runtime](https://go.dev/doc/faq#runtime). One of the purposes of the Go runtime is to clean up unused memory at runtime.

In other words, the Go compiler includes a small amount of extra logic in every Go program to make it easier for developers to write code that's memory efficient.

## COMPARISON

As a general rule, Java programs use more memory than comparable Go programs. There are several reasons for this, but one of them is that Java uses an entire virtual machine to interpret bytecode at runtime. Go programs are compiled into machine code, and the overhead of the Go runtime is typically less than the overhead of the Java virtual machine.

On the other hand, Rust and C programs use slightly _less_ memory than Go programs because more control is given to the developer to optimize the memory usage of the program. The Go runtime just handles it for us automatically.

IDLE MEMORY USAGE

![idle-memory-usage](https://miro.medium.com/max/1400/1*Ggs-bJxobwZmrbfuoWGpFw.png)

In the chart above, [Dexter Darwich compares the memory usage](https://medium.com/@dexterdarwich/comparison-between-java-go-and-rust-fdb21bd5fb7c) of three very simple programs written in Java, Go, and Rust. As you can see, Go and Rust use very little memory when compared to Java.


## Questions:

### Generally speaking, which language uses the most memory?

* Rust
* Java
* Go
* C

`Java`

### What's one of the purposes of the Go runtime?

* To style Go code and make it easier to read
* To scare off JavaScript developers
* To compile Go code
* To cleanup unused memory

`To cleanup unused memory`

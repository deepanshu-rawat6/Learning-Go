# COMPILED VS INTERPRETED

You can run a compiled program _without_ the original source code. You don't even need a compiler. For example, when your browser executes the code you write in this course, it doesn't use the original code, just the compiled executable.

This is different than interpreted languages like Python and JavaScript. With Python and JavaScript, the code is interpreted at [runtime](https://en.wikipedia.org/wiki/Runtime_(program_lifecycle_phase)) by a separate program known as the "interpreter". Distributing code for users to run can be a pain because they need to have an interpreter installed, and they need access to the source code.

### EXAMPLES OF COMPILED LANGUAGES

* Go
* C
* C++
* Rust

### EXAMPLES OF INTERPRETED LANGUAGES

* JavaScript
* Python
* Ruby

![COMPILED VS INTERPRETED](https://i.imgur.com/ovHaWmS.jpg)

## WHY BUILD TEXTIO IN A COMPILED LANGUAGE?

One of the most convenient things about using a compiled language like Go for Textio is that when we deploy our server we don't need to include any runtime language dependencies like Node or a Python interpreter. We just add the pre-compiled binary to the server and start it up!

## Questions

### Which do users of a compiled program need to run the program?

* The source code
* The compiler
* The compiled executable
* Blue hair

`The compiled executable`

### Which language is interpreted?

* Rust
* Python
* Go
* C++

`Python`
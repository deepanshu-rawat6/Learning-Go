# SHORT VARIABLE DECLARATION

Inside a function (like the main function) the `:=` short assignment statement can be used in place of a [var](https://go.dev/tour/basics/9) declaration. The `:=` operator infers the type of the new variable based on the value. It's colloquially called the walrus operator because it looks like a walrus... sort of.

These two lines of code are equivalent:

```go
var empty string
```

```go
empty := ""
```

```go
numCars := 10 // inferred as an integer
temperature := 0.0 // temperature is inferred as a float because it has a decimal
var isFunny = true // inferred as a boolean
```

Outside of a function (in the [global/package scope](https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables)), every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

## ASSIGNMENT

Many of our users send birthday messages using the Textio API. Finish the `main` function so it prints "Happy birthday! You are now 21 years old!".

First, create the first half of the message. Use a walrus operator to declare a string named `messageStart` with the value "Happy birthday! You are now".

Next, use another walrus operator to declare an integer named `age` with the value `21`.

Finally, declare another string named `messageEnd` with the value "years old!".

You should get a nicely formatted message that smashes the variables together.
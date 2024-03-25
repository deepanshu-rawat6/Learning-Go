# COMPUTED CONSTANTS

Constants must be known at compile time. They are usually declared with a static value:

```go
const myInt = 15
```

However, constants _can be computed_ as long as the computation can happen at _compile time_.

For example, this is valid:

```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

That said, you _cannot_ declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

```go
// the current time can only be known when the program is running
const currentTime = time.Now()
```

## ASSIGNMENT

Keeping track of time in a message-sending application is _critical_. Imagine getting an appointment reminder an hour **after** your doctor's visit.

Complete the code using a computed constant to print the number of seconds in an hour.
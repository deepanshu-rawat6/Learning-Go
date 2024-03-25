# CONDITIONALS

`if` statements in Go do not use parentheses around the condition:

```go
if height > 4 {
    fmt.Println("You are tall enough!")
}
```

`else if` and `else` are supported as you might expect:

```go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
```

## ASSIGNMENT

Fix the bug on line `12`. The `if` statement should print "Message sent" if the `messageLen` is _less than or equal_ to the `maxMessageLen`, or "Message not sent" otherwise.

## TIPS

Here are some of the comparison operators in Go:

* `==` equal to
* `!=` not equal to
* `<` less than
* `>` greater than
* `<=` less than or equal to
* `>=` greater than or equal to
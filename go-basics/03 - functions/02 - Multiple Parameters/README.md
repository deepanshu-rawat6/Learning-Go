# MULTIPLE PARAMETERS

When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

Here are some examples:

```go
func addToDatabase(hp, damage int) {
  // ...
}
```

```go
func addToDatabase(hp, damage int, name string) {
  // ?
}
```

```go
func addToDatabase(hp, damage int, name string, level int) {
  // ?
}
```

## Questions:

### Which of the following is the most succinct (and valid) way to write a function signature?

* func createUser(firstName, lastName string, age int)
* func createUser(firstName, age, lastName string int)
* func createUser(firstName string, lastName string, age int)
* func createUser(firstName string, age int, lastName string)

`func createUser(firstName, age, lastName string int)`
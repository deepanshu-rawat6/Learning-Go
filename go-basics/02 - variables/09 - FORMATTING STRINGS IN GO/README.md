# FORMATTING STRINGS IN GO

Go follows the [printf tradition](https://cplusplus.com/reference/cstdio/printf/) from the C language. In my opinion, string formatting/interpolation in Go is _less_ elegant than Python's f-strings, unfortunately.

* [fmt.Printf](https://pkg.go.dev/fmt#Printf) - Prints a formatted string to [standard out](https://stackoverflow.com/questions/3385201/confused-about-stdin-stdout-and-stderr).
* [fmt.Sprintf()](https://pkg.go.dev/fmt#Sprintf) - Returns the formatted string

These following "formatting verbs" work with the formatting functions above:

## DEFAULT REPRESENTATION

The `%v` variant prints the Go syntax representation of a value, it's a nice default.

```go
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old
```

If you want to print in a more specific way, you can use the following formatting verbs:

## STRING

```go
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old
```

## INTEGER

```go
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
```

## FLOAT

```go
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```

If you're interested in all the formatting options, you can look at the `fmt` package's [docs](https://pkg.go.dev/fmt#hdr-Printing).

## ASSIGNMENT

Create a new variable called `msg` on line 9. It's a string that contains the following:

```
Hi NAME, your open rate is OPENRATE percentNEWLINE
```

Where `NAME` is the given `name`, `OPENRATE` is the `openRate` rounded to the nearest "tenths" place, and `NEWLINE` is the [\n](https://en.wikipedia.org/wiki/Newline) escape sequence.

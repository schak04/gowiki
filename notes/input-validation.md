# Input validation with regular expressions

When user input is used for things like file paths, URLs, queries, etc., validate it before using it.

Go's `regexp` package provides regular expressions:

```go
var validPath = regexp.MustCompile(`^/(edit|save|view)/([a-zA-Z0-9]+)$`)
```

- `regexp.Compile` returns `(*Regexp, error)`.
- `regexp.MustCompile` panics if the regex is invalid, otherwise returns `*Regexp`.
- `MustCompile` is useful for regexes known at program startup.

`FindStringSubmatch` matches the string and returns the full match plus captured groups:

```go
m := validPath.FindStringSubmatch(path)
```

If there's no match, it returns `nil`.

A common validation pattern:

```go
m := validPath.FindStringSubmatch(input)
if m == nil {
    // invalid input
}
```

For reusable validation logic, put it in a separate function that returns the validated value and an error:

```go
value, err := validate(input)
if err != nil {
    return
}
```

The important idea is:

```text
untrusted input
      |
   validate
   /      \
invalid   valid
  |         |
reject     use
```

**Never trust user input just because it came through your own HTTP endpoint.**

# Function literals and closures

### Function literal

A function can be created as a value without giving it a name:

```go
func(x int) int {
    return x * 2
}
```

This is a function literal. Functions can be passed around like other values.

### Closure

A closure is a function that captures variables from its surrounding scope:

```go
func makeHandler(fn func(...)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fn(w, r)
    }
}
```

The returned function captures `fn`, so it can still access it after `makeHandler` returns.

Think:

```text
makeHandler(fn)
      |
      v
returns function
      |
      v
function still has access to fn
```

### Why use this pattern?

A wrapper can handle common logic once:

```text
request
   |
   v
wrapper
   |
   +-- validate input
   |
   +-- common error handling
   |
   v
actual handler
```

This avoids repeating the same logic in every handler.

### Key idea

**Function literals let us create functions as values. Closures let those functions remember variables from their surrounding scope.**

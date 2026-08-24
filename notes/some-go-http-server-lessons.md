# Go HTTP Server: Key Concepts

## `net/http`

- `http.HandleFunc(path, handler)` registers a handler for a path.
- `http.ListenAndServe(":8080", nil)` starts the server and blocks.
- `http.Redirect(w, r, url, code)` sends a redirect, usually `302`.
- `http.Error(w, msg, code)` sends an HTTP error response.
- `http.NotFound(w, r)` sends a `404 Not Found`.

## Handlers

```go
func handler(w http.ResponseWriter, r *http.Request)
```

- `w`: interface used to write the HTTP response.
- `r`: pointer to the incoming HTTP request.
- `http.HandleFunc` expects exactly this function signature.

## `makeHandler` / closures

```go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc
```

Used to share common logic between handlers.

```text
request
  -> validate path
  -> extract title
  -> actual handler
```

The returned function is a **closure** because it captures `fn` from the surrounding scope.

```go
makeHandler(viewHandler)
```

creates a function that remembers `viewHandler` and later calls:

```go
fn(w, r, title)
```

## Regex capture groups

For:

```text
^/(edit|save|view)/([a-zA-Z0-9]+)$
```

```text
m[0] = entire match
m[1] = edit/save/view
m[2] = page title
```

`FindStringSubmatch` returns `nil` if there is no match.

## Templates

- `ParseFiles` parses templates.
- `Execute` executes a template.
- `ExecuteTemplate` executes a specific named template from a collection.
- Parse once and reuse templates for efficiency.
- `template.Must` panics if initialisation fails.

## File I/O

- `os.ReadFile` -> `[]byte, error`
- `os.WriteFile` -> `error`
- `[]byte` represents raw byte data, including file contents.

## HTTP status

- `302 Found` -> temporary redirect; `Location` tells the client where to go.
- `404` -> requested resource/path not found.
- `500` -> server-side error.

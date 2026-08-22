# `net/http`

`net/http` is Go's standard library for building HTTP servers and handling HTTP requests.

Basic flow:

```go
http.HandleFunc("/pathName/", handler)
log.Fatal(http.ListenAndServe(":8080", nil))
```

- `HandleFunc` registers a handler for a URL path.
- `ListenAndServe` starts the server and blocks while it's running.
- `log.Fatal` handles the error if the server stops unexpectedly.

A handler looks like:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

- `w` is used to write the HTTP response.
- `r` contains information about the incoming request.
- `r.URL.Path` gives us the requested URL path.

For example:

```text
/view/test
```

We can strip `/view/` with:

```go
title := r.URL.Path[len("/view/"):]
```

giving us:

```text
test
```

Multiple routes can simply be registered with multiple `HandleFunc` calls:

```go
http.HandleFunc("/view/", viewHandler)
http.HandleFunc("/edit/", editHandler)
http.HandleFunc("/save/", saveHandler)
```

**register routes -> receive requests -> handle them -> write responses**

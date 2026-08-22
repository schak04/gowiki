# `http.ResponseWriter` vs `*http.Request`

```go
func handler(w http.ResponseWriter, r *http.Request)
```

- `w` is an **interface**, so we pass the interface value. It still refers to the underlying response writer.
- `r` is a **pointer** to the `http.Request`, so we avoid copying the whole request and work with the actual request object.

Basically:

w -> interface describing the response writer
r -> pointer to the actual request

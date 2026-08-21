## `io.Reader` and `io.Writer`

Go interfaces for byte streams:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

- `io.Reader`: something we can **read bytes from**
- `io.Writer`: something we can **write bytes to**

Examples: files, network connections, buffers, stdin/stdout, HTTP bodies.

Basically:
io.Reader -> data comes out
io.Writer <- data goes in

### Similar to C++:

```text
std::istream -> io.Reader
std::ostream -> io.Writer
```

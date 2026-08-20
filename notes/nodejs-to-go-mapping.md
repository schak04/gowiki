## Node.js -> Go mental mapping since I already know backend dev:

| Node.js         | Go                                                                            |
| --------------- | ----------------------------------------------------------------------------- |
| Express app     | `http.Server` / `http.ServeMux`                                               |
| Route           | `mux.HandleFunc(...)`                                                         |
| Controller      | handler function                                                              |
| `req`           | `*http.Request`                                                               |
| `res`           | `http.ResponseWriter`                                                         |
| `req.query.q`   | `r.URL.Query().Get("q")`                                                      |
| `res.json(...)` | `json.NewEncoder(w).Encode(...)`                                              |
| Middleware      | function wrapping a handler                                                   |
| Service layer   | ordinary Go functions/types                                                   |
| Model           | structs + persistence code                                                    |
| Router          | `http.ServeMux` or another router                                             |
| `app.listen()`  | `server.ListenAndServe()`                                                     |
| npm package     | Go module dependency                                                          |
| `async/await`   | usually synchronous functions; goroutines when concurrency is actually needed |

The important difference is that Go's standard library gives us most of this directly. There is no need of an Express-like library to create a normal HTTP API.

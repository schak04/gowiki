# Handling pages, redirects and errors

## Non-existent pages

If a page doesn't exist, don't continue with empty data. Redirect to the edit page:

```go
if err != nil {
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    return
}
```

`http.Redirect` sends a `302` response with a `Location` header.

## Saving pages

`saveHandler` gets the form data, creates a `Page`, saves it, then redirects to the view page:

```go
body := r.FormValue("body")
p := &Page{Title: title, Body: []byte(body)}
err := p.save()
```

`FormValue` returns a `string`, so we convert it to `[]byte` for `Page.Body`.

## Error handling

Don't ignore errors with `_` when they can actually matter.

Use `http.Error` to send an error response:

```go
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```

`http.StatusInternalServerError` is HTTP `500`.

So the general pattern is:

```text
do something
    |
    v
error?
 /   \
yes   no
 |     |
500   continue
```

Handling errors explicitly makes the server behave predictably instead of silently continuing with bad/missing data.

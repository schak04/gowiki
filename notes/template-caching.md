# Template caching

Previously, `ParseFiles` ran every time we rendered a page. That's unnecessary work.

Instead, parse all templates **once at startup**:

```go
var templates = template.Must(
    template.ParseFiles("edit.html", "view.html"),
)
```

Now we can reuse them:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

- `ParseFiles` parses the templates once.
- `ExecuteTemplate` executes an already-parsed template.
- `template.Must` panics if parsing fails. That's fine here because the server can't work without its templates.

**Basically:** parse once, reuse many times. This avoids unnecessary work on every request.

# `html/template`

`html/template` lets us keep HTML separate from Go code.

Instead of hard-coding HTML in a Go program, we put it in a template:

```html
<h1>{{.Title}}</h1>
```

Then pass our data to it:

```go
t, _ := template.ParseFiles("view.html")
t.Execute(w, p)
```

- `ParseFiles` loads the `.html` file into a `*template.Template`.
- `Execute` fills the template using the provided data and writes the result to `w`.
- `.Title` means `p.Title`, `.Body` means `p.Body`.

`{{ ... }}` is template syntax.

We can avoid repeating the template code by making a helper:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}
```

Then:

```go
renderTemplate(w, "view", p)
```

loads `view.html`.

`html/template` also automatically escapes HTML, which helps prevent user-provided data from breaking the page or injecting HTML.

**Basically:** Go handles the logic, templates handle the HTML.

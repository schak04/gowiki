# `Execute` vs `ExecuteTemplate`

Both execute a parsed template and write the result to an `io.Writer`.

### `Execute`

```go
t.Execute(w, data)
```

Executes the **template itself**.

Useful when `t` represents the template you want to render.

### `ExecuteTemplate`

```go
templates.ExecuteTemplate(w, "view.html", data)
```

Executes a **specific named template** from a collection of parsed templates.

Useful when multiple templates have been parsed together:

```go
templates := template.Must(
    template.ParseFiles("view.html", "edit.html"),
)
```

Then choose which one to execute:

```go
templates.ExecuteTemplate(w, "view.html", data)
```

### In short

```text
Execute         -> execute this template
ExecuteTemplate -> execute this named template from a collection
```

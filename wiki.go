// ref: https://go.dev/doc/articles/wiki

package main

import (
	// "errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

// The function template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the *Template unaltered.
// A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is exit the program.
// The ParseFiles function takes any number of string arguments that identify our template files, and parses those files into templates that are named after the base file name. If we were to add more templates to our program, we would add their names to the ParseFiles call's arguments.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// The function regexp.MustCompile will parse and compile the regular expression, and return a regexp.Regexp. MustCompile is distinct from Compile in that it will panic if the expression compilation fails, while Compile returns an error as a second parameter.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte // a byte slice because files are fundamentally raw bytes, not strings
	// []byte lets Go represent arbitrary file contents: text, images, binary, files, JSON, etc.
}

// save is a method instead of a plain function because it operates on a page, and it being a method makes the relationship clearer
// if an operation belongs to a specific type, it is good to make it a method
// no specific reason for using a pointer receiver instead of a value receiver, it isn't necessary here.
func (p *Page) save() error { // since return type of WriteFile is error
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // 0600 = read write perms for the current user only. The leading 0 means octal notation.
}

// loadPage is a function instead of a method because there isn't an existing Page to operate on yet
// title -> find/load Page -> return Page
// that is why it being a function makes sense
// no specific reason for return *Page instead of just Page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename) // ReadFile returns []byte and error
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Old comments (pre-template caching):
// The function template.ParseFiles will read the contents of edit.html and return a *template.Template.
// The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter. The .Title and .Body dotted identifiers refer to p.Title and p.Body.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		// If the title is invalid, the function will write a "404 Not Found" error to the HTTP connection, and return an error to the handler.
// 		http.NotFound(w, r)
// 		return "", errors.New(invalid Page Title)
// 	}
// 	// If the title is valid, it will be returned along with a nil error value.
// 	return m[2], nil // The title is the second subexpression.
// }

// The closure returned by makeHandler is a function that takes an http.ResponseWriter and http.Request (in other words, an http.HandlerFunc).
// The closure extracts the title from the request path, and validates it with the validPath regexp. If the title is invalid, an error will be written to the ResponseWriter using the http.NotFound function. If the title is valid, the enclosed handler function fn will be called with the ResponseWriter, Request, and title as arguments.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// handles URLs prefixed with "/view/"
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// The http.Redirect function adds an HTTP status code of http.StatusFound (302) and a Location header to the HTTP response.
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// r.FormValue returns a string which needs to be converted to []byte
	// before it fits into the Page struct.
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)} // string -> []byte
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // ListenAndServe blocks while the server is running. If it fails or stops, it returns an error. log.Fatal logs that error and exits.
}

// ref: https://go.dev/doc/articles/wiki

package main

import (
	"fmt"
	"os"
)

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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

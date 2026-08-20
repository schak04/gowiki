// ref: https://go.dev/doc/articles/wiki

package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte // a byte slice instead of a string slice because that is what the io libs expect
}

func (p *Page) save() error { // since return type of WriteFile is error
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // 0600 = read write perms for the current user only
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, err := os.ReadFile(filename) // ReadFile returns []byte and error
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

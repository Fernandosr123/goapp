package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "./cmd/app/data/" + p.Title + ".txt"
	//leer y escribir
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := "./cmd/app/data/" + title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	page := &Page{Title: title, Body: body}
	return page
}

func main() {
	page1 := &Page{Title: "title1", Body: []byte("fernando")}
	err := page1.save()
	if err != nil {
		fmt.Println(err)
	}

	page := loadPage("title1")
	fmt.Println(loadPage("title1"))
	fmt.Println(page.Title)
	fmt.Println(string(page.Body))

}

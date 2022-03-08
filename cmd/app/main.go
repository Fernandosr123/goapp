package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func loadPage(title string) (*Page, error) {
	filename := "./cmd/app/data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	page := &Page{Title: title, Body: body}
	return page, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>%s</h1>", err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	//http://localhost:8080/edit/articulo
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	fmt.Fprintf(w, `
		<html>
			<head>
				<title>%s</title>
			</head>
			<body>
				<h1>%s<h1>
				<form method="POST" action="/save/%s">
					<textarea name="body">%s</textarea>
					<button>Guardar</button>
				</form>
			</body>
		</html>
		`, page.Title, page.Title, page.Title, page.Body)

	//http://localhost/save
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	page.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	page1 := &Page{Title: "title1", Body: []byte("This is my firts .txt")}
	err := page1.save()
	if err != nil {
		fmt.Println(err)
	}

	//page, _ := loadPage("title1")
	//fmt.Println(loadPage("title1"))
	//fmt.Println(page.Title)
	//fmt.Println(string(page.Body))

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)

}

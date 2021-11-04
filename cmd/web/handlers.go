package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	files := []string{
		"ui/html/home.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(writer, "Internal server error", 500)
		return
	}
	err = ts.Execute(writer, nil)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(writer, "Internal server error", 500)
		return
	}
}

func (app *application) showSnippet(writer http.ResponseWriter, request *http.Request)  {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprintf(writer, "Отражение определенной заметки с ID %d...", id)
}

func (app *application) createSnippet(writer http.ResponseWriter, request *http.Request)  {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Метод не дозволен", 405)
		return
	}
	writer.Write([]byte("Создание заметки"))
}


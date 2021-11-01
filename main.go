package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, request *http.Request){
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	writer.Write([]byte("Привет из Snippetbox"))
}

//Обработчик отбражения заметки
func showSnippet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprintf(writer, "Отборажение выбранной заметки с ID %id...", id)
}

//Обработчик создания заметки
func createSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		//writer.WriteHeader(405)
		//writer.Write([]byte("GET метод запрещен"))
		http.Error(writer, "Метод запрещен", 405)
		return
	}
 		writer.Write([]byte("Создания заметки"))
}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
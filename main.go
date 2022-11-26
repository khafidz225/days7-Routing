package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })
	// route.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	// Route untuk menginisialisasi folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/projectDetail{id}", projectDetail).Methods("GET")
	route.HandleFunc("/add-project", addProjects).Methods("POST")

	fmt.Println("Server sedang berjalan pada port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/addProject.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func addProjects(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title : " + r.PostForm.Get("title"))
	fmt.Println("StartDate : " + r.PostForm.Get("startdate"))
	fmt.Println("EndDate : " + r.PostForm.Get("enddate"))
	fmt.Println("Description : " + r.PostForm.Get("description"))
	fmt.Println("nodejs : " + r.PostForm.Get("nodejs"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/profilDetail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	w.Write([]byte("Message : " + err.Error()))

	data := map[string]interface{}{
		"Title": "Pasar Coding dari Dumbways",
		"id":    id,
	}

	tmpt.Execute(w, data)
}

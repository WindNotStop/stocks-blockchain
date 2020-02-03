package main

import (
	"html/template"
	"log"
	"net/http"
	"stocks-blockchain/application/controller"
//	. "stocks-blockchain/application/model"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("application/view/mainPage.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, nil)
}

func query(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	err, res := controller.Query(r.Form.Get("uid"))
	if err != nil{
		log.Fatal(err.Error())
	}
	t1, err := template.ParseFiles("application/view/queryPage.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, res)
}

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/mainPage", mainPage)
	http.HandleFunc("/query", query)
	log.Fatal(server.ListenAndServe())
}
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hastechnologyltd/ordis-core/audit"
	"net/http"
)

func main() {

	audits := audit.CreateAudit()

	audits.AddAudit("Audit One")
	audits.AddAudit("Audit Two")
	audits.AddAudit("Audit Three")

	//var size = unsafe.Sizeof(audits)
	//fmt.Println(size)

	audits.Display()

	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	//router.HandleFunc("/todos", TodoIndex)
	//router.HandleFunc("/todos/{todoId}", TodoShow)
	//
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

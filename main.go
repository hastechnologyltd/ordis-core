package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hastechnologyltd/ordis-core/audit"
	"log"
	"net/http"
)

var audits *audit.Audits

func main() {

	audits = audit.CreateAudit()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	//router.HandleFunc("/audit", AddAudit)
	router.HandleFunc("/audit/{auditData}", AddAudit)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//func AddAudit(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Todo Index!")
//}

func AddAudit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auditData := vars["auditData"]
	audits.AddAudit(auditData)
	//fmt.Fprintln(w, "Todo show:", auditData)
	audits.Display()
}

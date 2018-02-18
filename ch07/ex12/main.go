package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type database map[string]int

const tmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>ch7 ex12</title>
  </head>
  <body>
    {{range $i, $v := .}}
      <p>{{$i}}: {{$v}}</p>
    {{end}}
  </body>
</html>
`

const errTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>ch7 ex12</title>
  </head>
  <body>
    {{.}}
  </body>
</html>
`

func main() {
	db := database{"socks": 5, "shoes": 50}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	showTmpl(w, db)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		showErrTmpl(w, db)
		return
	}
	num, err := strconv.Atoi(price)
	if err != nil {
		showErrTmpl(w, db)
		return
	}
	db[item] = num
	showTmpl(w, db)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		showErrTmpl(w, db)
		return
	}
	num, err := strconv.Atoi(price)
	if err != nil {
		showErrTmpl(w, db)
		return
	}
	db[item] = num
	showTmpl(w, db)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		showErrTmpl(w, db)
		return
	}
	delete(db, item)
	showTmpl(w, db)
}

func showTmpl(w http.ResponseWriter, db database) {
	report := template.Must(template.New("showList").Parse(tmpl))
	if err := report.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func showErrTmpl(w http.ResponseWriter, db database) {
	report := template.Must(template.New("error").Parse(errTmpl))
	if err := report.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

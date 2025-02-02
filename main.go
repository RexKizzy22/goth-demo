package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/RexKizzy22/goth-demo/components"
	"github.com/RexKizzy22/goth-demo/model"
	"github.com/RexKizzy22/goth-demo/public"
	"github.com/RexKizzy22/goth-demo/util"
)

var state model.State

const DEFAULT_SIZE = 15

func init() {
	stocks, err := util.ReadCsv()
	if err != nil {
		log.Fatal(err)
	}

	state = model.State{
		Rows: util.ParseStocks(&stocks),
	}
}

func main() {
	http.Handle("GET /static/", http.FileServer(http.FS(public.StaticFiles)))

	http.HandleFunc("GET /", index)

	http.HandleFunc("GET /table", table)

	http.HandleFunc("GET /search", search)

	http.HandleFunc("POST /add", add)

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	o := util.Paginate(r, DEFAULT_SIZE)
	components.Page(state.Rows[o:(o+DEFAULT_SIZE)]).Render(r.Context(), w)
}

func table(w http.ResponseWriter, r *http.Request) {
	o := util.Paginate(r, DEFAULT_SIZE)
	components.Table(state.Rows[o:(o+DEFAULT_SIZE)]).Render(r.Context(), w)
}

func search(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("search")
	if s == "" {
		http.Redirect(w, r, "/table", http.StatusSeeOther)
		return
	}

	o := util.Paginate(r, DEFAULT_SIZE)
	ns := util.FilterStocks(state.Rows, s)

	if len(ns) > DEFAULT_SIZE {
		components.Table(ns[o:(o+DEFAULT_SIZE)]).Render(r.Context(), w)
	} else {
		components.Table(ns).Render(r.Context(), w)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	nr := util.ParseForm(r)
	slices.Reverse(state.Rows)
	state.Rows = append(state.Rows, nr)
	slices.Reverse(state.Rows)
	o := util.Paginate(r, DEFAULT_SIZE)

	components.Table(state.Rows[o:(o+DEFAULT_SIZE)]).Render(r.Context(), w)
}

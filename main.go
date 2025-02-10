package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RexKizzy22/goth-demo/handler"
	"github.com/RexKizzy22/goth-demo/model"
	"github.com/RexKizzy22/goth-demo/public"
	"github.com/RexKizzy22/goth-demo/util"
)

var state model.State

func init() {
	state = util.LoadAppState()
}

func main() {
	handler := handler.New(&state)

	http.Handle("GET /static/", http.FileServer(http.FS(public.StaticFiles)))

	http.HandleFunc("GET /", handler.Index)
	
	http.HandleFunc("GET /refresh", handler.RefreshMain)

	http.HandleFunc("GET /table", handler.Table)
	
	http.HandleFunc("GET /main", handler.Main)

	http.HandleFunc("GET /slide/next", handler.NextSlide)

	http.HandleFunc("GET /slide/prev", handler.PrevSlide)

	http.HandleFunc("POST /add", handler.Add)

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

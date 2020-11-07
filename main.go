package main

import (
	"html/template"
	"net/http"
)

func main() {
	// マルチプレクサ
	mux := http.NewServeMux()

	// 静的ファイルのルーティング
	// http://localhost/static/css/bootstrap.min.css
	// -> <root>/css/bootstrap.min.css
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// ハンドラ関数
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}

}


package main

import (
	"flag"
	"html/template"
	"net/http"
	"os"
)

// templateHandler assigns each URL to the corresponding template structure.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	layout := "../templates/structure.html"
  var content string

	// TODO: don't hard code
	switch r.URL.Path {
	case "/":
		content = "../pages/index.html"
	case "/stories":
		content = "../pages/stories/index.html"
	case "/contribute":
		content = "../pages/contribute/index.html"
	case "/download":
		content = "../pages/download/index.html"
	case "/elsewhere":
		content ="../pages/elsewhere/index.html"
	case "/about":
		content ="../pages/about/index.html"                
	default:
		content = "../pages/" + r.URL.Path + ".html"
	}

	info, err := os.Stat(content)
	if err != nil {
		if os.IsNotExist(err) {
			errorHandler(w, r, http.StatusNotFound)
		} else {
			errorHandler(w, r, http.StatusBadRequest)
		}
		return
	}

	if info.IsDir() {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(layout, content)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "structure", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if (status == http.StatusNotFound) {
		layout := "../templates/structure.html"
		content := "../pages/notfound.html"
		tmpl, err := template.ParseFiles(layout, content)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		tmpl.ExecuteTemplate(w, "structure", nil)
	}
}

// main starts the web server and routes URLS.
func main() {
	listenAddress := flag.String("host", "localhost:8080", "address and port to listen on, type -host=yourdomain.com:80 for accessing online")

	flag.Parse()

	http.Handle("/public/images/", http.StripPrefix("/public/images/", http.FileServer(http.Dir("../public/images"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("../public"))))
	http.HandleFunc("/", templateHandler)

	http.ListenAndServe(*listenAddress, nil)
}

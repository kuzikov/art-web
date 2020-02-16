package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"text/template"
)

// uninit template struct
var tpl *template.Template

// init initialize stuct tpl
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", index)

	static := http.FileServer(http.Dir("public"))

	http.Handle("/public/", http.StripPrefix("/public/", static))

	http.ListenAndServe(":5500", nil)
}

// index func is handler for .root location
func index(w http.ResponseWriter, r *http.Request) {
	for i, v := range r.RequestURI {
		if i == 1 && v != '?' {
			GetInfo(w, "error.html" /*"ERROR 404"*/, http.StatusNotFound)
			return
		}
	}
	switch r.Method {
	case "GET":
		GetInfo(w, "index.html", nil)
	case "POST":

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
			return
		}
		text, err := url.ParseQuery(string(body))
		if err != nil {
			GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
			return
		}

		for i := range text {

			switch i {
			case "text":
				if len(text[i][0]) == 0 {
					GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
					return
				}
			case "selector":
				switch text[i][0] {
				case "standard", "shadow", "thinkertoy":
					continue
				// case "":
				// 	GetInfo(w, "error.html", "ERROR 400")
				default:
					GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
					return
				}
			default:
				GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
				return
			}
		}

		if len(text["text"]) != 1 || len(text["selector"]) != 1 {
			GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
			return
		}
		// newStr := strings.Replace(strings.Replace(text.Get("text"), "\r\n", "\\n", -1), "\r", "\\n", -1)
		if !fFind(text.Get("selector")+".txt", "public/generator/files/") || !fFind("ascii-art", "public/generator/") {
			GetInfo(w, "error.html" /*"ERROR 500"*/, http.StatusInternalServerError)
			return
		}

		art, err := exec.Command("./public/generator/ascii-art" /*newStr*/, text.Get("text"), text.Get("selector")).Output()
		if err != nil {
			GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
			return
		}
		GetInfo(w, "index.html", string(art))
	default:

		GetInfo(w, "error.html" /*"ERROR 400"*/, http.StatusBadRequest)
		return
	}
}

// fFind check existing file
func fFind(file, dir string) bool {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.Name() != file {
			continue
		} else {
			return true
		}
	}
	return false
}

// GetInfo func take inform from server and send it to browser
func GetInfo(w http.ResponseWriter, html string, info interface{}) {
	if html == "error.html" && (info.(int)) == 400 {
		w.WriteHeader(http.StatusBadRequest)
	}
	if html == "error.html" && (info.(int)) == 404 {
		w.WriteHeader(http.StatusNotFound)
	}
	if html == "error.html" && (info.(int)) == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err := tpl.ExecuteTemplate(w, html, info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		GetInfo(w, "error.html", "ERROR 500")
		return
		// http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, error2 string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	//?w.WriteBody(code)
	//w.Body(./public/error.html)
	//GetInfo(w, "error.html", "ERROR 400")
	//fmt.Fprintln(w, GetInfo(w , "error.html", "ERROR 404"))
	fmt.Fprintln(w, error2)
}

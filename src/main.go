package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type UploadedFile struct {
	Name string
	Body []byte
}

func (f *UploadedFile) save() error {
	return ioutil.WriteFile(f.Name, f.Body, 0600)
}

func main() {
	tmpl := template.Must(template.ParseFiles("src/result.json"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			return
		}

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var Buf bytes.Buffer

		file, header, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		name := header.Filename
		fmt.Printf("File name %s\n", name)

		io.Copy(&Buf, file)

		f := &UploadedFile{Name: name, Body: Buf.Bytes()}
		f.save()

		Buf.Reset()

		// TODO: change the command to actual one that executes genreXpose and let it print the result to stdout
		out, err := exec.Command("file", name).Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Remove(name)
		var Result string
		Result = string(out)
		Result = strings.TrimSuffix(Result, "\n")

		tmpl.Execute(w, struct {
			Success bool
			Output  string
		}{true, Result})
	})

	http.ListenAndServe(":8081", nil)
}

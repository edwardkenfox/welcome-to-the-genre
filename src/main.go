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
)

type UploadedFile struct {
	Name string
	Body []byte
}

func (f *UploadedFile) save() error {
	return ioutil.WriteFile(f.Name, f.Body, 0600)
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

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

		tmpl.Execute(w, struct {
			Success bool
			Output  string
		}{true, string(out)})
	})

	http.ListenAndServe(":8081", nil)
}

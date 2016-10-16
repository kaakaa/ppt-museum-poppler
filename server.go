package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func convertPDFtoText(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	fmt.Fprintf(w, "Hello World")

	if r.Method == "POST" {
		f, err := extractFile(w, r)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		t, err := pdf2text(f.Name())
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Fprintf(w, t)
	}
}

func extractFile(w http.ResponseWriter, r *http.Request) (*os.File, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)

	f, err := ioutil.TempFile("", "pm-poppler")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	io.Copy(f, file)

	return f, nil
}

func main() {
	http.HandleFunc("/", convertPDFtoText)
	http.ListenAndServe(":8080", nil)
}

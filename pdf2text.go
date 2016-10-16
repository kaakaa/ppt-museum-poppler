package main

import (
	"strings"

	poppler "github.com/cheggaaa/go-poppler"
)

func pdf2text(filepath string) (string, error) {
	doc, err := poppler.Open(filepath)
	if err != nil {
		return "", err
	}

	var result []string
	for i := 0; i < doc.GetNPages(); i++ {
		p := doc.GetPage(i)
		str := p.Text()
		result = append(result, str)
	}
	return strings.Join(result, " "), nil
}

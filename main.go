package main

import (
	"bufio"
	"log"
	"os"
	"sec-form4/form4"
	"strings"
)

func main() {
	data := &form4.FileContent{}

	err := extractForm4Data("./0000046619-22-000004.nc", data)
	if err != nil {
		os.Exit(1)
	}
	log.Printf("data: %v\n", form4.Serialize(data))

	os.Exit(0)
}

func extractForm4Data(filename string, data *form4.FileContent) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("read form4 error: %v", err)
		//return errors.New("open form4 error")
	}
	defer file.Close()

	xmlPart := false
	var d *form4.ExtractData
	var buffer, docBuffer []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if xmlPart == false {
			d = form4.GetParser(l, form4.ExtractingDocument)
			if d != nil {
				xmlPart = true
				buffer = append(buffer, l)
			} else {
				docBuffer = append(docBuffer, l)
			}
		} else if xmlPart {
			buffer = append(buffer, l)
			if strings.EqualFold(l, d.End) {
				xmlPart = false
				_ = d.Proc(buffer, data)
				buffer = nil
			}
		}
	}
	_ = form4.ExtractReportDocument(docBuffer, data)
	return nil
}

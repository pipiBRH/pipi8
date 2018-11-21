package execute

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

func ExecuteTemplate(templateInfo map[string]string, imageTag *string) {
	type Recipient struct {
		ImageTag *string
	}
	var recipients = Recipient{imageTag}

	for key, value := range templateInfo {
		t := template.Must(template.New(key).ParseFiles(value))
		var b bytes.Buffer
		err := t.Execute(&b, recipients)
		fmt.Println(b.String())
		if err != nil {
			log.Println("executing template:", err)
		}
		err = ioutil.WriteFile(value, []byte(b.String()), 0644)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

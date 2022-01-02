package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

func main() {

	templateFunctions := template.FuncMap{
		"asset": assetFileName,
	}
	baseTemplatePath := "./templates/base.gohtml"
	baseTemplate, err := template.New("base").Funcs(templateFunctions).ParseFiles(baseTemplatePath)
	if err != nil {
		log.Println(err.Error())
	}
	personTplFilename := "person-index.gohtml"
	personTplPath := "./templates/" + personTplFilename
	fileBuf, err := ioutil.ReadFile(personTplPath)
	baseName := strings.TrimSuffix(path.Base(personTplFilename), ".gohtml")
	log.Println("base trimmed name: " + baseName)

	s := string(fileBuf)
	log.Println("\n" + s)
	tpl2, err := template.Must(baseTemplate.Clone()).New(baseName).ParseFiles(personTplPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var tw bytes.Buffer
	err = tpl2.Execute(&tw, nil)
	if err != nil {
		log.Fatalln("execute error: " + err.Error())
	}
	log.Println(tw.String())
}

func assetFileName(fileName string) string {
	assetFileNames := map[string]string{}
	assetFileNames["css/styles.css"] = "css/styles.css"
	return "/" + assetFileNames[fileName]
}

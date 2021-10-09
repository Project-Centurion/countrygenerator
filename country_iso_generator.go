package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/dave/jennifer/jen"
)

func generateCode(packageName string, codeToName map[string]string) *File {
	f := NewFile(packageName)

	f.Type().Id("CountryCode").Id("string")
	f.Type().Id("CountryName").Id("string")

	f.Comment("CodeToName is a mapping between countries ISO codes and their corresponding english name.")
	f.
		Var().Id("CodeToName").
		Op("=").Map(Qual("", "CountryCode")).Qual("", "CountryName").
		Values(DictFunc(func(d Dict) {
			for code, country := range codeToName {
				d[Lit(code)] = Lit(country)
			}
		}))

	f.Line()

	f.Comment("NameToCode is a mapping between english country names and their corresponding ISO code.")
	f.
		Var().Id("NameToCode").
		Op("=").Map(Qual("", "CountryName")).Qual("", "CountryCode").
		Values(DictFunc(func(d Dict) {
			for code, country := range codeToName {
				d[Lit(country)] = Lit(code)
			}
		}))

	return f
}

func generateIsoCodeFile(packageName string) *File {
	resp, err := http.Get("http://country.io/names.json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var codeToName map[string]string
	err = json.Unmarshal(body, &codeToName)
	if err != nil {
		log.Fatalln(err)
	}

	return generateCode(packageName, codeToName)
}

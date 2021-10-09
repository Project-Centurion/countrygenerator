package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) < 2 {
		log.Fatalf("please provide a file path and a package name. format is countrygenerator file.go mypackage. file location must be relative to the current directory")
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	code := generateIsoCodeFile(argsWithProg[1])

	file, err := os.Create(fmt.Sprintf("%s/%s", path, argsWithProg[0]))
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	err = code.Render(file)
	if err != nil {
		log.Fatalln(err)
	}
}
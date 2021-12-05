package main

import (
	"fizz/internal/conf"
	"fizz/internal/io"
	"log"
	"path/filepath"
)

func main() {
	err := filepath.Walk(conf.ProjectRoot(), io.Walk)
	if err != nil {
		log.Println(err)
		panic("Failed to update some of the generated HTMLs.")
	}
}

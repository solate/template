package main

import (
	"github.com/solate/generation"
	"github.com/solate/template/virhis"
	"log"
)

func main() {

	err := virhis.SetConfig()
	if err != nil {
		log.Println(err)
	}

	generation.GenerateProjectCode()

}

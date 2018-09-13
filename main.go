package main

import (
	"github.com/solate/generation"
	"gitlab.dev.okapp.cc/med-his-group/Template/virhis"
	"log"
)

func main() {

	err := virhis.SetConfig()
	if err != nil {
		log.Println(err)
	}

	generation.GenerateProjectCode()

}

package main

import (
	"flag"
	"fmt"
	"github.com/xerardoo/invoice-cfdi-validator/internal/validator/app"
)

func main() {
	args := app.NewAppArgs()
	app.ReadConfig()

	flag.StringVar(&args.Filepath, "filepath", "", "Location of a CFDI in the file directory")
	flag.Parse()

	if args.Filepath == "" {
		fmt.Println("'-filepath' argument filepath is mandatory, more info see -help")
		return
	}

	err := app.Run(args)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

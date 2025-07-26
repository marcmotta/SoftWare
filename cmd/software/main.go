// cmd/software/main.go
package main

import (
	"flag"
	"log"
	"os"

	"software/internal/software"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := software.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

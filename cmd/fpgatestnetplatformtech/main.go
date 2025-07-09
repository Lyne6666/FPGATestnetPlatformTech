// cmd/fpgatestnetplatformtech/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"fpgatestnetplatformtech/internal/fpgatestnetplatformtech"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := fpgatestnetplatformtech.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

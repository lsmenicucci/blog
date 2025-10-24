package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	serveCmd := flag.Bool("serve", false, "start the dev server")
	exportCmd := flag.String("export", "", "export all posts to folder")
	flag.Parse()

	// loag config
	err := LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Failed to load config: %s\n", err)
		os.Exit(1)
	}

	if *serveCmd {
		serve()
		return
	}

	if *exportCmd != "" {
		err := Export(*exportCmd)
		if err != nil {
			log.Println("Export error:", err)
			os.Exit(1)
		}
		return
	}

	flag.Usage()
}

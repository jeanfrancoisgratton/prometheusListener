package main

import (
	"errors"
	"flag"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var mu sync.Mutex
var cfg Config_s

func main() {
	var ce *cerr.CustomError

	// Parse command-line flags
	setupFlag := flag.Bool("setup", false, "Run setup and exit")
	versionFlag := flag.Bool("version", false, "Show version")
	flag.Parse()

	// Check if the "-setup" flag is set
	if *setupFlag {
		// Call the setup function and exit
		if ce = setup(); ce != nil {
			ce.Error()
		} else {
			return
		}
	}
	// -version flag
	if *versionFlag {
		fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), hf.White(fmt.Sprintf("1.03.03-0-%s 2024.09.16", runtime.GOARCH)))
		os.Exit(0)
	}

	if _, err := os.Stat("/etc/prometheus/prometheusListener.json"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("The configuration file is absent, please run this tool with the -setup flag")
		os.Exit(0)
	}

	// Load the config file
	if cfg, ce = loadConfig(); ce != nil {
		ce.Error()
	}

	http.HandleFunc("/file", fileHandler)

	log.Printf("Starting server on :%d\n", cfg.Port)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", cfg.Port), cfg.Cert, cfg.Key, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

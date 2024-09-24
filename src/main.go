package main

import (
	"crypto/tls"
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
			fmt.Println(ce.Error())
		} else {
			return
		}
	}
	// -version flag
	if *versionFlag {
		fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), hf.White(fmt.Sprintf("2.00.00TEST-%s 2024.09.24", runtime.GOARCH)))
		os.Exit(0)
	}

	if _, err := os.Stat("/etc/prometheusSDlistener.json"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("The configuration file is absent, please run this tool with the -setup flag")
		os.Exit(0)
	}

	// Load the config file
	if cfg, ce = loadConfig(); ce != nil {
		fmt.Println(ce.Error())
	}

	// Ensure certFile and keyFile exist
	if _, err := os.Stat(cfg.Cert); os.IsNotExist(err) {
		log.Fatalf("Certificate file not found: %v", err)
	}
	if _, err := os.Stat(cfg.Key); os.IsNotExist(err) {
		log.Fatalf("Key file not found: %v", err)
	}

	// Setup HTTPS server
	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.Port),
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	// Start the HTTPS server
	log.Printf("Starting HTTPS server on port %d\n", cfg.Port)
	err := server.ListenAndServeTLS(cfg.Cert, cfg.Key)
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}

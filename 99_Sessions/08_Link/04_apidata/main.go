package main

import (
	"flag"
	"fmt"
	log1 "github.com/Sirupsen/logrus"
	"github.com/boseji/goboseji"
	"github.com/valyala/fasthttp"
	"os"
)

var (
	//addr        = flag.String("addr", "localhost:8080", "TCP Address:Port to listen to")
	compress    = flag.Bool("compress", false, "Whether to enable transparent response compression")
	HostAddress = ""
	logger      *log1.Logger
	logdir      string
	serverip    string
	serverport  string
)

func init() {

	// Start Commandline parsing
	flag.Parse()

	// Initialize the Logger
	logger = NewLogger()

	// Get Environment to Make the Host IP + Port info
	serverip = os.Getenv("OPENSHIFT_GO_IP")
	if len(serverip) == 0 {
		serverip = "localhost"
	}
	serverport = os.Getenv("OPENSHIFT_GO_PORT")
	if len(serverport) == 0 {
		serverport = "8080"
	}
	logdir = os.Getenv("OPENSHIFT_GO_LOG_DIR")
	if len(logdir) == 0 {
		logdir = ""
	}

	// Generate the Address Info
	HostAddress = fmt.Sprintf("%s:%s", serverip, serverport)
}

func main() {

	// Main Request handler
	h := mainRequestHandler
	if *compress {
		h = fasthttp.CompressHandlerLevel(h, fasthttp.CompressBestSpeed)
	}

	logger.Println("Starting Server At " + HostAddress)
	// Launch an independent Web Server
	go func() {
		if err := fasthttp.ListenAndServe(HostAddress, h); err != nil {
			logger.Fatalf("Error in ListenAndServe: %s", err)
		}
	}()

	// Set up Cleanup
	goboseji.SetIntrCleanup(cleanup, 0)

	// Select for all other Channels to run in Infinite Loop
	select {}
}

func cleanup() {
}

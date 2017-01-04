package main

import (
	_ "expvar"
	"flag"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/expvarhandler"
	"log"
)

var (
	addr     = flag.String("addr", "localhost:8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Enables transparent response compression if set to true")
)

func main() {
	// Parse command-line flags.
	flag.Parse()

	// Create RequestHandler serving server stats on /stats and files
	// on other requested paths.
	// /stats output may be filtered using regexps. For example:
	//
	//   * /stats?r=fs will show only stats (expvars) containing 'fs'
	//     in their names.
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/stats":
			expvarhandler.ExpvarHandler(ctx)
		default:
			ctx.Write([]byte("Hello World"))
		}
	}

	// Enable Compression
	if *compress {
		requestHandler = fasthttp.CompressHandler(requestHandler)
	}

	// Start HTTP server.
	if len(*addr) > 0 {
		log.Printf("Starting HTTP server on %q", *addr)
		go func() {
			if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServe: %s", err)
			}
		}()
	}

	log.Printf("See stats at http://%s/stats", *addr)

	// Wait forever.
	select {}
}

// Various counters - see https://golang.org/pkg/expvar/ for details.
var (
/*
	// Counter for total number of fs calls
	fsCalls = expvar.NewInt("fsCalls")

	// Counters for various response status codes
	fsOKResponses          = expvar.NewInt("fsOKResponses")
	fsNotModifiedResponses = expvar.NewInt("fsNotModifiedResponses")
	fsNotFoundResponses    = expvar.NewInt("fsNotFoundResponses")
	fsOtherResponses       = expvar.NewInt("fsOtherResponses")

	// Total size in bytes for OK response bodies served.
	fsResponseBodyBytes = expvar.NewInt("fsResponseBodyBytes")*/
)

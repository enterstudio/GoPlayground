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

}

func main() {

	// Main Request handler
	h := mainRequestHandler
	if *compress {
		h = fasthttp.CompressHandlerLevel(h, fasthttp.CompressBestSpeed)
	}

	// Generate the Address Info
	HostAddress = fmt.Sprintf("%s:%s", serverip, serverport)

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

func requestDetailer(ctx *fasthttp.RequestCtx) {

	// Acquire the Buffer from Pool
	b := fasthttp.AcquireByteBuffer()

	b.B = append(b.B, "\nByteBuffered Magic\n"...)

	b.B = append(b.B, "\nRequest method                = "...)
	b.B = append(b.B, ctx.Method()...)

	b.B = append(b.B, "\nIP Address                    = "...)
	b.B = fasthttp.AppendIPv4(b.B, ctx.RemoteIP())

	b.B = append(b.B, "\nRequestURI                    = "...)
	b.B = append(b.B, ctx.RequestURI()...)

	b.B = append(b.B, "\nRequested path                = "...)
	b.B = append(b.B, ctx.Path()...)

	b.B = append(b.B, "\nHost                          = "...)
	b.B = append(b.B, ctx.Host()...)

	b.B = append(b.B, "\nQuery string                  = "...)
	b.B = append(b.B, ctx.QueryArgs().QueryString()...)

	b.B = append(b.B, "\nUser-Agent                    = "...)
	b.B = append(b.B, ctx.UserAgent()...)

	b.B = append(b.B, "\nConnection Establishment Time = "...)
	b.B = append(b.B, goboseji.TimeIST(ctx.ConnTime()).String()...)

	b.B = append(b.B, "\nRequest Starting Time         = "...)
	b.B = append(b.B, goboseji.TimeIST(ctx.Time()).String()...)

	b.B = append(b.B, "\nRequest Count on Your IP      = "...)
	b.B = append(b.B, fmt.Sprintf("%d", ctx.ConnRequestNum())...)

	b.B = append(b.B, "\nRAW Request                   = ...\n"...)
	b.B = append(b.B, "\n------ REQUEST ------\n\n"...)
	b.B = append(b.B, ctx.Request.String()...)
	b.B = append(b.B, "\n------   END   ------\n"...)

	b.B = append(b.B, "\nIs this connection is WEBSOCK = "...)
	b.B = append(b.B, fmt.Sprintf("%v", ctx.Hijacked())...)

	b.B = append(b.B, "\nProtocol                      = "...)
	b.B = append(b.B, ctx.UserValue("protocol").(string)...)

	// RawWrite all Data at once
	ctx.Write(b.B)

	// Safe to Release the Buffer
	fasthttp.ReleaseByteBuffer(b)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-Custom-Header", "Custom-header-value")

	// Set cookies - As it needs to be created and reference preserved
	var c fasthttp.Cookie
	c.SetKey("Custom-cookie-name")
	c.SetValue("Custom-cookie-value_" + fmt.Sprintf("%d", ctx.ConnRequestNum()))
	ctx.Response.Header.SetCookie(&c)
}

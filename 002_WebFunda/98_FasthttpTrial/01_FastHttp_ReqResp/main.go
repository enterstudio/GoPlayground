package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"regexp"
	"time"
)

var (
	//mode     = flag.Bool("devl", false, " Enable Developer Controls ")
	mode     = flag.Bool("devl", true, " Enable Developer Controls ")
	addr     = flag.String("addr", ":8080", "TCP Address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandlerLevel(h, fasthttp.CompressBestSpeed)
	}

	port := *addr

	// For Open Shift Deployment
	if !*mode {
		serverip := os.Getenv("OPENSHIFT_GO_IP")
		if len(serverip) == 0 {
			serverip = "localhost"
		}
		serverport := os.Getenv("OPENSHIFT_GO_PORT")
		if len(serverport) == 0 {
			serverport = "8080"
		}
		bind := fmt.Sprintf("%s:%s", serverip, serverport)
		port = bind
	}

	log.Println("Starting Server At Localhost " + port)
	if err := fasthttp.ListenAndServe(port, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	// Regex for Path Matching
	checkPath, _ := regexp.MatchString("^/check", string(ctx.Path()))

	// Selector Switch
	switch {

	case checkPath:
		requestDetailer(ctx)
		break

	default:
		ctx.Error("Unknown", fasthttp.StatusBadRequest)
		break
	}

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
	b.B = append(b.B, TimeIST(ctx.ConnTime()).String()...)

	b.B = append(b.B, "\nRequest Starting Time         = "...)
	b.B = append(b.B, TimeIST(ctx.Time()).String()...)

	b.B = append(b.B, "\nRequest Count on Your IP      = "...)
	b.B = append(b.B, fmt.Sprintf("%d", ctx.ConnRequestNum())...)

	b.B = append(b.B, "\nRAW Request                   = ...\n"...)
	b.B = append(b.B, "\n------ REQUEST ------\n\n"...)
	b.B = append(b.B, ctx.Request.String()...)
	b.B = append(b.B, "\n------   END   ------\n"...)

	b.B = append(b.B, "\nIs this connection is WEBSOCK = "...)
	b.B = append(b.B, fmt.Sprintf("%v", ctx.Hijacked())...)

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

// TimeIST function converts timezone of given input time to Bharat (India) Standard time
//
// Example:
//  fmt.Println(TimeIST(time.Now()).String())
//
func TimeIST(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return t.In(loc)
}

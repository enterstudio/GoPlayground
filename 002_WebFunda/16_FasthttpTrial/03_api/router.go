package main

import (
	"github.com/valyala/fasthttp"
	"regexp"
	"runtime"
)

func mainRequestHandler(ctx *fasthttp.RequestCtx) {

	// Regex for Path Matching
	rootPath := string(ctx.Path()) == "/"
	checkPath, _ := regexp.MatchString("^/check", string(ctx.Path()))

	// Add config User Values
	if ctx.IsTLS() {
		ctx.SetUserValue("protocol", "https://")
	} else {
		ctx.SetUserValue("protocol", "http://")
	}

	// Selector Switch
	switch {
	case rootPath:
		host := string(ctx.Host())
		protocol := string(ctx.UserValue("protocol").(string))
		ctx.SetContentType("text/html")
		ctx.WriteString("<!doctype html><html><head><title>" + host)
		ctx.WriteString("</title></head><body><p>")
		ctx.WriteString(" Welcome to " + runtime.Version())
		ctx.WriteString("<br /> Visit <a href=\"" + protocol + host +
			"/check\">API check functionality</a>")
		ctx.WriteString("</p></body></html>")
		break
	case checkPath:
		requestDetailer(ctx)
		break

	default:
		ctx.Error("Unknown", fasthttp.StatusBadRequest)
		break
	}

}

package main

import (
	"github.com/boseji/goboseji"
	"github.com/valyala/fasthttp"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func mainRequestHandler(ctx *fasthttp.RequestCtx) {

	// Path to String
	sPath := string(ctx.Path())
	spath := strings.ToLower(sPath)
	// Regex for Path Matching
	rootPath := sPath == "/"
	checkPath, _ := regexp.MatchString("^/check", spath)
	apiPath, _ := regexp.MatchString("^/api", spath)
	formPath, _ := regexp.MatchString("^/frm", spath)

	// Add config User Values
	if ctx.IsTLS() {
		ctx.SetUserValue("protocol", "https://")
	} else {
		ctx.SetUserValue("protocol", "http://")
	}

	// Selector Switch
	switch {

	// Generate the Home page
	case rootPath:
		host := string(ctx.Host())
		protocol := string(ctx.UserValue("protocol").(string))
		ctx.SetContentType("text/html")
		ctx.WriteString("<!doctype html><html><head><title>" + host)
		ctx.WriteString("</title></head><body><p>")
		ctx.WriteString(" Welcome to " + runtime.Version())
		ctx.WriteString("<br /> Visit <a href=\"" + protocol + host +
			"/check\">Check functionality</a>")
		ctx.WriteString("<br /> Visit <a href=\"" + protocol + host +
			"/api\">Test API functionality</a>")
		ctx.WriteString("<br /> Visit <a href=\"" + protocol + host +
			"/frm\">Test Form functionality</a>")
		ctx.WriteString("</p></body></html>")
		break

	// Move to Check Path
	case checkPath:
		checkEpController(ctx)
		break

	// Actual API Layer
	case apiPath:
		apiEpController(ctx)
		break

	// Trial Form for posting Comments
	case formPath:
		frmEpController(ctx)
		break

	default:
		ctx.Error("Unknown", fasthttp.StatusBadRequest)
		break
	}

	// Add Special Time to Response Header - Tail Chained
	ctx.Response.Header.Add("Time",
		goboseji.TimeIST(time.Now()).Format("January Monday 02-01-2006 15:04:05.000 MST"))
}

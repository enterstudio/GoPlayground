package main

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

func apiEpController(ctx *fasthttp.RequestCtx) {

	// Search if we have Slashes (Rune is 2 bytes) - Boolean
	spaths := bytes.Split(ctx.Path()[1:], []byte("/"))

	ctx.WriteString("API Hit ")

	// Print Other Parameters
	if len(spaths) > 1 {
		ctx.WriteString("\nHas multiple slashes :\n")
		spaths = spaths[1:]
		// Skip the Root Slash
		for i := range spaths {
			ctx.WriteString(string(spaths[i]) + "\n")
		}
	}

	if ctx.QueryArgs().Len() > 0 {
		ctx.WriteString("\nHas Query Parameters :")
		ctx.WriteString("\n" + ctx.QueryArgs().String() + "\n")
	}
}

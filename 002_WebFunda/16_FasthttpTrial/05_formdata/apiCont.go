package main

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

const (
	authKeyPath = "sh3hd9sk2ls0"
)

func apiEpController(ctx *fasthttp.RequestCtx) {

	// Search if we have Slashes (Rune is 2 bytes) - Boolean
	spaths := bytes.Split(ctx.Path()[1:], []byte("/"))

	// Default Response Start
	ctx.WriteString("API Hit")

	// Print Other Parameters
	if len(spaths) > 1 {
		// Special Post Request Processing on Auth
		if ctx.IsPost() && string(spaths[1]) == authKeyPath {
			w := "\nAuth Received\n" + string(ctx.PostBody()) + "\n"
			ctx.WriteString(w)
			ctx.Logger().Printf("%s", w)
		} else {
			ctx.WriteString("\nHas Slashes :\n")
			// Skip the Root Slash
			for i := 1; i < len(spaths); i++ {
				ctx.WriteString(string(spaths[i]) + "\n")
			}
		}
		ctx.WriteString("--\n")
	}

	if ctx.QueryArgs().Len() > 0 {
		ctx.WriteString("\nHas Query Parameters :")
		ctx.WriteString("\n" + ctx.QueryArgs().String() + "\n")
	}

}

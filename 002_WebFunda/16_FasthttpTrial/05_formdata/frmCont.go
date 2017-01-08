package main

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

func frmEpController(ctx *fasthttp.RequestCtx) {

	// Search if we have Slashes (Rune is 2 bytes) - Boolean
	//spaths := bytes.Split(ctx.Path()[1:], []byte("/"))

	switch {
	case bytes.Compare(ctx.Method(), []byte("GET")) == 0:
		ctx.WriteString("GET Path")
		break
	case bytes.Compare(ctx.Method(), []byte("POST")) == 0:
		ctx.WriteString("POST Path")
		break
	default:
		ctx.Error("Unkon Type"+string(ctx.Method()), fasthttp.StatusBadRequest)
		break
	}

}

package main

import (
	"fmt"
	"github.com/boseji/goboseji"
	"github.com/valyala/fasthttp"
)

func checkEpController(ctx *fasthttp.RequestCtx) {

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

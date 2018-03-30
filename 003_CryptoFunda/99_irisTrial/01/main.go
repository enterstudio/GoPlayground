package main

import (
	"gopkg.in/kataras/iris.v4"
	"time"
)

func main() {
	api := iris.New()
	api.Get("/", func(ctx *iris.Context) {
		ctx.Write("Aum at " + time.Now().String())
	})
	api.Listen(":8080")
}

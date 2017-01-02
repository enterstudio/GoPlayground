package main

import (
	"gopkg.in/kataras/iris.v4"
	"time"
)

func main() {
	api := iris.New()
	api.Get("/", func(ctx *iris.Context) {
		ctx.Write("Aum at " + timeNow())
	})
	api.Listen(":8080")
}

func timeNow() string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).String()
}

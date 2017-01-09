package main

import (
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/valyala/fasthttp"
	"net/url"
	"strings"
)

const (
	cookieKeyName = "message"
	messageSep    = " "
	infoSep       = ":"
)

func prepForm(ctx *fasthttp.RequestCtx, ckarray []string) {
	ctx.SetContentType("text/html")
	ctx.WriteString(`
		<h1>GET Path</h1>
		<form method="POST">
		<h3>Message:</h3>
	  <input type="text" name="mesg" placeholder="your Message">
	  <input type="submit">
		</form><br />
		<ul>
		`)
	if len(ckarray) > 0 {
		for i := range ckarray {
			sp := strings.Split(ckarray[i], infoSep)
			w := fmt.Sprintf("<li>%s</li>", sp[1])
			ctx.WriteString(w)
		}
	}
	ctx.WriteString("</ul>")
}

func frmEpController(ctx *fasthttp.RequestCtx) {

	// Get Cookie Information
	lastCookie := string(ctx.Request.Header.Cookie(cookieKeyName))

	switch {

	// GET Request
	case bytes.Compare(ctx.Method(), []byte("GET")) == 0:

		// Create the Post Array
		if len(lastCookie) > 0 {
			// Get all the Messages from the Last cookie
			arr := strings.Fields(lastCookie)
			// Render form with Processed Messages
			prepForm(ctx, arr)
		} else {
			// Blank form
			prepForm(ctx, []string{})
		}

		// Log the Request
		ctx.Logger().Printf("Cookie = " + lastCookie)
		break

	// POST Request
	case bytes.Compare(ctx.Method(), []byte("POST")) == 0:

		// Get the Posted Message field
		messageValue := string(ctx.PostArgs().Peek("mesg"))

		// Check if we really have Some Message
		if len(messageValue) > 0 {

			// Create a New UUID
			s := uuid.NewV4().String()

			// URI Compatible values
			messageValue = url.QueryEscape(messageValue)

			// Attach it to Message to Create the New Cookie Addendum
			newValue := strings.Join([]string{s, messageValue}, infoSep)
			// Add the Values with the Older Cookie Value
			newValue = strings.Join([]string{lastCookie, newValue}, messageSep)

			// Fresh Cookie to be Set
			var newCookie fasthttp.Cookie
			newCookie.SetKey(cookieKeyName)
			newCookie.SetValue(newValue)
			newCookie.SetHTTPOnly(true)
			ctx.Response.Header.SetCookie(&newCookie)

		}

		// Do a Hard redirect after processing the request
		ctx.Redirect(string(ctx.RequestURI()), fasthttp.StatusOK)

		// Log the Request
		ctx.Logger().Printf("mesg = %s", messageValue)

		break

	// Unknown Request
	default:
		ctx.Logger().Printf("Unknown" + string(ctx.Method()))
		ctx.Error("Unkon Type"+string(ctx.Method()), fasthttp.StatusBadRequest)
		break
	}

}

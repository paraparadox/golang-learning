package main

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	users := []string{"rvasiliy"}
	body, _ := json.Marshal(users)

	ctx.SetBody(body)
}

func GetUser(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "you try to see user %s\n", ctx.UserValue("id"))
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/users/:id", GetUser)

	fmt.Println("starting server at :4009")
	log.Fatal(fasthttp.ListenAndServe(":4009", router.Handler))
}

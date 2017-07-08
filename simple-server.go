// Credit: Luke Kysow https://github.com/lkysow
package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Now listening on http://localhost:8080 ...")
	fasthttp.ListenAndServe(":8080", fastHandler)
}

func fastHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "OK")
}

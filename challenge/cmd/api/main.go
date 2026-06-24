package main

import (
	"fmt"

	"meshnode/internal/app"
)

func main() {
	rt := app.Load()
	fmt.Println(rt.Bind(), rt.StoreName(), rt.RouteCount())
}

package main

import (
	authz "github.com/carlossantin/resource-policy-framework/auth"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	r := authz.Role{}
	println(r.ID)

	e.Start(":8080")
	println("Server started on port 8080")
}

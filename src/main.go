package main

import (
	"fmt"

	"github.com/go-http-test/handler"
	"github.com/go-http-test/injector"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

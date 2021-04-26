package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/masamichhhhi/go-clean/handler"
	"github.com/masamichhhhi/go-clean/injector"
)

func main() {
	fmt.Println("sever start")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":9000"))
}

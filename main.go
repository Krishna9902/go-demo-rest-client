package main

import (
	"GoBegin/controllers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Application")
	controllers.RegisterControllers()
	fmt.Println("Started  Application on 8030 port")
	http.ListenAndServe(":8030", nil)
}

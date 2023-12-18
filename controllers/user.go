package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	if request.Method == "GET" {
		fmt.Println("This is GET Method")
		writer.Write([]byte("This is GET Hello World!!!"))
	}

	if request.Method == "POST" {
		fmt.Println("This is POST Method")
		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("server: could not read request body: %s\n", err)
		}
		fmt.Printf("server: request body: %s\n", reqBody)
		writer.Write(reqBody)
	}

}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

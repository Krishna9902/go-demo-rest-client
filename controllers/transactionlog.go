package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type transactionController struct {
	userIDPattern *regexp.Regexp
}

func (uc transactionController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me

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

func newTransactionController() *transactionController {
	return &transactionController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

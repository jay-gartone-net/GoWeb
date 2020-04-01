package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/models/controllers"
)

func main() {

	port := 3000
	retries := 2
	_, err := startWebServer(port, retries)
	fmt.Println(port, err)

}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	fmt.Println("Number of retires:", numberOfRetries)
	fmt.Println("Server started on port", port)
	return port, nil
}

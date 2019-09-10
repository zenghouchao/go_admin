package main

import (
	"fmt"
	_ "github.com/go_admin/router"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("The HTTP server failed to start:\n", err.Error())
	}
}

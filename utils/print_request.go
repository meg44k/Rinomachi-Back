package utils

import (
	"fmt"
	"net/http"
)

func PrintRequest(r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL.String())
}
	
// Serve current dir with HTTP
package main

import (
	"net/http"
)

func main() {
	panic(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))
}

package main
import (
	"net/http"
	"fmt"
	"log"
	"os"
)
func main() {

	host := os.Getenv("HOSTNAME")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, World, from: %v", host)
	})
	log.Fatal(http.ListenAndServe(":3005", nil))
}
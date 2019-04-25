package main

import (
	"fmt"
	"log"
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/rts"
	"github.com/gorilla/handlers"
)

// Entry point of the program
func main() {
	r := rts.RTS()
	fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.")
	fmt.Println("-.-.-.-.-.-.                            -.-.-.-.-.")
	fmt.Println("-.-.-.-.-.-.    Listen on port :4000    -.-.-.-.-.")
	fmt.Println("-.-.-.-.-.-.                            -.-.-.-.-.")
	fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.")
	log.Fatal(http.ListenAndServe(":4000", handlers.CORS()(r)))
}

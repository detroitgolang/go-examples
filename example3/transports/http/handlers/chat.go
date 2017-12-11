package handlers

import (
	"log"
	"net/http"

	"github.com/mccurdyc/go-examples/example3/transports/http/connections"
)

func Chat(chub *connections.ConnectionHub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go chub.WriteMessage()
		log.Println("wrote message")
	})
}

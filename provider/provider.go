package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pact-foundation/pact-go/types"
)

//StartProvider hi yes I be comment
func StartProvider() {
	mux := http.NewServeMux()
	lastName := "billy"

	fmt.Println("Starting server")
	mux.HandleFunc("/foobar", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf(`{"lastName":"%s"}`, lastName))
	})

	mux.HandleFunc("/setup", func(w http.ResponseWriter, req *http.Request) {
		var s *types.ProviderState
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&s)
		if s.State == "User foo exists" {
			lastName = "bar"
		}

		w.Header().Add("Content-Type", "application/json")
	})
	log.Fatal(http.ListenAndServe(":8000", mux))
	fmt.Println("Server Started")

}

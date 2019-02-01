package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"log"

	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/dsl"
)

func startProvider() {
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

func TestProvider(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "ResultsRenderer",
		Provider: "ResultsAPI",
	}

	go startProvider()

	pactDir := "../consumer/pacts"
	fmt.Println("About to start pact")
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:        "http://localhost:8000",
		PactURLs:               []string{filepath.ToSlash(fmt.Sprintf("%s/resultsrenderer-resultsapi.json", pactDir))},
		ProviderStatesSetupURL: "http://localhost:8000/setup",
	})
	if err != nil {
		t.Fatal(err)
	}
}

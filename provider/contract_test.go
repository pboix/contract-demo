package provider

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "ResultsRenderer",
		Provider: "ResultsAPI",
	}

	go StartProvider()

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

package provider

import (
	"fmt"
	"log"
	"net/http"
)

//StartProvider hi yes I be comment
func StartProvider() {
	mux := http.NewServeMux()

	mux.HandleFunc("/forms/abcdef/responses", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		o := `{
			"title": "My Awesome Form",
			"form": [
				{
					"token": "21085286190ffad1248d17c4135ee56f",
					"answers": [
						{
							"field": {
								"id": "hVONkQcnSNRj",
								"type": "short_text",
								"title": "What's your favourite cheese?"
							},
							"type": "text",
							"answer": "Gouda"
						},
						{
							"field": {
								"id": "RUqkXSeXBXSd",
								"type": "yes_no",
								"title": "do you eat cheese everyday?"
							},
							"type": "boolean",
							"answer": true
						}
					]
				}
			]
		}`
		fmt.Fprintf(w, fmt.Sprintf(o))
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}

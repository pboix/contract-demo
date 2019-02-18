package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/forms/abcdef/responses", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf(`{
			"title": "My Awesome Form",
			"form":
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
					"answer": "true"
				  }
				]
			  }	
		  }`))
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}

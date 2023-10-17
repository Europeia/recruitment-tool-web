package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// required query params: nations [comma separated list of nations], template [telegram template -- without %s]
	// example: localhost:8080/?nations=darcness,decacon,upc&template=TEMPLATE-8675309
	nations := r.URL.Query().Get("nations")
	if nations == "" {
		http.Error(w, "Please enter nations", http.StatusBadRequest)
		return
	}

	template := r.URL.Query().Get("template")
	if template == "" {
		http.Error(w, "Please enter template", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://www.nationstates.net/page=compose_telegram?tgto=%s&message=%s", nations, template)
	fmt.Fprintf(w, `<a href="%s"><button style="background-color: red; color: white; font-size: 20px; padding: 15px 32px; text-align: center; display: inline-block; border-radius: 12px; border: none; cursor: pointer;">Send Telegram</button></a>`, url)
}

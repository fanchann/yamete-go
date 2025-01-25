package main

import (
	"encoding/json"
	"log"
	"net/http"

	yametego "github.com/fanchann/yamete-go"
)

// CensoredText represents the response structure for censored text.
type CensoredText struct {
	OriginalText  string   `json:"original_text"`
	CensoredText  string   `json:"censored_text"`
	CensoredCount int      `json:"censored_count"`
	CensoredWords []string `json:"censored_words"`
}

// BodyRequest represents the incoming request body.
type BodyRequest struct {
	Text string `json:"text"`
}

// Controller handles HTTP requests.
type controller struct {
	service service
}

func (c *controller) analyze(w http.ResponseWriter, r *http.Request) {
	var body BodyRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := c.service.analyze(body.Text)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

type service struct {
	yg yametego.Yamete
}

func (s *service) analyze(phrase string) CensoredText {
	result := s.yg.AnalyzeText(phrase)
	return CensoredText{
		OriginalText:  result.OriginalText,
		CensoredText:  result.CensoredText,
		CensoredCount: result.CensoredCount,
		CensoredWords: result.CensoredWords,
	}
}

func main() {
	// Initialize the Yamete instance.
	yameteInstance, err := yametego.NewYamete(
		&yametego.YameteConfig{
			URL: "https://raw.githubusercontent.com/fanchann/toxic-word-list/refs/heads/master/id_toxic_371.txt",
		})
	if err != nil {
		log.Fatalf("Failed to initialize yametego: %v", err)
	}

	// Create the service and controller.
	svc := service{yg: yameteInstance}
	ctrl := controller{service: svc}

	// Set up HTTP routes.
	http.HandleFunc("/analyze", ctrl.analyze)

	// Start the main HTTP server.
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

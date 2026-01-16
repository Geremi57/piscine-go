package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message Message `json:"message"`
}

type ChatResponse struct {
	Choices []Choice `json:"choices"`
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		return
	}

	// Read question
	body, _ := io.ReadAll(r.Body)
	question := string(body)

	// Prepare request for OpenAI
	reqBody, _ := json.Marshal(ChatRequest{
		Model: "gpt-4o-mini", // you can also try "gpt-3.5-turbo"
		Messages: []Message{
			{Role: "user", Content: question},
		},
	})

	req, _ := http.NewRequest("POST", openaiURL, strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")) // ✅ FIXED

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error calling OpenAI: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	// Log raw response for debugging
	raw, _ := io.ReadAll(resp.Body)
	fmt.Println("Raw response:", string(raw))

	var chatResp ChatResponse
	if err := json.Unmarshal(raw, &chatResp); err != nil {
		http.Error(w, "Error parsing OpenAI response: "+err.Error(), 500)
		return
	}

	if len(chatResp.Choices) > 0 {
		fmt.Fprint(w, chatResp.Choices[0].Message.Content)
	} else {
		fmt.Fprint(w, "No response from OpenAI")
	}
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

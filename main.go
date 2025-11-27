package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Request Structure
type TransferRequest struct {
	User   string `json:"user"`
	Amount int    `json:"amount"`
}

// Response Structure
type TransferResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	// 1. Static Files (Frontend)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// 2. API Endpoint
	http.HandleFunc("/transfer", handleTransfer)

	// 3. Health Check (for Kubernetes Liveness Probe)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST method required", http.StatusMethodNotAllowed)
		return
	}

	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Server Log (Infrastructure Monitoring)
	fmt.Printf("[LOG] Transfer Request: User=%s, Amount=%d JPY\n", req.User, req.Amount)

	// Response to Client (Japanese)
	resp := TransferResponse{
		Message: fmt.Sprintf("%d円の送金が完了しました。", req.Amount),
		Status:  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
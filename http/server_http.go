package http

import (
	"encoding/json"
	"fmt"
	"github.com/erik-olsson-op/yakvs/engine"
	"github.com/erik-olsson-op/yakvs/logger"
	"github.com/erik-olsson-op/yakvs/model"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"sync"
)

func Init(wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/execute", executeHandler)
	http.HandleFunc("/health", healthHandler)
	addr := fmt.Sprintf(":%v", viper.Get("HTTP_PORT"))
	logger.Logger.Infof("HTTP server is running on port %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method only GET", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method only POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("ERR: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var query model.Query
	if err := json.Unmarshal(body, &query); err != nil {
		http.Error(w, fmt.Sprintf("ERR: %v", err), http.StatusBadRequest)
		return
	}
	response, err := engine.Parse(&query)
	if err != nil {
		http.Error(w, fmt.Sprintf("ERR: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response.Value)
	if err != nil {
		return
	}
}

package http

import (
	"encoding/json"
	"fmt"
	"github.com/erik-olsson-op/yakvs/engine"
	"github.com/erik-olsson-op/yakvs/logger"
	"github.com/erik-olsson-op/yakvs/model"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"sync"
)

func Init(wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/execute", executeHandler)
	http.HandleFunc("/health", healthHandler)
	addr := viper.GetString("HTTPS_PORT")
	logger.Logger.Infof("HTTPS server is running on port %v", addr)
	cert := viper.GetString("SERVER_CERT")
	key := viper.GetString("SERVER_KEY")
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", addr), cert, key, nil)
	if err != nil {
		logger.Logger.Fatal(err)
	}
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
		http.Error(w, fmt.Sprintf("ERR: %v", err), http.StatusBadRequest)
		return
	}
}

package main

import (
	"github.com/erik-olsson-op/yakvs/grpc"
	"github.com/erik-olsson-op/yakvs/http"
	"github.com/spf13/viper"
	"log"
	"sync"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error loading .env file, falling back to environment variables")
	}
	viper.AutomaticEnv()
	var wg sync.WaitGroup
	wg.Add(2)
	go grpc.Init(&wg)
	go http.Init(&wg)
	wg.Wait()
}

package main

import (
	"fmt"
	"os"

	"go-land/pkg/config"
	lru "go-land/pkg/lrudemo"
)

func main() {
	fmt.Println("Hello, Go's play land!")

	demoStructAlias()
	demoLruCache()
}

func demoStructAlias() {
	fmt.Println("\n* Demonstrates using type alias to handle special fields during JSON marshaling and unmarshaling.")
	data, err := os.ReadFile("./pkg/config/kafka.json")

	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var kafkaConfig config.KafkaConfig
	if err := kafkaConfig.UnmarshalJSON(data); err != nil {
		fmt.Println("Error parsing config:", err)
		return
	}

	fmt.Printf("Id: %s\n", kafkaConfig.Id)
	fmt.Printf("Interval: %s\n", kafkaConfig.Interval)
}

func demoLruCache() {
	fmt.Println("\n* Demonstrates LRU cache.")
	lru.Main()
}
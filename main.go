package main

import (
	"fmt"
	"os"

	"github.com/andrxu/go-land/pkg/config"
)

func main() {
	fmt.Println("Hello, Go's play land!")

	demoStructAlias()
}

func demoStructAlias() {
	fmt.Println("* Demonstrates using type alias to handle special fields during JSON marshaling and unmarshaling.")
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

package main

import (
	"flag"
	"log"
	"strings"
	internal "wipcocli/internal/app"
)

func main() {
	var apiKey, todo string
	var isCompleted bool

	flag.StringVar(&apiKey, "apikey", "", "Set the API key")
	flag.StringVar(&todo, "todo", "", "Set the todo message")
	flag.Parse()

	apiKey = handleAPIKey(apiKey)

	client := internal.NewClient(apiKey)

	if todo == "" {
		args := flag.Args()
		if len(args) == 0 {
			log.Fatal("⚠️ No message provided")
		}

		todo = strings.Join(args, " ")
		isCompleted = true
	}

	response := internal.PostTodo(client, todo, isCompleted)
	log.Println("✅ Posted: " + response.Data.CreateTodo.URL)
}

func handleAPIKey(apiKey string) string {
	if apiKey != "" {
		internal.SaveAPIKey(apiKey)
		log.Fatal("🔑API key has been set.")
	} else {
		apiKey = internal.LoadAPIKey()
		if apiKey == "" {
			log.Fatal("🔑API key not set. Use 'wip apikey YOUR_API_KEY' to set it.")
		}
	}

	return apiKey
}

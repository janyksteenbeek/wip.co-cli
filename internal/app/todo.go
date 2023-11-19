package internal

import (
	"encoding/json"
	"log"
	"time"
)

type TodoInput struct {
	Body        string `json:"body"`
	CompletedAt string `json:"completed_at,omitempty"`
}

type CreateTodoMutation struct {
	Mutation string    `json:"mutation"`
	Input    TodoInput `json:"input"`
}

type GraphqlQuery struct {
	Query     string    `json:"query"`
	Variables TodoInput `json:"variables"`
}

type CreateTodoResponse struct {
	Data struct {
		CreateTodo struct {
			URL string `json:"url"`
		}
	}
}

func PostTodo(client *Client, todo string, completed bool) CreateTodoResponse {
	var completedAt string

	if completed {
		completedAt = time.Now().Format(time.RFC3339)
	}

	mutation := `
  		mutation CreateTodoMutation($body: String!, $completed_at: DateTime) {
			createTodo(input: {body: $body, completed_at: $completed_at}) {
				url
			}
		}
		`

	query := GraphqlQuery{
		Query: mutation,
		Variables: TodoInput{
			Body:        todo,
			CompletedAt: completedAt,
		},
	}

	queryJSON, err := json.Marshal(query)
	response, err := client.Request(queryJSON)
	if err != nil {
		log.Fatal(err)
	}

	// json decode
	todoResponse := CreateTodoResponse{}
	err = json.Unmarshal(response, &todoResponse)
	if err != nil {
		log.Fatal(err)
	}

	return todoResponse
}

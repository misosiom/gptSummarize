package summarize

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func SummarizeText(content string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", ErrMissingAPIKey
	}

	payload := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "user", "content": "Please summarize the following text and list the important points in bullet form. Then, translate the summary into Japanese.:\n\n" + content},
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", ErrNoChoices
	}

	return result.Choices[0].Message.Content, nil
}

var (
	ErrMissingAPIKey = &SummarizeError{"OPENAI_API_KEY が設定されていません"}
	ErrNoChoices     = &SummarizeError{"GPTの応答に 'choices' がありません"}
)

type SummarizeError struct {
	Message string
}

func (e *SummarizeError) Error() string {
	return e.Message
}

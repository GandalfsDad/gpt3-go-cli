package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	key := os.Getenv("OPEN_API_KEY")
	c := gogpt.NewClient(key)
	ctx := context.Background()

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Input Max Tokens")
	maxtokens, _ := inputReader.ReadString('\n')
	maxTokensInt, _ := strconv.Atoi(maxtokens)

	fmt.Println("Input Temperature")
	temp, _ := inputReader.ReadString('\n')
	tempFloat, _ := strconv.ParseFloat(temp, 32)

	fmt.Println("Input Prompt:")
	input, _ := inputReader.ReadString('\n')

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-002",
		MaxTokens:   maxTokensInt,
		Prompt:      input,
		Temperature: float32(tempFloat),
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}

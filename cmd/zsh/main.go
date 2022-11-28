package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	key := os.Getenv("OPEN_API_KEY")
	c := gogpt.NewClient(key)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-002",
		MaxTokens:   20,
		Prompt:      string(loadHistory()),
		Temperature: 0.7,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Choices[0].Text)

}

func loadHistory() []byte {

	dirname, _ := os.UserHomeDir()
	path := filepath.Join(dirname, ".zsh_history")

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 2000)
	stat, _ := os.Stat(path)
	start := stat.Size() - 2000
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("%s\n", buf[:10])
	}

	re := regexp.MustCompile(`: ([\d]+) *:0;`)
	commands := re.ReplaceAll(buf, []byte(""))

	ga := regexp.MustCompile(`ga `)
	commands = ga.ReplaceAll(commands, []byte("git add"))

	gc := regexp.MustCompile(`gc `)
	commands = gc.ReplaceAll(commands, []byte("git commit"))

	gp := regexp.MustCompile(`gp\n`)
	commands = gp.ReplaceAll(commands, []byte("git push\n"))

	gst := regexp.MustCompile(`gst\n`)
	commands = gst.ReplaceAll(commands, []byte("git status\n"))

	return commands
}

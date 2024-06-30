package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type FileChange struct {
	Filename string `json:"filename"`
	Patch    string `json:"patch"`
}

var (
	prNumber  string
	repoOwner string
	repoName  string
)

var rootCmd = &cobra.Command{
	Use:   "github_pr_diff",
	Short: "A tool to fetch and display GitHub PR diffs",
	Long:  `This tool fetches the diff for a specified GitHub Pull Request and displays it in a markdown-friendly format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchAndDisplayDiff()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&prNumber, "pr", "p", "", "Pull Request number (required)")
	rootCmd.Flags().StringVarP(&repoOwner, "owner", "o", "", "Repository owner (required)")
	rootCmd.Flags().StringVarP(&repoName, "repo", "r", "", "Repository name (required)")
	rootCmd.MarkFlagRequired("pr")
	rootCmd.MarkFlagRequired("owner")
	rootCmd.MarkFlagRequired("repo")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fetchAndDisplayDiff() {
	// auth token is needed.
	cmd := exec.Command("gh", "auth", "token")
	tokenBytes, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error getting GitHub token: %v", err)
	}
	token := string(tokenBytes[:len(tokenBytes)-1]) // 改行を削除

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%s/files", repoOwner, repoName, prNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var files []FileChange
	err = json.Unmarshal(body, &files)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	for _, file := range files {
		fmt.Printf("```diff\n")
		fmt.Printf("# %s\n", file.Filename)
		fmt.Printf("%s\n", file.Patch)
		fmt.Printf("```\n\n")
	}
}

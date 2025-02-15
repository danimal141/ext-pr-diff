package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type FileChange struct {
	Filename string `json:"filename"`
	Patch    string `json:"patch"`
}

var prLink string

var rootCmd = &cobra.Command{
	Use:   "ext_pr_diff",
	Short: "A tool to fetch and display GitHub PR diffs",
	Long:  `This tool fetches the diff for a specified GitHub Pull Request and displays it in a markdown-friendly format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchAndDisplayDiff()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&prLink, "link", "l", "", "GitHub Pull Request link (required)")
	rootCmd.MarkFlagRequired("link")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fetchAndDisplayDiff() {
	repoOwner, repoName, prNumber := parsePRLink(prLink)

	cmd := exec.Command("gh", "auth", "token")
	tokenBytes, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error getting GitHub token: %v", err)
	}
	token := strings.TrimSpace(string(tokenBytes))

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

func parsePRLink(link string) (string, string, string) {
	re := regexp.MustCompile(`github\.com/([^/]+)/([^/]+)/pull/(\d+)`)
	matches := re.FindStringSubmatch(link)
	if len(matches) != 4 {
		log.Fatalf("Invalid GitHub PR link: %s", link)
	}
	return matches[1], matches[2], matches[3]
}

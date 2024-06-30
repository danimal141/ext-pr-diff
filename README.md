# ext-pr-diff
This tool is designed to be used for getting AI-generated explanations of code changes. It fetches GitHub Pull Request (PR) diffs and displays them in a markdown format, making it easy to input the changes into a generative AI system for analysis and explanation.

ext-pr-diff is a command-line tool to easily fetch and display GitHub Pull Request (PR) diffs in markdown format.

## Features

- Retrieve the list of changed files in a GitHub PR
- Display the diff for each file in markdown format
- Authenticate using GitHub CLI

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [GitHub CLI](https://cli.github.com/)

## Installation

1. Clone the repository:
h`git clone https://github.com/danimal141/ext-pr-diff.git`

2. Navigate to the project directory:

`cd ext-pr-diff`

3. Build the tool:

`go build -o bin/ext-pr-diff`

## Usage

1. Authenticate with GitHub CLI (if you haven't already):

`gh auth login`

2. Run the tool:

`./ext-pr-diff --pr <PR_NUMBER> --owner <REPO_OWNER> --repo <REPO_NAME>`

Or use short flags:

`./ext-pr-diff -p <PR_NUMBER> -o <REPO_OWNER> -r <REPO_NAME>`

Example:

`./ext-pr-diff -p 1 -o danimal141 -r ext-pr-diff`

3. The diff for each file in the PR will be displayed in markdown format.

## Options

- `--pr`, `-p`: Pull request number (required)
- `--owner`, `-o`: Repository owner (required)
- `--repo`, `-r`: Repository name (required)
- `--help`, `-h`: Display help message

## Contributing

Contributions to ext-pr-diff are welcome! Here's how you can contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

Please ensure your code adheres to the existing style to maintain consistency.

## Bug Reports and Feature Requests

If you encounter any bugs or have ideas for new features, please [open an issue](https://github.com/danimal141/ext-pr-diff/issues) on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
- [GitHub CLI](https://cli.github.com/) - GitHub's official command line tool

# github_cli
Scalabit Task


github_cli is a command-line application that allows
	you to fetch various data related to a target user from GitHub.
	It provides convenient and efficient ways to retrieve information
	such as pull requests, repositories, followers, and more.

Usage:
  github_cli [command] [target user]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  following   Get repos followed by the target user
  help        Help about any command
  prs         Get a list of pull requests for the target user
  repos       Retrieve information about repositories owned by the target user

Flags:
      --config string   config file (default is $HOME/.github_cli.yaml)
  -h, --help            help for github_cli
  -t, --toggle          Help message for toggle

Use "github_cli [command] --help" for more information about a command.

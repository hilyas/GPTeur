# gpteur
A GPT CLI tools

## Installation

```bash
git clone github.com/hilyas/gpteur
go build -o bin/gpteur main.go 
```

## Usage

```bash
Usage:
  gpteur [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate a response from the ChatGPT model
  help        Help about any command

Flags:
      --apikey string   API key for ChatGPT authentication
  -h, --help            help for gpteur
  -t, --toggle          Help message for toggle

Use "gpteur [command] --help" for more information about a command.
```

## License

See [LICENSE](LICENSE).

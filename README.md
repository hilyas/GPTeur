# GPTeur

A GPT CLI tool to interact with the [ChatGPT](https://chatgpt.com/) API.

## Installation

```bash
git clone github.com/hilyas/gpteur
go build -o bin/gpteur main.go 
```

You will need your API key from [ChatGPT](https://chatgpt.com/). You can get it from the [dashboard](https://chatgpt.com/dashboard).

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

Use "gpteur [command] --help" for more information about a command.
```

Example:

```bash
./bin/gpteur generate --prompt "Write a simple for loop in python'" --max-tokens 150 --temperature 0.9 --apikey XXXXXXXXXXXXXXXXXXXXXXXXX
```

## Model

Currently the default model is `gpt-3.5-turbo`. Work is scheduled to support more models in the future.

## Roadmap

- [ ] Support more models
- [ ] Support more API endpoints
- [ ] Support more CLI flags
- [ ] Support more CLI commands
- [ ] Tests
- [ ] UX improvements

## License

See [LICENSE](LICENSE).

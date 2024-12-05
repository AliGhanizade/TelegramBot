# Telegram Bot in Go

This repository contains a simple Telegram bot implemented in the Go programming language. The bot is designed to respond to messages sent to the @Alex_Simple_bot Telegram account.

## Features

- Responds to user messages
- Simple and easy to understand code structure
- Built with Go for performance and efficiency

## Getting Started

### Prerequisites

- Go (version 1.15 or higher) installed on your machine

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AliGhanizade/TelegramBot.git
   ```
   
2. Navigate to the project directory:
   ```bash
   cd telegram_bot
   ```

3. Install the necessary dependencies:
   ```bash
   go mod tidy
   ```

### Configuration

Before running the bot, you need to set your Telegram Bot Token. You can obtain a token by creating a new bot using [BotFather](https://core.telegram.org/bots#botfather).

Add your token to the environment variable:
```bash
TELEGRAM_BOT_TOKEN="your_bot_token_here"
```

### Running the Bot

To start the bot, run the following command:
```bash
go run firstbot.go
```

## License

This project does not have a specified license. Please check the repository for more details.

## Acknowledgements

- [Telegram Bot API](https://core.telegram.org/bots/api) for providing the API documentation.

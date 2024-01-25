package main

import (
	"github.com/vikpe/qwza-discord-bot/internal/qwza"
	"github.com/vikpe/qwza-discord-bot/internal/qwza/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	cfg, err := config.FromFile("config.json")

	if err != nil {
		log.Fatal("unable to load config", err)
		return
	}

	bot, err := qwza.New(
		os.Getenv("BOT_TOKEN"),
		os.Getenv("GUILD_ID"),
		cfg,
	)

	if err != nil {
		log.Fatal("unable to create qwzaBot", err)
		return
	}

	bot.Start() // blocking operation
}

package main

import (
	"dgb/meter.readings.bot/internal/application"
	"dgb/meter.readings.bot/internal/configuration"
	"dgb/meter.readings.bot/internal/database"
)

func main() {
	conf := configuration.NewConfig()

	repo := database.NewRepository(conf)
	application.HandleRequests(conf)
	application.StartTelegramBot(conf, *repo)
}

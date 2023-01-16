package configuration

import (
	"os"
	"strconv"
	"strings"
)

type Configuration struct {
	TELEGRAM_BOT_TOKEN   string
	TELEGRAM_VALID_USERS []int64
	MONGO_CONNECTION     string
	MONGO_DB             string
	MONGO_COLLECTION     string
	HTTP_PORT            string
}

func NewConfig() Configuration {

	configuration := &Configuration{}
	configuration.TELEGRAM_BOT_TOKEN = os.Getenv("METER_READINGS_TELEGRAM_BOT_TOKEN")

	validUsers := strings.Split(os.Getenv("METER_READINGS_TELEGRAM_VALID_USERS"), ",")
	for _, s := range validUsers {
		i, _ := strconv.ParseInt(s, 10, 64)
		configuration.TELEGRAM_VALID_USERS = append(configuration.TELEGRAM_VALID_USERS, i)
	}

	configuration.MONGO_CONNECTION = os.Getenv("METER_READINGS_MONGO_CONNECTION")
	configuration.MONGO_COLLECTION = os.Getenv("METER_READINGS_MONGO_COLLECTION")
	configuration.MONGO_DB = os.Getenv("METER_READINGS_MONGO_DB")
	configuration.HTTP_PORT = os.Getenv("METER_READINGS_HTTP_PORT")

	return *configuration
}

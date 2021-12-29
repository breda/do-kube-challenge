package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/jaswdr/faker"
)

type Log struct {
	Level   string
	Message string
}

func RandomLogLevel() string {
	levels := []string{"debug", "info", "notice", "warning", "error", "critical", "alert", "emergency"}
	random := rand.Intn(len(levels))

	return levels[random]
}

func main() {
	faker := faker.New()

	fluentPort, fluentHost := os.Getenv("FLUENTD_PORT"), os.Getenv("FLUENTD_HOST")

	if fluentPort == "" || fluentHost == "" {
		fmt.Println("Missing Fluentd port or host in environment variables.")
		os.Exit(1)
	}

	fluentPortInt, err := strconv.Atoi(fluentPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger, err := fluent.New(fluent.Config{
		FluentPort: fluentPortInt,
		FluentHost: fluentHost,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logger.Close()

	for {
		log := Log{
			Level:   RandomLogLevel(),
			Message: faker.Lorem().Sentence(rand.Intn(25)),
		}

		err := logger.Post("app.logs", log)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}

}

package main

import (
	"github.com/codeuniversity/smag-mvp/kafka"
	inserter "github.com/codeuniversity/smag-mvp/postgres-inserter"
	"github.com/codeuniversity/smag-mvp/service"
	"github.com/codeuniversity/smag-mvp/utils"
)

func main() {
	groupID := utils.MustGetStringFromEnv("KAFKA_GROUPID")
	rTopic := utils.MustGetStringFromEnv("KAFKA_RTOPIC")
	wTopic := utils.MustGetStringFromEnv("KAFKA_WTOPIC")
	isUserDiscovery := utils.GetBoolFromEnvWithDefault("USER_DISCOVERY", false)

	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "127.0.0.1:9092")
	postgresHost := utils.GetStringFromEnvWithDefault("POSTGRES_HOST", "127.0.0.1")
	postgresPassword := utils.GetStringFromEnvWithDefault("POSTGRES_PASSWORD", "")

	qReaderConfig := kafka.NewReaderConfig(kafkaAddress, groupID, topic)
	qWriterConfig := kafka.NewWriterConfig(kafkaAddress, topic, true)

	i := inserter.New(
		kafkaAddress,
		postgresHost,
		postgresPassword,
		kafka.NewReader(qReaderConfig),
		kafka.NewWriter(qWriterConfig),
	)

	service.CloseOnSignal(i)
	go i.Run()

	i.WaitUntilClosed()
}

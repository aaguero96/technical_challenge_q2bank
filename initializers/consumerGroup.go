package initializers

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateConsumerGroup() {
	streamName := os.Getenv("STREAM_REDIS_NAME")
	consumerGroup := os.Getenv("CONSUMER_GROUP_REDIS_NAME")
	if _, err := Client.XGroupCreateMkStream(streamName, consumerGroup, "0").Result(); err != nil {
		if !strings.Contains(fmt.Sprint(err), "BUSYGROUP") {
			log.Fatalf("Error on create Consumer Group: %v ...\n", consumerGroup)
		}
	}
}

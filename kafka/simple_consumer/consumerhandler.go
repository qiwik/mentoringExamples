package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/cockroachdb/errors"
)

type ConsumerHandler struct{}

type Msg struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *ConsumerHandler) ConsumeClaim(gs sarama.ConsumerGroupSession, gc sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-gc.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}

			var m Msg

			if err := json.Unmarshal(message.Value, &m); err != nil {
				return errors.Wrap(err, "unmarshal kafka message")
			}

			fmt.Println(fmt.Sprintf("The message is: %v. Partition: %d", m, message.Partition))

			gs.MarkMessage(message, "")
		case <-gs.Context().Done():
			return nil
		}
	}
}

func (c *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

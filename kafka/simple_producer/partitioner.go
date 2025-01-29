package main

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/cockroachdb/errors"
)

type Partition struct{}

func NewPartition(topic string) sarama.Partitioner {
	return &Partition{}
}

func (p *Partition) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	var m Msg

	b, err := message.Value.Encode()
	if err != nil {
		return 0, errors.Wrap(err, "encode message")
	}

	if err = json.Unmarshal(b, &m); err != nil {
		return 0, errors.Wrap(err, "unmarshal kafka message of inner topic")
	}

	if m.ID%2 == 0 {
		return 0, nil
	}

	return 1, nil
}

func (p *Partition) RequiresConsistency() bool {
	return false
}

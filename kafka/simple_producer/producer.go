package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/cockroachdb/errors"
)

var topicName = "study_topic"

type Producer struct {
	kafkaClient sarama.SyncProducer
}

type Msg struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewProducer(client sarama.SyncProducer) *Producer {
	return &Producer{kafkaClient: client}
}

func (p *Producer) Send(msg Msg) error {
	encoded, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "marshal message to json")
	}

	m := &sarama.ProducerMessage{
		Topic: topicName,
		Key:   sarama.StringEncoder(rune(msg.ID)), // Ключ, благодаря которому сообщение будет отправляться в разные партиции.
		// На его основе Кафка генерирует хеш, который из себя будет в конечном итоге представлять число,
		// по которому принимается решение в какую ПАРТИЦИЮ отправлять сообщение.
		// Таким образом достигается параллелизм обработки данных. Ну а если ключа нет, то сообщения по партициям
		// раскидываются по round-robinу.
		Value:   sarama.ByteEncoder(encoded), // Сами данные, отправляющиеся в топик.
		Headers: nil,                         // Заголовки схожи с тем, что есть в http. В виде key-value они позволяют консьюмеру
		// понять, что же ему можно сделать с сообщением, кому отправить далее, куда сохранить, кто был отправителем. То есть
		// заголовки нужны для распределения сообщений, для версионирования, идентификации клиента, уровень приоритизации сообщения,
		// для аналитики в ui и так далее. Тут надо быть осторожным, потому что заголовки все же добавляют веса сообщению.
		Metadata: nil, // В метадату сваливаются все остальные произвольные данные, которые ни на что не влияют, но их нахождение
		// решает какую-то проблему этого сообщения. Сарама просто прокидывает их до Кафки и ничего больше не делает.
		Offset:    0,           // Продюсер сам проставляет сюда позицию, на которую встало сообщение, если получено подтверждение от Кафки.
		Partition: 0,           // Продюсер сам проставит сюда партицию, в которую попало сообщение, если получено подтверждение от Кафки.
		Timestamp: time.Time{}, // Продюсер проставит сюда время добавления сообщения в топик. Могут быть разные стратегии.
	}

	// Отправка сообщения. Дополнительно получаем данные о том, в какую партицию упало сообщение и какую позицию в топике имеет.
	partition, offset, err := p.kafkaClient.SendMessage(m)
	if err != nil {
		return errors.Wrap(err, "send message to kafka")
	}

	fmt.Println(fmt.Sprintf("To partition %d on offset %d", partition, offset))

	return nil
}

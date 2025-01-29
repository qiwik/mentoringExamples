package main

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
)

var topicName = "study_topic"

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0  // Минимально поддерживаемая версия кафки.
	config.ClientID = "test-consumer" // Идентификатор сервиса.

	config.Net.MaxOpenRequests = 10            // Сколько запросов выдерживает подключение.
	config.Net.DialTimeout = 30 * time.Second  // Таймаут на подключение.
	config.Net.ReadTimeout = 30 * time.Second  // Таймаут на ожидание ответа от брокера.
	config.Net.WriteTimeout = 30 * time.Second // Таймаут на ожидание передачи данных.

	// config.Consumer.Group - настройки управления консьюмер группой.
	config.Consumer.Retry.Backoff = 1 * time.Second // Интервал между попытками вычитать сообщение.
	// config.Consumer.Fetch - настройка объема байтов, вычитываемого из топика.
	config.Consumer.MaxWaitTime = 500 * time.Millisecond // Максимальное время ожидание брокером минимального
	// количества байт, который установил консьюмер. Если не набирается, то отдается столько, сколько есть.
	config.Consumer.MaxProcessingTime = 2 * time.Second // Максимальное время обработки сообщений, которое
	// ожидает консьюмер.
	config.Consumer.Return.Errors = true
	// config.Consumer.Offsets - организует работу с коммитами.
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Вычитываем ли с самого старого оффсета в партиции
	// или начинаем с самого нового.
	// config.Consumer.IsolationLevel - 2 уровня изоляции при выполнении транзакции.
	// config.Consumer.Interceptors - настройка интерсепторов.

	client, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "groupOne", config)
	if err != nil {
		log.Fatalf("can't create consumer group")
	}

	ctx := context.Background()
	ch := ConsumerHandler{}

	for {
		if err := client.Consume(ctx, []string{topicName}, &ch); err != nil {
			log.Fatalf("consume message: %v", err)
		}
	}
}

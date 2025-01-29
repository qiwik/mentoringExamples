package main

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0  // Минимально поддерживаемая версия кафки.
	config.ClientID = "test-producer" // Идентификатор сервиса.

	config.Net.MaxOpenRequests = 10            // Сколько запросов выдерживает подключение.
	config.Net.DialTimeout = 30 * time.Second  // Таймаут на подключение.
	config.Net.ReadTimeout = 30 * time.Second  // Таймаут на ожидание ответа от брокера.
	config.Net.WriteTimeout = 30 * time.Second // Таймаут на ожидание передачи данных.

	// config.Producer.MaxMessageBytes - максимальный вес сообщения, который может отправлять
	// этот продюсер.
	config.Producer.RequiredAcks = sarama.WaitForLocal // Параметр описывает гарантии доставки. Продюсеру достаточно
	// или когда сообщение просто отправлено в брокера, или когда брокер-лидер записал сообщение, или когда еще и все
	// реплики сделали то же самое.
	// config.Producer.Timeout - длительность ожидания получения номера брокером. Хз, что это пока
	// config.Producer.Compression - алгоритм сжатия сообщения продюсером.
	// config.Producer.CompressionLevel - уровень сжатия сообщения.
	// config.Producer.Partitioner - кастомная логика для выбора партиции продюсером.
	config.Producer.Partitioner = NewPartition
	// config.Producer.Idempotent - флаг позволяет проверять, что сообщение уже было записано. То есть
	// помогает устранять дубликаты. Брокер проверяет порядковый номер сообщения, и если оно уже было, то
	// продюсер получит сообщение о дубликате и пойдет дальше. Важно, чтобы сообщения попадали в одни
	// и те же партиции.
	// config.Producer.Transaction - управление транзакциями.
	// config.Producer.Return - управление каналами внутри продюсера.
	config.Producer.Return.Successes = true
	// config.Producer.Flush - позволяет отправлять сообщения батчами вместо моментальной отправки.
	// config.Producer.Retry - позволяет управлять попытками переотправки сообщения.
	config.Producer.Retry.Max = 3 // Максимальное количество попыток.
	// config.Producer.Retry.Backoff - настраивает время между попытками отправки сообщения.
	// config.Producer.Interceptors - настраиваем интерсепторы.

	syncProducer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("create sync producer: " + err.Error())
	}

	testProducer := NewProducer(syncProducer)

	for i := range 10 {
		msg := Msg{
			ID:   i,
			Name: "Important message",
		}

		err = testProducer.Send(msg)
		if err != nil {
			log.Fatal("sending message to topic: " + err.Error())
		}

		//time.Sleep(1 * time.Minute)
	}
}

package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() (*ckafka.Producer, error) {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	p, err := ckafka.NewProducer(configMap)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChannel chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(msg),
	}
	err := producer.Produce(message, deliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChannel chan ckafka.Event) {
	for e := range deliveryChannel {
		switch event := e.(type) {
		case *ckafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery failed", event.TopicPartition)
			} else {
				fmt.Println("Delivered message to: ", event.TopicPartition)
			}

		}
	}
}

package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/luizeduu/imersao/codepix-go/application/factory"
	appModel "github.com/luizeduu/imersao/codepix-go/application/model"
	"github.com/luizeduu/imersao/codepix-go/application/usecase"
	"github.com/luizeduu/imersao/codepix-go/domain/model"
)

type KafkaProcessor struct {
	Database        *gorm.DB
	Producer        *ckafka.Producer
	DeliveryChannel chan ckafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChannel chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:        database,
		Producer:        producer,
		DeliveryChannel: deliveryChannel,
	}
}

func (k *KafkaProcessor) Consume(topics []string) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}
	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(topics, nil)

	log.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			k.processMessage(msg)
		}
	}
}

func (k *KafkaProcessor) processMessage(message *ckafka.Message) {
	transactionsTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *message.TopicPartition.Topic; topic {
	case transactionsTopic:
		k.processTransaction(message)
	case transactionConfirmationTopic:
		k.ProcessTransactionConfirmation(message)
	default:
		fmt.Println("not a valid topic", string(message.Value))
	}
}

func (k *KafkaProcessor) processTransaction(message *ckafka.Message) error {
	transaction := appModel.NewTransaction()

	err := transaction.ParseJson(message.Value)

	if err != nil {
		return err
	}

	registerTransactionUseCase := factory.RegisterTransactionUseCaseFactory(k.Database)

	createdTransaction, err := registerTransactionUseCase.Execute(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description,
	)

	if err != nil {
		fmt.Println("Error registering transaction", err)

		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = model.TransactionPending
	transactionJson, err := transaction.ToJson()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, k.Producer, k.DeliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaProcessor) ProcessTransactionConfirmation(message *ckafka.Message) error {
	transaction := appModel.NewTransaction()

	err := transaction.ParseJson(message.Value)

	if err != nil {
		return err
	}

	confirmTransactionUseCase := factory.ConfirmTransactionUseCaseFactory(k.Database)
	completeTransactionUseCase := factory.CompleteTransactionUseCaseFactory(k.Database)

	if transaction.Status == model.TransactionConfirmed {
		err = k.confirmTransaction(transaction, confirmTransactionUseCase)

		if err != nil {
			return err
		}

		return nil

	}

	_, err = completeTransactionUseCase.Execute(transaction.ID)

	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaProcessor) confirmTransaction(transaction *appModel.Transaction, confirmTransactionUseCase usecase.ConfirmTransactionUseCase) error {
	confirmedTransaction, err := confirmTransactionUseCase.Execute(transaction.ID)

	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJson()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, k.Producer, k.DeliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

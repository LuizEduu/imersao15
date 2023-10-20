/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/luizeduu/imersao/codepix-go/application/kafka"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",

	Run: func(cmd *cobra.Command, args []string) {
		producer, err := kafka.NewKafkaProducer()

		if err != nil {
			log.Fatal(err.Error())
		}

		deliveryChannel := make(chan ckafka.Event)
		err = kafka.Publish("Olá Kafka", "Test", producer, deliveryChannel)
		kafka.DeliveryReport(deliveryChannel)

		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package kafka

import (
	"log"
	"logAgent/config"

	"github.com/Shopify/sarama"
)


func InitKafka() (kafkaServerClient sarama.SyncProducer){
	kafkaServerConf := sarama.NewConfig()
	kafkaServerConf.Producer.RequiredAcks = sarama.WaitForAll
	kafkaServerConf.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaServerConf.Producer.Return.Successes = true

	var err error
	kafkaServerClient, err = sarama.NewSyncProducer(config.KafkaAddressList, kafkaServerConf)
	if err != nil {
		log.Fatalf("producer create failed, %s", err)
	}
	return
}



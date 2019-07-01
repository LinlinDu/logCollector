package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"log"
	"logCollector/config"
	"logCollector/elasticsearch"
	"sync"
)

var consumer sarama.Consumer

func InitKafka(){
	var err error
	consumer, err = sarama.NewConsumer(config.KafkaAddressList, nil)
	if err != nil {
		log.Fatal(err)
	}
	go createTopicTask()
	}


func createTopicTask() {
	var wg sync.WaitGroup
	partitionList, err := consumer.Partitions("log")
	if err != nil {
		logs.Error("get topic: [%s] partitions failed, err: %s", "log", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			logs.Warn("topic: [%s] start consumer partition failed, err: %s", "log", err)
			continue
		}

		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				elasticsearch.ESChan <- string(msg.Value)
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
}

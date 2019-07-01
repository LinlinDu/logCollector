package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"logAgent/config"
	"logAgent/kafka"
	"logAgent/logger"
	"logAgent/tail"
)

func init()  {
	config.InitConfig()
	logger.InitLogger()
}

type Msg struct {
	IP string
	Log string
}

func main() {
	tailClient := tail.InitTail()
	kafkaClient := kafka.InitKafka()

	logs.Info("Log Agent start running...")
	for line := range tailClient.Lines {
		if line.Text != ""{
		    sendToKafka(kafkaClient, line.Text)
	    }
	}
}

func sendToKafka(kafkaClient sarama.SyncProducer, log string) {
	msg := Msg{config.LocalIP, log}
	jmsg, err := json.Marshal(msg)
	if err != nil{
		logs.Error("send to kafka topic:[log] log:[%v] failed, %v", log, err)
		return
	}
	pid, offset, err := kafkaClient.SendMessage(&sarama.ProducerMessage{
		Topic:"log",
		Value:sarama.StringEncoder(jmsg),
	})

	if err != nil {
		logs.Error("send to kafka topic:[log] log:[%v] failed, %v", log, err)
		return
	}

	logs.Debug("topic: [log] pid: [%v], offset: [%v]", pid, offset)
}

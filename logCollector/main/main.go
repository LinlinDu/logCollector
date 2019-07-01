package main

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"logCollector/config"
	"logCollector/elasticsearch"
	"logCollector/kafka"
	"logCollector/logger"
)

func init()  {
	config.InitConfig()
	logger.InitLogger()
	kafka.InitKafka()
}

type Msg struct {
	IP string `json:"ip"`
	Log string `json:"log"`
}

func main() {
	client := elasticsearch.InitElastic()
	logs.Info("Log Collector start running...")
	for jmsg := range elasticsearch.ESChan{
		msg := Msg{}
		err := json.Unmarshal([]byte(jmsg), &msg)
		if err != nil{
			logs.Error(err)
		}
		_, err = client.Index().
			Index("log").
			BodyJson(msg).
			Do(context.Background())
		if err != nil {
			logs.Error("msg:%s failed to send to ES", msg)
		}
	}
}

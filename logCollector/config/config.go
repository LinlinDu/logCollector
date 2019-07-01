package config

import (
	"github.com/astaxie/beego/config"
	"log"
	"strings"
)

var(
	LogLevel     string
	KafkaAddressList []string
	ESAddress    string
)

func InitConfig() {

	conf, err := config.NewConfig("ini", "config/logCollector.ini")
	if err != nil {
		log.Fatal(err)
	}

	LogLevel = conf.String("base::log_level")
	if len(LogLevel) == 0 {
		log.Fatal("config log level failed")
	}

	KafkaAddressList = strings.Split(conf.String("kafka::kafka_address_list"),",")
	if len(KafkaAddressList) == 0 {
		log.Fatal("Transger config kafka address error")
	}

	ESAddress = conf.String("elasticsearch::es_address")
	if len(ESAddress) == 0 {
		log.Fatal("Transger config elasticsearch address error")
	}
}


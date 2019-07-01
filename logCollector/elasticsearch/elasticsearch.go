package elasticsearch

import (
"github.com/olivere/elastic"
"log"
"logCollector/config"
)

var ESChan = make(chan string, 1)

func InitElastic() (client *elastic.Client){
	var err error
	client, err = elastic.NewClient(elastic.SetURL(config.ESAddress))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

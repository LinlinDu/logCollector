package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"log"
	"net"
	"strings"
)

var(
	LocalIP      string
	LogLevel     string
	LogPath      string
	KafkaAddressList []string
)

func InitConfig() () {
	conf, err := config.NewConfig("ini", "config/logagent.ini")
	if err != nil {
		log.Fatal(err)
	}

	LocalIP, err = getLocalIP()
	if err != nil{
		log.Fatalf("logger config local IP error: %s", err)
	}

	LogLevel = conf.String("base::log_level")
	if len(LogLevel) == 0 {
		log.Fatalf("logger config logger level error: %s", err)

	}

	LogPath = conf.String("base::log_path")
	if len(LogPath) == 0 {
		log.Fatalf("logger config logger path error")
	}

	KafkaAddressList = strings.Split(conf.String("kafka::kafka_address_list"), ",")
	if len(KafkaAddressList) == 0 {
		log.Fatal("logger config kafka address error")
	}
}

func getLocalIP() (localIP string, err error){
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						localIP = ipnet.IP.String()
					}
				}
			}
		}
	}
	if localIP == ""{
		return "", fmt.Errorf("get local ip failed")
	}
	return
}

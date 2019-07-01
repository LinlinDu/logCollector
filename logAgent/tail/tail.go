package tail

import (
"log"
"github.com/hpcloud/tail"
"logAgent/config"
)

func InitTail() (t *tail.Tail){
	var err error
	t, err = tail.TailFile(config.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		log.Fatalf("tail init failed, err: %v", err)
	}
	return
}

package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetOutput(os.Stderr)
	//	log.SetLevel(logrus.InfoLevel)

	log.Infoln("[service] Start...")

	for i := 10; i > 0; i-- {
		if i%3 == 0 {
			log.Errorf("[service] wait %d...", i)
		} else {
			log.Printf("[service] wait %d...", i)
		}
		time.Sleep(time.Second)
	}

	log.Infoln("[service] End...")
}

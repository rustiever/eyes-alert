package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	mins = 20 * time.Minute
	secs = 20 * time.Second
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	minsTicker := time.NewTicker(mins)

	for {
		select {
		case <-sigs:
			os.Exit(0)
		case <-minsTicker.C:
			alert("Look 20 feet away")
			minsTicker.Stop()
			secsTicker := time.NewTicker(secs)
			select {
			case <-secsTicker.C:
				alert("Resume your awesome work")
				minsTicker.Reset(mins)
			case <-sigs:
				os.Exit(0)
			}
		}
	}
}

func alert(body string) {
	err := beeep.Alert("Eyes Alert", body, "")
	if err != nil {
		log.Fatalln("error in sending system-notification", err)
	}
}

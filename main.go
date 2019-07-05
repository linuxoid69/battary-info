package main

import (
	"flag"
	"github.com/linuxoid69/battery-info/dbus"
	"github.com/linuxoid69/battery-info/play"
	"log"
	"time"
)

var (
	p int
)

func init() {
	flag.IntVar(&p, "p", 15, "-p percentage")
	flag.Parse()
}

func main() {
//go:generate go-bindata -pkg play -o play/bindata.go sound/
	var second time.Duration
	second = 1
	for {
		sysdbus, err := dbus.SystemBusConnect()

		if err != nil {
			log.Println("Can't connect to system dbus")
		}

		sessiondbus, err := dbus.SessionBusConnect()

		if err != nil {
			log.Println("Can't connect to session dbus")
		}

		pp, err := dbus.GetPercentage(sysdbus)

		if err != nil {
			log.Println("Can't get percentage")
		}

		ac, err := dbus.GetACStatus(sysdbus)

		if err != nil {
			log.Println("Can't get status AC")
		}

		if p >= int(pp) && ac == false {
			err := dbus.BattNotification(sessiondbus, int(pp))
			if err != nil {
				log.Printf("Can't send notifycation: %s", err)
			}
			play.PlaySound()
			second = 60
		} else {
			second = 1
		}

		time.Sleep(time.Second * second)
	}
}

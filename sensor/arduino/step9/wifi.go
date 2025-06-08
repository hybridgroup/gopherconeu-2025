package main

import (
	"time"

	"tinygo.org/x/drivers/netlink"
	"tinygo.org/x/drivers/netlink/probe"
)

var (
	ssid, pass string
)

func connectWifi() {
	link, _ := probe.Probe()

	err := link.NetConnect(&netlink.ConnectParams{
		ConnectMode: netlink.ConnectModeAP,
		Ssid:        ssid,
		Passphrase:  pass,
	})
	if err != nil {
		failMessage(err.Error())
	}
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}

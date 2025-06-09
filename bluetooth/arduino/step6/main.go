package main

import (
	"fmt"
	"image/color"

	"tinygo.org/x/bluetooth"
	"tinygo.org/x/tinyterm"
)

var (
	terminal *tinyterm.Terminal

	black = color.RGBA{0, 0, 0, 255}
)

var DeviceAddress string

var (
	adapter = bluetooth.DefaultAdapter

	heartRateServiceUUID        = bluetooth.ServiceUUIDHeartRate
	heartRateCharacteristicUUID = bluetooth.CharacteristicUUIDHeartRateMeasurement
)

func main() {
	initTerminal()

	terminalOutput("enabling")

	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	ch := make(chan bluetooth.ScanResult, 1)

	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		terminalOutput(fmt.Sprintf("found device: %s %d %s", result.Address.String(), result.RSSI, result.LocalName()))
		if result.Address.String() == DeviceAddress {
			adapter.StopScan()
			ch <- result
		}
	})
	must("start scanning", err)

	var device bluetooth.Device
	select {
	case result := <-ch:
		device, err = adapter.Connect(result.Address, bluetooth.ConnectionParams{})
		if err != nil {
			terminalOutput(err.Error())
			return
		}

		terminalOutput("connected to " + result.Address.String())
	}

	// get services
	terminalOutput("discovering services/characteristics")
	srvcs, err := device.DiscoverServices([]bluetooth.UUID{heartRateServiceUUID})
	must("discover services", err)

	if len(srvcs) == 0 {
		panic("could not find heart rate service")
	}

	srvc := srvcs[0]

	terminalOutput("found service" + srvc.UUID().String())

	chars, err := srvc.DiscoverCharacteristics([]bluetooth.UUID{heartRateCharacteristicUUID})
	if err != nil {
		terminalOutput(err.Error())
	}

	if len(chars) == 0 {
		panic("could not find heart rate characteristic")
	}

	char := chars[0]
	terminalOutput("found characteristic" + char.UUID().String())

	char.EnableNotifications(func(buf []byte) {
		terminalOutput(fmt.Sprintf("data: %d", uint8(buf[1])))
	})

	select {}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}

func terminalOutput(s string) {
	println(s)
	fmt.Fprintf(terminal, "\n%s", s)

	terminal.Display()
}

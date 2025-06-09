package main

import (
	"fmt"
	"strconv"

	"tinygo.org/x/bluetooth"
	"tinygo.org/x/tinyterm"
)

var DeviceAddress string

var (
	terminal *tinyterm.Terminal

	adapter = bluetooth.DefaultAdapter
)

func main() {
	initTerminal()

	terminalOutput("enable interface...")

	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	ch := make(chan bluetooth.ScanResult, 1)

	// Start scanning.
	terminalOutput("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		terminalOutput(fmt.Sprintf("found device: %s %d %s", result.Address.String(), result.RSSI, result.LocalName()))
		if result.Address.String() == DeviceAddress {
			adapter.StopScan()
			ch <- result
		}
	})

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

	srvcs, err := device.DiscoverServices(nil)
	must("discover services", err)

	// buffer to retrieve characteristic data
	buf := make([]byte, 255)

	for _, srvc := range srvcs {
		terminalOutput("- srv " + srvc.UUID().String())

		chars, err := srvc.DiscoverCharacteristics(nil)
		if err != nil {
			terminalOutput(err.Error())
		}
		for _, char := range chars {
			terminalOutput("-- chr " + char.UUID().String())
			mtu, err := char.GetMTU()
			if err != nil {
				terminalOutput("  mtu: error: " + err.Error())
			} else {
				terminalOutput("  mtu: " + strconv.Itoa(int(mtu)))
			}
			n, err := char.Read(buf)
			if err != nil {
				terminalOutput("  " + err.Error())
			} else {
				terminalOutput("  data size " + strconv.Itoa(n))
				terminalOutput("  value = " + string(buf[:n]))
			}
		}
	}

	err = device.Disconnect()
	if err != nil {
		terminalOutput(err.Error())
	}

	terminalOutput("done")
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

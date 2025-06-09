# Bluetooth Tutorial

This series of tutorials is intended to help you learn about programming for Bluetooth using TinyGo and the TinyGo [Bluetooth package](https://github.com/tinygo-org/bluetooth)


## What you need

    - Arduino Nano RP2040 Connect IoT board
    - SSD1306 display and cables
    - Personal computer with Go 1.23+ and TinyGo installed, and a serial port.

## Code

### step0.go - Built-in LED

![Arduino Nano RP2040 Connect](../../sensor/arduino/assets/step0.jpg)

This tests that you can compile and flash your Arduino with TinyGo code, by blinking the built-in LED.

Run the following command to compile your code, and flash it onto the Arduino:

```
tinygo flash -target nano-rp2040 ./step0/
```

Once the Arduino is flashed correctly, the built-in amber LED to the right of the USB jack should start to turn on and off once per second. Now everything is setup correctly and you are ready to continue.

### step1.go - Bluetooth scan

First run some code to scan your local area for Bluetooth devices. This program will use the Bluetooth interface in your Arduino Nano RP2040 and show the output from it on your computer terminal using the USB interface.

Run the code.

```shell
tinygo flash -target nano-rp2040 -stack-size=8kb -monitor ./step1
```

You will see output on your terminal. Each of the devices list is a Bluetooth device nearby that is advertising itself.


### step2.go - Bluetooth scan on OLED display

Next step is to scan for Bluetooth device as in the previous step, but displaying the output on the OLED display.

We will add a SSD1306 OLED display to show the results from the Bluetooth scan. We will control this display using an [I2C interface](https://en.wikipedia.org/wiki/I%C2%B2C).

- Connect one of the "Ground" pins on the Arduino to the breadboard's ground rail (-) using a black or brown jumper cable.

- Connect the "3.3V" pin on the Arduino to the breadboard's power rail (+) using a red jumper cable.

- Connect a jumper wire from the "GND" pin on the breadboard next to the OLED display, to the breadboard's top left set of pins (-).

- Connect a jumper wire from the "VCC" pin on the breadboard next to the OLED display, to the breadboard's top right (+) set of pins.

- Connect a jumper wire from the "SDA" pin on the breadboard next to the OLED display, to the Arduino Nano RP2040 A4 pin.

- Connect a jumper wire from the "SCL" pin on the breadboard next to the OLED display, to the Arduino Nano RP2040 A5 pin.

We have 2 TinyGo packages to make it easier to use small displays such as the SSD1306 in the kit. 

The TinyFont package renders fonts to any of the supported displays in the TinyGo drivers repo. 

The TinyTerm package provides a terminal-style output on supported displays in the TinyGo drivers repo. 

Run the code.

```shell
tinygo flash -target nano-rp2040 -stack-size=8kb -monitor ./step2
```

You will see the bluetooth scan output on both your monitor and on the SSD1306 display.


### step3.go - Bluetooth discover

Now that you know how to find Bluetooth devices that are nearby you, you can proceed to try to connect to one of them and find out what services it can offer.

You will need to use the MAC address (Linux or Windows) or the Bluetooth ID (macOS) to connect to a specific device.

Try one of of the devices you found when you were scanning in step1/step2.

Run the code.

Note that not all devices will allow you to connect to them, and that some that allow you to connect will not allow you to view the details of every service/characteristic.

```shell
tinygo flash -target nano-rp2040 -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./step3
```

### step4.go - Bluetooth discover on nano-rp2040 display

This is the same Bluetooth service discovery as the previous example, but it also shows the data on the nano-rp2040 display.

Run the code.

```shell
tinygo flash -target nano-rp2040 -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./step4
```

You should see the output on both your terminal, and also on the nano-rp2040 display.


### step5.go - Bluetooth heart rate

Now that you know how to find Bluetooth devices that are nearby you and how to connect to them, you can proceed to try to do something useful.

Let's connect the nano-rp2040 to a Bluetooth heart rate sensor.

If you already have a smart watch or app on your phone that is a heart rate sensor, you can connect to it. Otherwise you can obtain one for your mobile device such as those listed here:

https://www.cnet.com/health/how-to-track-your-heart-rate-with-a-smartphone/

You can also run a simulator on your laptop computer:

```shell

go run ./tutorial/bluetooth/heartsim
```

Run the code.

```shell
tinygo flash -target nano-rp2040 -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./step5
```

You can connect from the nano-rp2040 to your mobile phone or any other device/software that can produce the data from a standard Bluetooth heart rate device.


### step6.go - Bluetooth heart rate on nano-rp2040 display

Run the code.

```shell
tinygo flash -target nano-rp2040 -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./step6
```

This is the same heart rate device as the previous example, but it also shows the data on the nano-rp2040 display. You will still need to connect to your mobile phone or any other device/software that can produce the data for a standard Bluetooth heart rate device.

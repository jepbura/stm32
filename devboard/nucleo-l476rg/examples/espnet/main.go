// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espnet is ESP-AT based TCP echo server. Because of the insufficient RAM in
// STM32L476 it may work unstable. See ../espat for less memory consuming
// version of this program that uses the espat package directly. See also
// the same example written for F4-Discovery and other development boards.
package main

import (
	"io"
	"net"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"
)

func logErr(err error) bool {
	if err != nil {
		if err != io.EOF {
			println("error:", err.Error())
		}
		return true
	}
	return false
}

func fatalErr(err error) {
	for err != nil {
		println("error:", err.Error())
		time.Sleep(time.Second)
	}
}

func main() {
	system.Setup80(0, 0)
	rtcst.Setup(rtcst.LSE, 1, 32768)

	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	tx := pa.Pin(9)  // CN5 D8
	rx := pa.Pin(10) // CN9 D2

	// Configure and enable USART
	u := usart1.Driver()
	u.UsePin(tx, usart.TXD)
	u.UsePin(rx, usart.RXD)
	u.Setup(usart.Word8b, 115200)
	u.EnableTx()
	u.EnableRx(256)

	print("\n* L7 ready *\n\n")

	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(false))
	fatalErr(espnet.SetPasvRecv(dev, true))

	/*
		for msg := range dev.Async() {
			if msg == "WIFI GOT IP" {
				break
			}
		}
	*/

	ls, err := espnet.ListenDev(dev, "tcp", 1111)
	fatalErr(err)

	println("listen on:", ls.Addr().String())
	for {
		c, err := ls.Accept()
		fatalErr(err)
		go handle(c)
	}
}

func handle(c net.Conn) {
	var buf [64]byte
	println("connect:", c.RemoteAddr().String())
	_, err := io.WriteString(c, "Echo Server\n\n")
	if logErr(err) {
		goto end
	}
	for {
		n, err := c.Read(buf[:])
		if logErr(err) {
			goto end
		}
		_, err = c.Write(buf[:n])
		if logErr(err) {
			goto end
		}
	}
end:
	c.Close()
	println("close:  ", c.RemoteAddr().String())
}

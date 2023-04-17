// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espat is ESP-AT based TCP echo server. It uses the espat package directly
// (instead of espat/espnet) because of the insufficient RAM in STM32L476. See
// ../espnet for the example of same TCP server implemented using espat/espnet.
package main

import (
	"io"
	"time"

	"github.com/embeddedgo/espat"

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

	d := espat.NewDevice("esp0", u, u)
	fatalErr(d.Init(false))
	_, err := d.Cmd("+CIPMUX=1")
	fatalErr(err)
	_, err = d.Cmd("+CIPRECVMODE=1")
	fatalErr(err)

	/*
		for msg := range dev.Async() {
			if msg == "WIFI GOT IP" {
				break
			}
		}
	*/

	d.SetServer(true)
	_, err = d.Cmd("+CIPSERVER=1,1111")
	fatalErr(err)

	println("listen on :1111")
	for conn := range d.Server() {
		go handle(d, conn)
	}
}

var welcome = []byte("Echo Server\n\n")

func handle(d *espat.Device, conn *espat.Conn) {
	println("connect", conn.ID)
	if logErr(send(d, conn, welcome)) {
		return
	}
	var buf [64]byte
	for {
		if _, ok := <-conn.Ch; !ok {
			goto end // connection closed by remote part
		}
		n, err := d.CmdInt("+CIPRECVDATA=", buf[:], conn.ID, len(buf))
		if logErr(err) {
			return
		}
		if logErr(send(d, conn, buf[:n])) {
			return
		}
	}
end:
	d.Cmd("+CIPCLOSE=", conn.ID)
	println("close", conn.ID)
}

func send(d *espat.Device, conn *espat.Conn, p []byte) error {
	d.Lock()
	defer d.Unlock()
	if _, err := d.UnsafeCmd("+CIPSEND=", conn.ID, len(p)); err != nil {
		return err
	}
	if _, err := d.UnsafeWrite(p); err != nil {
		return err
	}
	_, err := d.UnsafeCmd("")
	return err
}

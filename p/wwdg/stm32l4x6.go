// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package wwdg provides access to the registers of the WWDG peripheral.
//
// Instances:
//  WWDG  WWDG_BASE  APB1  WWDG  System window watchdog
// Registers:
//  0x000 32  CR   Control register
//  0x004 32  CFR  Configuration register
//  0x008 32  SR   Status register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package wwdg

const (
	T    CR = 0x7F << 0 //+ 7-bit counter (MSB to LSB)
	WDGA CR = 0x01 << 7 //+ Activation bit
)

const (
	Tn    = 0
	WDGAn = 7
)

const (
	W     CFR = 0x7F << 0 //+ 7-bit window value
	WDGTB CFR = 0x03 << 7 //+ Timer base
	EWI   CFR = 0x01 << 9 //+ Early wakeup interrupt
)

const (
	Wn     = 0
	WDGTBn = 7
	EWIn   = 9
)

const (
	EWIF SR = 0x01 << 0 //+ Early wakeup interrupt flag
)

const (
	EWIFn = 0
)

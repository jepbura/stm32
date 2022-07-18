// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package rng provides access to the registers of the RNG peripheral.
//
// Instances:
//  RNG  RNG_BASE  AHB2  HASH_RNG  Random number generator
// Registers:
//  0x000 32  CR  control register
//  0x004 32  SR  status register
//  0x008 32  DR  data register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package rng

const (
	RNGEN CR = 0x01 << 2 //+ Random number generator enable
	IE    CR = 0x01 << 3 //+ Interrupt enable
)

const (
	RNGENn = 2
	IEn    = 3
)

const (
	DRDY SR = 0x01 << 0 //+ Data ready
	CECS SR = 0x01 << 1 //+ Clock error current status
	SECS SR = 0x01 << 2 //+ Seed error current status
	CEIS SR = 0x01 << 5 //+ Clock error interrupt status
	SEIS SR = 0x01 << 6 //+ Seed error interrupt status
)

const (
	DRDYn = 0
	CECSn = 1
	SECSn = 2
	CEISn = 5
	SEISn = 6
)

const (
	RNDATA DR = 0xFFFFFFFF << 0 //+ Random data
)

const (
	RNDATAn = 0
)

// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package vrefbuf provides access to the registers of the VREFBUF peripheral.
//
// Instances:
//  VREFBUF  VREFBUF_BASE  -  -  Voltage reference buffer
// Registers:
//  0x000 32  VREFBUF_CSR  VREF_BUF Control and Status Register
//  0x004 32  VREFBUF_CCR  VREF_BUF Calibration Control Register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package vrefbuf

const (
	ENVR VREFBUF_CSR = 0x01 << 0 //+ Enable Voltage Reference
	HIZ  VREFBUF_CSR = 0x01 << 1 //+ High impedence mode for the VREF_BUF
	VRR  VREFBUF_CSR = 0x01 << 3 //+ Voltage reference buffer ready
	VRS  VREFBUF_CSR = 0x03 << 4 //+ Voltage reference scale
)

const (
	ENVRn = 0
	HIZn  = 1
	VRRn  = 3
	VRSn  = 4
)

const (
	TRIM VREFBUF_CCR = 0x3F << 0 //+ Trimming code
)

const (
	TRIMn = 0
)
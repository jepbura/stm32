// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package mpu provides access to the registers of the MPU peripheral.
//
// Instances:
//  MPU  MPU_BASE  -  -  Memory protection unit
// Registers:
//  0x000 32  TYPER  MPU type register
//  0x004 32  CTRL   MPU control register
//  0x008 32  RNR    MPU region number register
//  0x00C 32  RBAR   MPU region base address register
//  0x010 32  RASR   MPU region attribute and size register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package mpu

const (
	SEPARATE TYPER = 0x01 << 0  //+ Separate flag
	DREGION  TYPER = 0xFF << 8  //+ Number of MPU data regions
	IREGION  TYPER = 0xFF << 16 //+ Number of MPU instruction regions
)

const (
	SEPARATEn = 0
	DREGIONn  = 8
	IREGIONn  = 16
)

const (
	ENABLE     CTRL = 0x01 << 0 //+ Enables the MPU
	HFNMIENA   CTRL = 0x01 << 1 //+ Enables the operation of MPU during hard fault
	PRIVDEFENA CTRL = 0x01 << 2 //+ Enable priviliged software access to default memory map
)

const (
	ENABLEn     = 0
	HFNMIENAn   = 1
	PRIVDEFENAn = 2
)

const (
	REGION RNR = 0xFF << 0 //+ MPU region
)

const (
	REGIONn = 0
)

const (
	REGION RBAR = 0x0F << 0      //+ MPU region field
	VALID  RBAR = 0x01 << 4      //+ MPU region number valid
	ADDR   RBAR = 0x7FFFFFF << 5 //+ Region base address field
)

const (
	REGIONn = 0
	VALIDn  = 4
	ADDRn   = 5
)

const (
	ENABLE RASR = 0x01 << 0  //+ Region enable bit.
	SIZE   RASR = 0x1F << 1  //+ Size of the MPU protection region
	SRD    RASR = 0xFF << 8  //+ Subregion disable bits
	B      RASR = 0x01 << 16 //+ memory attribute
	C      RASR = 0x01 << 17 //+ memory attribute
	S      RASR = 0x01 << 18 //+ Shareable memory attribute
	TEX    RASR = 0x07 << 19 //+ memory attribute
	AP     RASR = 0x07 << 24 //+ Access permission
	XN     RASR = 0x01 << 28 //+ Instruction access disable bit
)

const (
	ENABLEn = 0
	SIZEn   = 1
	SRDn    = 8
	Bn      = 16
	Cn      = 17
	Sn      = 18
	TEXn    = 19
	APn     = 24
	XNn     = 28
)
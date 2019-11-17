// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package opamp provides access to the registers of the OPAMP peripheral.
//
// Instances:
//  OPAMP  OPAMP_BASE  APB1  -  Operational amplifiers
// Registers:
//  0x000 32  OPAMP1_CSR    OPAMP1 control/status register
//  0x004 32  OPAMP1_OTR    OPAMP1 offset trimming register in normal mode
//  0x008 32  OPAMP1_LPOTR  OPAMP1 offset trimming register in low-power mode
//  0x010 32  OPAMP2_CSR    OPAMP2 control/status register
//  0x014 32  OPAMP2_OTR    OPAMP2 offset trimming register in normal mode
//  0x018 32  OPAMP2_LPOTR  OPAMP2 offset trimming register in low-power mode
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package opamp

const (
	OPAEN     OPAMP1_CSR = 0x01 << 0  //+ Operational amplifier Enable
	OPALPM    OPAMP1_CSR = 0x01 << 1  //+ Operational amplifier Low Power Mode
	OPAMODE   OPAMP1_CSR = 0x03 << 2  //+ Operational amplifier PGA mode
	PGA_GAIN  OPAMP1_CSR = 0x03 << 4  //+ Operational amplifier Programmable amplifier gain value
	VM_SEL    OPAMP1_CSR = 0x03 << 8  //+ Inverting input selection
	VP_SEL    OPAMP1_CSR = 0x01 << 10 //+ Non inverted input selection
	CALON     OPAMP1_CSR = 0x01 << 12 //+ Calibration mode enabled
	CALSEL    OPAMP1_CSR = 0x01 << 13 //+ Calibration selection
	USERTRIM  OPAMP1_CSR = 0x01 << 14 //+ allows to switch from AOP offset trimmed values to AOP offset
	CALOUT    OPAMP1_CSR = 0x01 << 15 //+ Operational amplifier calibration output
	OPA_RANGE OPAMP1_CSR = 0x01 << 31 //+ Operational amplifier power supply range for stability
)

const (
	OPAENn     = 0
	OPALPMn    = 1
	OPAMODEn   = 2
	PGA_GAINn  = 4
	VM_SELn    = 8
	VP_SELn    = 10
	CALONn     = 12
	CALSELn    = 13
	USERTRIMn  = 14
	CALOUTn    = 15
	OPA_RANGEn = 31
)

const (
	TRIMOFFSETN OPAMP1_OTR = 0x1F << 0 //+ Trim for NMOS differential pairs
	TRIMOFFSETP OPAMP1_OTR = 0x1F << 8 //+ Trim for PMOS differential pairs
)

const (
	TRIMOFFSETNn = 0
	TRIMOFFSETPn = 8
)

const (
	TRIMLPOFFSETN OPAMP1_LPOTR = 0x1F << 0 //+ Trim for NMOS differential pairs
	TRIMLPOFFSETP OPAMP1_LPOTR = 0x1F << 8 //+ Trim for PMOS differential pairs
)

const (
	TRIMLPOFFSETNn = 0
	TRIMLPOFFSETPn = 8
)

const (
	OPAEN    OPAMP2_CSR = 0x01 << 0  //+ Operational amplifier Enable
	OPALPM   OPAMP2_CSR = 0x01 << 1  //+ Operational amplifier Low Power Mode
	OPAMODE  OPAMP2_CSR = 0x03 << 2  //+ Operational amplifier PGA mode
	PGA_GAIN OPAMP2_CSR = 0x03 << 4  //+ Operational amplifier Programmable amplifier gain value
	VM_SEL   OPAMP2_CSR = 0x03 << 8  //+ Inverting input selection
	VP_SEL   OPAMP2_CSR = 0x01 << 10 //+ Non inverted input selection
	CALON    OPAMP2_CSR = 0x01 << 12 //+ Calibration mode enabled
	CALSEL   OPAMP2_CSR = 0x01 << 13 //+ Calibration selection
	USERTRIM OPAMP2_CSR = 0x01 << 14 //+ allows to switch from AOP offset trimmed values to AOP offset
	CALOUT   OPAMP2_CSR = 0x01 << 15 //+ Operational amplifier calibration output
)

const (
	OPAENn    = 0
	OPALPMn   = 1
	OPAMODEn  = 2
	PGA_GAINn = 4
	VM_SELn   = 8
	VP_SELn   = 10
	CALONn    = 12
	CALSELn   = 13
	USERTRIMn = 14
	CALOUTn   = 15
)

const (
	TRIMOFFSETN OPAMP2_OTR = 0x1F << 0 //+ Trim for NMOS differential pairs
	TRIMOFFSETP OPAMP2_OTR = 0x1F << 8 //+ Trim for PMOS differential pairs
)

const (
	TRIMOFFSETNn = 0
	TRIMOFFSETPn = 8
)

const (
	TRIMLPOFFSETN OPAMP2_LPOTR = 0x1F << 0 //+ Trim for NMOS differential pairs
	TRIMLPOFFSETP OPAMP2_LPOTR = 0x1F << 8 //+ Trim for PMOS differential pairs
)

const (
	TRIMLPOFFSETNn = 0
	TRIMLPOFFSETPn = 8
)
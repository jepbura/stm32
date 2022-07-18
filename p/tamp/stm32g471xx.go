// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package tamp provides access to the registers of the TAMP peripheral.
//
// Instances:
//  TAMP  TAMP_BASE  -  -  Tamper and backup registers
// Registers:
//  0x000 32  CR1     control register 1
//  0x004 32  CR2     control register 2
//  0x00C 32  FLTCR   TAMP filter control register
//  0x02C 32  IER     TAMP interrupt enable register
//  0x030 32  SR      TAMP status register
//  0x034 32  MISR    TAMP masked interrupt status register
//  0x03C 32  SCR     TAMP status clear register
//  0x100 32  BKP0R   TAMP backup register
//  0x104 32  BKP1R   TAMP backup register
//  0x108 32  BKP2R   TAMP backup register
//  0x10C 32  BKP3R   TAMP backup register
//  0x110 32  BKP4R   TAMP backup register
//  0x114 32  BKP5R   TAMP backup register
//  0x118 32  BKP6R   TAMP backup register
//  0x11C 32  BKP7R   TAMP backup register
//  0x120 32  BKP8R   TAMP backup register
//  0x124 32  BKP9R   TAMP backup register
//  0x128 32  BKP10R  TAMP backup register
//  0x12C 32  BKP11R  TAMP backup register
//  0x130 32  BKP12R  TAMP backup register
//  0x134 32  BKP13R  TAMP backup register
//  0x138 32  BKP14R  TAMP backup register
//  0x13C 32  BKP15R  TAMP backup register
//  0x140 32  BKP16R  TAMP backup register
//  0x144 32  BKP17R  TAMP backup register
//  0x148 32  BKP18R  TAMP backup register
//  0x14C 32  BKP19R  TAMP backup register
//  0x150 32  BKP20R  TAMP backup register
//  0x154 32  BKP21R  TAMP backup register
//  0x158 32  BKP22R  TAMP backup register
//  0x15C 32  BKP23R  TAMP backup register
//  0x160 32  BKP24R  TAMP backup register
//  0x164 32  BKP25R  TAMP backup register
//  0x168 32  BKP26R  TAMP backup register
//  0x16C 32  BKP27R  TAMP backup register
//  0x170 32  BKP28R  TAMP backup register
//  0x174 32  BKP29R  TAMP backup register
//  0x178 32  BKP30R  TAMP backup register
//  0x17C 32  BKP31R  TAMP backup register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package tamp

const (
	TAMP1E  CR1 = 0x01 << 0  //+ TAMP1E
	TAMP2E  CR1 = 0x01 << 1  //+ TAMP2E
	TAMP3E  CR1 = 0x01 << 2  //+ TAMP2E
	ITAMP3E CR1 = 0x01 << 18 //+ ITAMP3E
	ITAMP4E CR1 = 0x01 << 19 //+ ITAMP4E
	ITAMP5E CR1 = 0x01 << 20 //+ ITAMP5E
	ITAMP6E CR1 = 0x01 << 21 //+ ITAMP6E
)

const (
	TAMP1En  = 0
	TAMP2En  = 1
	TAMP3En  = 2
	ITAMP3En = 18
	ITAMP4En = 19
	ITAMP5En = 20
	ITAMP6En = 21
)

const (
	TAMP1NOER CR2 = 0x01 << 0  //+ TAMP1NOER
	TAMP2NOER CR2 = 0x01 << 1  //+ TAMP2NOER
	TAMP3NOER CR2 = 0x01 << 2  //+ TAMP3NOER
	TAMP1MSK  CR2 = 0x01 << 16 //+ TAMP1MSK
	TAMP2MSK  CR2 = 0x01 << 17 //+ TAMP2MSK
	TAMP3MSK  CR2 = 0x01 << 18 //+ TAMP3MSK
	TAMP1TRG  CR2 = 0x01 << 24 //+ TAMP1TRG
	TAMP2TRG  CR2 = 0x01 << 25 //+ TAMP2TRG
	TAMP3TRG  CR2 = 0x01 << 26 //+ TAMP3TRG
)

const (
	TAMP1NOERn = 0
	TAMP2NOERn = 1
	TAMP3NOERn = 2
	TAMP1MSKn  = 16
	TAMP2MSKn  = 17
	TAMP3MSKn  = 18
	TAMP1TRGn  = 24
	TAMP2TRGn  = 25
	TAMP3TRGn  = 26
)

const (
	TAMPFREQ  FLTCR = 0x07 << 0 //+ TAMPFREQ
	TAMPFLT   FLTCR = 0x03 << 3 //+ TAMPFLT
	TAMPPRCH  FLTCR = 0x03 << 5 //+ TAMPPRCH
	TAMPPUDIS FLTCR = 0x01 << 7 //+ TAMPPUDIS
)

const (
	TAMPFREQn  = 0
	TAMPFLTn   = 3
	TAMPPRCHn  = 5
	TAMPPUDISn = 7
)

const (
	TAMP1IE  IER = 0x01 << 0  //+ TAMP1IE
	TAMP2IE  IER = 0x01 << 1  //+ TAMP2IE
	TAMP3IE  IER = 0x01 << 2  //+ TAMP3IE
	ITAMP3IE IER = 0x01 << 18 //+ ITAMP3IE
	ITAMP4IE IER = 0x01 << 19 //+ ITAMP4IE
	ITAMP5IE IER = 0x01 << 20 //+ ITAMP5IE
	ITAMP6IE IER = 0x01 << 21 //+ ITAMP6IE
)

const (
	TAMP1IEn  = 0
	TAMP2IEn  = 1
	TAMP3IEn  = 2
	ITAMP3IEn = 18
	ITAMP4IEn = 19
	ITAMP5IEn = 20
	ITAMP6IEn = 21
)

const (
	TAMP1F  SR = 0x01 << 0  //+ TAMP1F
	TAMP2F  SR = 0x01 << 1  //+ TAMP2F
	TAMP3F  SR = 0x01 << 2  //+ TAMP3F
	ITAMP3F SR = 0x01 << 18 //+ ITAMP3F
	ITAMP4F SR = 0x01 << 19 //+ ITAMP4F
	ITAMP5F SR = 0x01 << 20 //+ ITAMP5F
	ITAMP6F SR = 0x01 << 21 //+ ITAMP6F
)

const (
	TAMP1Fn  = 0
	TAMP2Fn  = 1
	TAMP3Fn  = 2
	ITAMP3Fn = 18
	ITAMP4Fn = 19
	ITAMP5Fn = 20
	ITAMP6Fn = 21
)

const (
	TAMP1MF  MISR = 0x01 << 0  //+ TAMP1MF:
	TAMP2MF  MISR = 0x01 << 1  //+ TAMP2MF
	TAMP3MF  MISR = 0x01 << 2  //+ TAMP3MF
	ITAMP3MF MISR = 0x01 << 18 //+ ITAMP3MF
	ITAMP4MF MISR = 0x01 << 19 //+ ITAMP4MF
	ITAMP5MF MISR = 0x01 << 20 //+ ITAMP5MF
	ITAMP6MF MISR = 0x01 << 21 //+ ITAMP6MF
)

const (
	TAMP1MFn  = 0
	TAMP2MFn  = 1
	TAMP3MFn  = 2
	ITAMP3MFn = 18
	ITAMP4MFn = 19
	ITAMP5MFn = 20
	ITAMP6MFn = 21
)

const (
	CTAMP1F  SCR = 0x01 << 0  //+ CTAMP1F
	CTAMP2F  SCR = 0x01 << 1  //+ CTAMP2F
	CTAMP3F  SCR = 0x01 << 2  //+ CTAMP3F
	CITAMP3F SCR = 0x01 << 18 //+ CITAMP3F
	CITAMP4F SCR = 0x01 << 19 //+ CITAMP4F
	CITAMP5F SCR = 0x01 << 20 //+ CITAMP5F
	CITAMP6F SCR = 0x01 << 21 //+ CITAMP6F
)

const (
	CTAMP1Fn  = 0
	CTAMP2Fn  = 1
	CTAMP3Fn  = 2
	CITAMP3Fn = 18
	CITAMP4Fn = 19
	CITAMP5Fn = 20
	CITAMP6Fn = 21
)

const (
	BKP BKP0R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP1R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP2R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP3R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP4R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP5R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP6R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP7R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP8R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP9R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP10R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP11R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP12R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP13R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP14R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP15R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP16R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP17R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP18R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP19R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP20R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP21R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP22R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP23R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP24R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP25R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP26R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP27R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP28R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP29R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP30R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

const (
	BKP BKP31R = 0xFFFFFFFF << 0 //+ BKP
)

const (
	BKPn = 0
)

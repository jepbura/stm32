// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f407

// Package fsmc provides access to the registers of the FSMC peripheral.
//
// Instances:
//  FSMC  FSMC_BASE  AHB3  FSMC  Flexible static memory controller
// Registers:
//  0x000 32  BCT{CR,TR}[4]  chip-select control and timing registers
//  0x060 32  PCR2           PC Card/NAND Flash control register 2
//  0x064 32  SR2            FIFO status and interrupt register 2
//  0x068 32  PMEM2          Common memory space timing register 2
//  0x06C 32  PATT2          Attribute memory space timing register 2
//  0x074 32  ECCR2          ECC result register 2
//  0x080 32  PCR3           PC Card/NAND Flash control register 3
//  0x084 32  SR3            FIFO status and interrupt register 3
//  0x088 32  PMEM3          Common memory space timing register 3
//  0x08C 32  PATT3          Attribute memory space timing register 3
//  0x094 32  ECCR3          ECC result register 3
//  0x0A0 32  PCR4           PC Card/NAND Flash control register 4
//  0x0A4 32  SR4            FIFO status and interrupt register 4
//  0x0A8 32  PMEM4          Common memory space timing register 4
//  0x0AC 32  PATT4          Attribute memory space timing register 4
//  0x0B0 32  PIO4           I/O space timing register 4
//  0x104 32  BWTR[4]        write timing registers
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package fsmc

const (
	MBKEN     CR = 0x01 << 0  //+ MBKEN
	MUXEN     CR = 0x01 << 1  //+ MUXEN
	MTYP      CR = 0x03 << 2  //+ MTYP
	MWID      CR = 0x03 << 4  //+ MWID
	FACCEN    CR = 0x01 << 6  //+ FACCEN
	BURSTEN   CR = 0x01 << 8  //+ BURSTEN
	WAITPOL   CR = 0x01 << 9  //+ WAITPOL
	WAITCFG   CR = 0x01 << 11 //+ WAITCFG
	WREN      CR = 0x01 << 12 //+ WREN
	WAITEN    CR = 0x01 << 13 //+ WAITEN
	EXTMOD    CR = 0x01 << 14 //+ EXTMOD
	ASYNCWAIT CR = 0x01 << 15 //+ ASYNCWAIT
	CBURSTRW  CR = 0x01 << 19 //+ CBURSTRW
)

const (
	MBKENn     = 0
	MUXENn     = 1
	MTYPn      = 2
	MWIDn      = 4
	FACCENn    = 6
	BURSTENn   = 8
	WAITPOLn   = 9
	WAITCFGn   = 11
	WRENn      = 12
	WAITENn    = 13
	EXTMODn    = 14
	ASYNCWAITn = 15
	CBURSTRWn  = 19
)

const (
	ADDSET  TR = 0x0F << 0  //+ ADDSET
	ADDHLD  TR = 0x0F << 4  //+ ADDHLD
	DATAST  TR = 0xFF << 8  //+ DATAST
	BUSTURN TR = 0x0F << 16 //+ BUSTURN
	CLKDIV  TR = 0x0F << 20 //+ CLKDIV
	DATLAT  TR = 0x0F << 24 //+ DATLAT
	ACCMOD  TR = 0x03 << 28 //+ ACCMOD
)

const (
	ADDSETn  = 0
	ADDHLDn  = 4
	DATASTn  = 8
	BUSTURNn = 16
	CLKDIVn  = 20
	DATLATn  = 24
	ACCMODn  = 28
)

const (
	PWAITEN PCR2 = 0x01 << 1  //+ PWAITEN
	PBKEN   PCR2 = 0x01 << 2  //+ PBKEN
	PTYP    PCR2 = 0x01 << 3  //+ PTYP
	PWID    PCR2 = 0x03 << 4  //+ PWID
	ECCEN   PCR2 = 0x01 << 6  //+ ECCEN
	TCLR    PCR2 = 0x0F << 9  //+ TCLR
	TAR     PCR2 = 0x0F << 13 //+ TAR
	ECCPS   PCR2 = 0x07 << 17 //+ ECCPS
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR2 = 0x01 << 0 //+ IRS
	ILS   SR2 = 0x01 << 1 //+ ILS
	IFS   SR2 = 0x01 << 2 //+ IFS
	IREN  SR2 = 0x01 << 3 //+ IREN
	ILEN  SR2 = 0x01 << 4 //+ ILEN
	IFEN  SR2 = 0x01 << 5 //+ IFEN
	FEMPT SR2 = 0x01 << 6 //+ FEMPT
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSETx  PMEM2 = 0xFF << 0  //+ MEMSETx
	MEMWAITx PMEM2 = 0xFF << 8  //+ MEMWAITx
	MEMHOLDx PMEM2 = 0xFF << 16 //+ MEMHOLDx
	MEMHIZx  PMEM2 = 0xFF << 24 //+ MEMHIZx
)

const (
	MEMSETxn  = 0
	MEMWAITxn = 8
	MEMHOLDxn = 16
	MEMHIZxn  = 24
)

const (
	ATTSETx  PATT2 = 0xFF << 0  //+ ATTSETx
	ATTWAITx PATT2 = 0xFF << 8  //+ ATTWAITx
	ATTHOLDx PATT2 = 0xFF << 16 //+ ATTHOLDx
	ATTHIZx  PATT2 = 0xFF << 24 //+ ATTHIZx
)

const (
	ATTSETxn  = 0
	ATTWAITxn = 8
	ATTHOLDxn = 16
	ATTHIZxn  = 24
)

const (
	ECCx ECCR2 = 0xFFFFFFFF << 0 //+ ECCx
)

const (
	ECCxn = 0
)

const (
	PWAITEN PCR3 = 0x01 << 1  //+ PWAITEN
	PBKEN   PCR3 = 0x01 << 2  //+ PBKEN
	PTYP    PCR3 = 0x01 << 3  //+ PTYP
	PWID    PCR3 = 0x03 << 4  //+ PWID
	ECCEN   PCR3 = 0x01 << 6  //+ ECCEN
	TCLR    PCR3 = 0x0F << 9  //+ TCLR
	TAR     PCR3 = 0x0F << 13 //+ TAR
	ECCPS   PCR3 = 0x07 << 17 //+ ECCPS
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR3 = 0x01 << 0 //+ IRS
	ILS   SR3 = 0x01 << 1 //+ ILS
	IFS   SR3 = 0x01 << 2 //+ IFS
	IREN  SR3 = 0x01 << 3 //+ IREN
	ILEN  SR3 = 0x01 << 4 //+ ILEN
	IFEN  SR3 = 0x01 << 5 //+ IFEN
	FEMPT SR3 = 0x01 << 6 //+ FEMPT
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSETx  PMEM3 = 0xFF << 0  //+ MEMSETx
	MEMWAITx PMEM3 = 0xFF << 8  //+ MEMWAITx
	MEMHOLDx PMEM3 = 0xFF << 16 //+ MEMHOLDx
	MEMHIZx  PMEM3 = 0xFF << 24 //+ MEMHIZx
)

const (
	MEMSETxn  = 0
	MEMWAITxn = 8
	MEMHOLDxn = 16
	MEMHIZxn  = 24
)

const (
	ATTSETx  PATT3 = 0xFF << 0  //+ ATTSETx
	ATTWAITx PATT3 = 0xFF << 8  //+ ATTWAITx
	ATTHOLDx PATT3 = 0xFF << 16 //+ ATTHOLDx
	ATTHIZx  PATT3 = 0xFF << 24 //+ ATTHIZx
)

const (
	ATTSETxn  = 0
	ATTWAITxn = 8
	ATTHOLDxn = 16
	ATTHIZxn  = 24
)

const (
	ECCx ECCR3 = 0xFFFFFFFF << 0 //+ ECCx
)

const (
	ECCxn = 0
)

const (
	PWAITEN PCR4 = 0x01 << 1  //+ PWAITEN
	PBKEN   PCR4 = 0x01 << 2  //+ PBKEN
	PTYP    PCR4 = 0x01 << 3  //+ PTYP
	PWID    PCR4 = 0x03 << 4  //+ PWID
	ECCEN   PCR4 = 0x01 << 6  //+ ECCEN
	TCLR    PCR4 = 0x0F << 9  //+ TCLR
	TAR     PCR4 = 0x0F << 13 //+ TAR
	ECCPS   PCR4 = 0x07 << 17 //+ ECCPS
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR4 = 0x01 << 0 //+ IRS
	ILS   SR4 = 0x01 << 1 //+ ILS
	IFS   SR4 = 0x01 << 2 //+ IFS
	IREN  SR4 = 0x01 << 3 //+ IREN
	ILEN  SR4 = 0x01 << 4 //+ ILEN
	IFEN  SR4 = 0x01 << 5 //+ IFEN
	FEMPT SR4 = 0x01 << 6 //+ FEMPT
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSETx  PMEM4 = 0xFF << 0  //+ MEMSETx
	MEMWAITx PMEM4 = 0xFF << 8  //+ MEMWAITx
	MEMHOLDx PMEM4 = 0xFF << 16 //+ MEMHOLDx
	MEMHIZx  PMEM4 = 0xFF << 24 //+ MEMHIZx
)

const (
	MEMSETxn  = 0
	MEMWAITxn = 8
	MEMHOLDxn = 16
	MEMHIZxn  = 24
)

const (
	ATTSETx  PATT4 = 0xFF << 0  //+ ATTSETx
	ATTWAITx PATT4 = 0xFF << 8  //+ ATTWAITx
	ATTHOLDx PATT4 = 0xFF << 16 //+ ATTHOLDx
	ATTHIZx  PATT4 = 0xFF << 24 //+ ATTHIZx
)

const (
	ATTSETxn  = 0
	ATTWAITxn = 8
	ATTHOLDxn = 16
	ATTHIZxn  = 24
)

const (
	IOSETx  PIO4 = 0xFF << 0  //+ IOSETx
	IOWAITx PIO4 = 0xFF << 8  //+ IOWAITx
	IOHOLDx PIO4 = 0xFF << 16 //+ IOHOLDx
	IOHIZx  PIO4 = 0xFF << 24 //+ IOHIZx
)

const (
	IOSETxn  = 0
	IOWAITxn = 8
	IOHOLDxn = 16
	IOHIZxn  = 24
)

const (
	ADDSET BWTR = 0x0F << 0  //+ ADDSET
	ADDHLD BWTR = 0x0F << 4  //+ ADDHLD
	DATAST BWTR = 0xFF << 8  //+ DATAST
	CLKDIV BWTR = 0x0F << 20 //+ CLKDIV
	DATLAT BWTR = 0x0F << 24 //+ DATLAT
	ACCMOD BWTR = 0x03 << 28 //+ ACCMOD
)

const (
	ADDSETn = 0
	ADDHLDn = 4
	DATASTn = 8
	CLKDIVn = 20
	DATLATn = 24
	ACCMODn = 28
)

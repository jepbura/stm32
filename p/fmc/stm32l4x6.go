// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package fmc provides access to the registers of the FMC peripheral.
//
// Instances:
//  FMC  FMC_BASE  AHB3  FMC,FPU*  Flexible memory controller
// Registers:
//  0x000 32  BCT{BCR,BTR}[4]  chip-select control/timing register
//  0x080 32  PCR              PC Card/NAND Flash control register 3
//  0x084 32  SR               FIFO status and interrupt register 3
//  0x088 32  PMEM             Common memory space timing register 3
//  0x08C 32  PATT             Attribute memory space timing register 3
//  0x094 32  ECCR             ECC result register 3
//  0x104 32  BWTR[4]          write timing registers
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package fmc

const (
	MBKEN     BCR = 0x01 << 0  //+ MBKEN
	MUXEN     BCR = 0x01 << 1  //+ MUXEN
	MTYP      BCR = 0x03 << 2  //+ MTYP
	MWID      BCR = 0x03 << 4  //+ MWID
	FACCEN    BCR = 0x01 << 6  //+ FACCEN
	BURSTEN   BCR = 0x01 << 8  //+ BURSTEN
	WAITPOL   BCR = 0x01 << 9  //+ WAITPOL
	WAITCFG   BCR = 0x01 << 11 //+ WAITCFG
	WREN      BCR = 0x01 << 12 //+ WREN
	WAITEN    BCR = 0x01 << 13 //+ WAITEN
	EXTMOD    BCR = 0x01 << 14 //+ EXTMOD
	ASYNCWAIT BCR = 0x01 << 15 //+ ASYNCWAIT
	CBURSTRW  BCR = 0x01 << 19 //+ CBURSTRW
	CCLKEN    BCR = 0x01 << 20 //+ CCLKEN
	WFDIS     BCR = 0x01 << 21 //+ Write FIFO Disable
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
	CCLKENn    = 20
	WFDISn     = 21
)

const (
	ADDSET  BTR = 0x0F << 0  //+ ADDSET
	ADDHLD  BTR = 0x0F << 4  //+ ADDHLD
	DATAST  BTR = 0xFF << 8  //+ DATAST
	BUSTURN BTR = 0x0F << 16 //+ BUSTURN
	CLKDIV  BTR = 0x0F << 20 //+ CLKDIV
	DATLAT  BTR = 0x0F << 24 //+ DATLAT
	ACCMOD  BTR = 0x03 << 28 //+ ACCMOD
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
	PWAITEN PCR = 0x01 << 1  //+ PWAITEN
	PBKEN   PCR = 0x01 << 2  //+ PBKEN
	PTYP    PCR = 0x01 << 3  //+ PTYP
	PWID    PCR = 0x03 << 4  //+ PWID
	ECCEN   PCR = 0x01 << 6  //+ ECCEN
	TCLR    PCR = 0x0F << 9  //+ TCLR
	TAR     PCR = 0x0F << 13 //+ TAR
	ECCPS   PCR = 0x07 << 17 //+ ECCPS
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
	IRS   SR = 0x01 << 0 //+ IRS
	ILS   SR = 0x01 << 1 //+ ILS
	IFS   SR = 0x01 << 2 //+ IFS
	IREN  SR = 0x01 << 3 //+ IREN
	ILEN  SR = 0x01 << 4 //+ ILEN
	IFEN  SR = 0x01 << 5 //+ IFEN
	FEMPT SR = 0x01 << 6 //+ FEMPT
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
	MEMSETx  PMEM = 0xFF << 0  //+ MEMSETx
	MEMWAITx PMEM = 0xFF << 8  //+ MEMWAITx
	MEMHOLDx PMEM = 0xFF << 16 //+ MEMHOLDx
	MEMHIZx  PMEM = 0xFF << 24 //+ MEMHIZx
)

const (
	MEMSETxn  = 0
	MEMWAITxn = 8
	MEMHOLDxn = 16
	MEMHIZxn  = 24
)

const (
	ATTSETx  PATT = 0xFF << 0  //+ ATTSETx
	ATTWAITx PATT = 0xFF << 8  //+ ATTWAITx
	ATTHOLDx PATT = 0xFF << 16 //+ ATTHOLDx
	ATTHIZx  PATT = 0xFF << 24 //+ ATTHIZx
)

const (
	ATTSETxn  = 0
	ATTWAITxn = 8
	ATTHOLDxn = 16
	ATTHIZxn  = 24
)

const (
	ECCx ECCR = 0xFFFFFFFF << 0 //+ ECCx
)

const (
	ECCxn = 0
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
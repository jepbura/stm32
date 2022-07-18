// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package dma provides access to the registers of the DMA peripheral.
//
// Instances:
//  DMA1  DMA1_BASE  AHB1  -  DMA controller
//  DMA2  DMA2_BASE  AHB1  -  DMA controller
// Registers:
//  0x000 32  ISR[2]                           interrupt status register
//  0x008 32  IFCR[2]                          interrupt flag clear register
//  0x010 32  S{CR,NDTR,PAR,M0AR,M1AR,FCR}[8]  stream configuration and controll registers
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package dma

const (
	FEIF0  ISR = 0x01 << 0  //+ Stream x FIFO error interrupt flag (x=3..0)
	DMEIF0 ISR = 0x01 << 2  //+ Stream x direct mode error interrupt flag (x=3..0)
	TEIF0  ISR = 0x01 << 3  //+ Stream x transfer error interrupt flag (x=3..0)
	HTIF0  ISR = 0x01 << 4  //+ Stream x half transfer interrupt flag (x=3..0)
	TCIF0  ISR = 0x01 << 5  //+ Stream x transfer complete interrupt flag (x = 3..0)
	FEIF1  ISR = 0x01 << 6  //+ Stream x FIFO error interrupt flag (x=3..0)
	DMEIF1 ISR = 0x01 << 8  //+ Stream x direct mode error interrupt flag (x=3..0)
	TEIF1  ISR = 0x01 << 9  //+ Stream x transfer error interrupt flag (x=3..0)
	HTIF1  ISR = 0x01 << 10 //+ Stream x half transfer interrupt flag (x=3..0)
	TCIF1  ISR = 0x01 << 11 //+ Stream x transfer complete interrupt flag (x = 3..0)
	FEIF2  ISR = 0x01 << 16 //+ Stream x FIFO error interrupt flag (x=3..0)
	DMEIF2 ISR = 0x01 << 18 //+ Stream x direct mode error interrupt flag (x=3..0)
	TEIF2  ISR = 0x01 << 19 //+ Stream x transfer error interrupt flag (x=3..0)
	HTIF2  ISR = 0x01 << 20 //+ Stream x half transfer interrupt flag (x=3..0)
	TCIF2  ISR = 0x01 << 21 //+ Stream x transfer complete interrupt flag (x = 3..0)
	FEIF3  ISR = 0x01 << 22 //+ Stream x FIFO error interrupt flag (x=3..0)
	DMEIF3 ISR = 0x01 << 24 //+ Stream x direct mode error interrupt flag (x=3..0)
	TEIF3  ISR = 0x01 << 25 //+ Stream x transfer error interrupt flag (x=3..0)
	HTIF3  ISR = 0x01 << 26 //+ Stream x half transfer interrupt flag (x=3..0)
	TCIF3  ISR = 0x01 << 27 //+ Stream x transfer complete interrupt flag (x = 3..0)
)

const (
	FEIF0n  = 0
	DMEIF0n = 2
	TEIF0n  = 3
	HTIF0n  = 4
	TCIF0n  = 5
	FEIF1n  = 6
	DMEIF1n = 8
	TEIF1n  = 9
	HTIF1n  = 10
	TCIF1n  = 11
	FEIF2n  = 16
	DMEIF2n = 18
	TEIF2n  = 19
	HTIF2n  = 20
	TCIF2n  = 21
	FEIF3n  = 22
	DMEIF3n = 24
	TEIF3n  = 25
	HTIF3n  = 26
	TCIF3n  = 27
)

const (
	CFEIF0  IFCR = 0x01 << 0  //+ Stream x clear FIFO error interrupt flag (x = 3..0)
	CDMEIF0 IFCR = 0x01 << 2  //+ Stream x clear direct mode error interrupt flag (x = 3..0)
	CTEIF0  IFCR = 0x01 << 3  //+ Stream x clear transfer error interrupt flag (x = 3..0)
	CHTIF0  IFCR = 0x01 << 4  //+ Stream x clear half transfer interrupt flag (x = 3..0)
	CTCIF0  IFCR = 0x01 << 5  //+ Stream x clear transfer complete interrupt flag (x = 3..0)
	CFEIF1  IFCR = 0x01 << 6  //+ Stream x clear FIFO error interrupt flag (x = 3..0)
	CDMEIF1 IFCR = 0x01 << 8  //+ Stream x clear direct mode error interrupt flag (x = 3..0)
	CTEIF1  IFCR = 0x01 << 9  //+ Stream x clear transfer error interrupt flag (x = 3..0)
	CHTIF1  IFCR = 0x01 << 10 //+ Stream x clear half transfer interrupt flag (x = 3..0)
	CTCIF1  IFCR = 0x01 << 11 //+ Stream x clear transfer complete interrupt flag (x = 3..0)
	CFEIF2  IFCR = 0x01 << 16 //+ Stream x clear FIFO error interrupt flag (x = 3..0)
	CDMEIF2 IFCR = 0x01 << 18 //+ Stream x clear direct mode error interrupt flag (x = 3..0)
	CTEIF2  IFCR = 0x01 << 19 //+ Stream x clear transfer error interrupt flag (x = 3..0)
	CHTIF2  IFCR = 0x01 << 20 //+ Stream x clear half transfer interrupt flag (x = 3..0)
	CTCIF2  IFCR = 0x01 << 21 //+ Stream x clear transfer complete interrupt flag (x = 3..0)
	CFEIF3  IFCR = 0x01 << 22 //+ Stream x clear FIFO error interrupt flag (x = 3..0)
	CDMEIF3 IFCR = 0x01 << 24 //+ Stream x clear direct mode error interrupt flag (x = 3..0)
	CTEIF3  IFCR = 0x01 << 25 //+ Stream x clear transfer error interrupt flag (x = 3..0)
	CHTIF3  IFCR = 0x01 << 26 //+ Stream x clear half transfer interrupt flag (x = 3..0)
	CTCIF3  IFCR = 0x01 << 27 //+ Stream x clear transfer complete interrupt flag (x = 3..0)
)

const (
	CFEIF0n  = 0
	CDMEIF0n = 2
	CTEIF0n  = 3
	CHTIF0n  = 4
	CTCIF0n  = 5
	CFEIF1n  = 6
	CDMEIF1n = 8
	CTEIF1n  = 9
	CHTIF1n  = 10
	CTCIF1n  = 11
	CFEIF2n  = 16
	CDMEIF2n = 18
	CTEIF2n  = 19
	CHTIF2n  = 20
	CTCIF2n  = 21
	CFEIF3n  = 22
	CDMEIF3n = 24
	CTEIF3n  = 25
	CHTIF3n  = 26
	CTCIF3n  = 27
)

const (
	EN     CR = 0x01 << 0  //+ Stream enable / flag stream ready when read low
	DMEIE  CR = 0x01 << 1  //+ Direct mode error interrupt enable
	TEIE   CR = 0x01 << 2  //+ Transfer error interrupt enable
	HTIE   CR = 0x01 << 3  //+ Half transfer interrupt enable
	TCIE   CR = 0x01 << 4  //+ Transfer complete interrupt enable
	PFCTRL CR = 0x01 << 5  //+ Peripheral flow controller
	DIR    CR = 0x03 << 6  //+ Data transfer direction
	CIRC   CR = 0x01 << 8  //+ Circular mode
	PINC   CR = 0x01 << 9  //+ Peripheral increment mode
	MINC   CR = 0x01 << 10 //+ Memory increment mode
	PSIZE  CR = 0x03 << 11 //+ Peripheral data size
	MSIZE  CR = 0x03 << 13 //+ Memory data size
	PINCOS CR = 0x01 << 15 //+ Peripheral increment offset size
	PL     CR = 0x03 << 16 //+ Priority level
	DBM    CR = 0x01 << 18 //+ Double buffer mode
	CT     CR = 0x01 << 19 //+ Current target (only in double buffer mode)
	PBURST CR = 0x03 << 21 //+ Peripheral burst transfer configuration
	MBURST CR = 0x03 << 23 //+ Memory burst transfer configuration
	CHSEL  CR = 0x07 << 25 //+ Channel selection
)

const (
	ENn     = 0
	DMEIEn  = 1
	TEIEn   = 2
	HTIEn   = 3
	TCIEn   = 4
	PFCTRLn = 5
	DIRn    = 6
	CIRCn   = 8
	PINCn   = 9
	MINCn   = 10
	PSIZEn  = 11
	MSIZEn  = 13
	PINCOSn = 15
	PLn     = 16
	DBMn    = 18
	CTn     = 19
	PBURSTn = 21
	MBURSTn = 23
	CHSELn  = 25
)

const (
	NDT NDTR = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	PA PAR = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	M0A M0AR = 0xFFFFFFFF << 0 //+ Memory 0 address
)

const (
	M0An = 0
)

const (
	M1A M1AR = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	M1An = 0
)

const (
	FTH   FCR = 0x03 << 0 //+ FIFO threshold selection
	DMDIS FCR = 0x01 << 2 //+ Direct mode disable
	FS    FCR = 0x07 << 3 //+ FIFO status
	FEIE  FCR = 0x01 << 7 //+ FIFO error interrupt enable
)

const (
	FTHn   = 0
	DMDISn = 2
	FSn    = 3
	FEIEn  = 7
)

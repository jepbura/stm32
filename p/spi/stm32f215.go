// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

// Package spi provides access to the registers of the SPI peripheral.
//
// Instances:
//  SPI1  SPI1_BASE  APB2  SPI1  Serial peripheral interface
//  SPI2  SPI2_BASE  APB1  SPI2  Serial peripheral interface
//  SPI3  SPI3_BASE  APB1  SPI3  Serial peripheral interface
// Registers:
//  0x000 32  CR1         control register 1
//  0x004 32  CR2         control register 2
//  0x008 32  SR          status register
//  0x00C 32  DR(uint32)  data register
//  0x010 32  CRCPR       CRC polynomial register
//  0x014 32  RXCRCR      RX CRC register
//  0x018 32  TXCRCR      TX CRC register
//  0x01C 32  I2SCFGR     I2S configuration register
//  0x020 32  I2SPR       I2S prescaler register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package spi

const (
	CPHA     CR1 = 0x01 << 0  //+ Clock phase
	CPOL     CR1 = 0x01 << 1  //+ Clock polarity
	MSTR     CR1 = 0x01 << 2  //+ Master selection
	BR       CR1 = 0x07 << 3  //+ Baud rate control
	SPE      CR1 = 0x01 << 6  //+ SPI enable
	LSBFIRST CR1 = 0x01 << 7  //+ Frame format
	SSI      CR1 = 0x01 << 8  //+ Internal slave select
	SSM      CR1 = 0x01 << 9  //+ Software slave management
	RXONLY   CR1 = 0x01 << 10 //+ Receive only
	DFF      CR1 = 0x01 << 11 //+ Data frame format
	CRCNEXT  CR1 = 0x01 << 12 //+ CRC transfer next
	CRCEN    CR1 = 0x01 << 13 //+ Hardware CRC calculation enable
	BIDIOE   CR1 = 0x01 << 14 //+ Output enable in bidirectional mode
	BIDIMODE CR1 = 0x01 << 15 //+ Bidirectional data mode enable
)

const (
	CPHAn     = 0
	CPOLn     = 1
	MSTRn     = 2
	BRn       = 3
	SPEn      = 6
	LSBFIRSTn = 7
	SSIn      = 8
	SSMn      = 9
	RXONLYn   = 10
	DFFn      = 11
	CRCNEXTn  = 12
	CRCENn    = 13
	BIDIOEn   = 14
	BIDIMODEn = 15
)

const (
	RXDMAEN CR2 = 0x01 << 0 //+ Rx buffer DMA enable
	TXDMAEN CR2 = 0x01 << 1 //+ Tx buffer DMA enable
	SSOE    CR2 = 0x01 << 2 //+ SS output enable
	FRF     CR2 = 0x01 << 4 //+ Frame format
	ERRIE   CR2 = 0x01 << 5 //+ Error interrupt enable
	RXNEIE  CR2 = 0x01 << 6 //+ RX buffer not empty interrupt enable
	TXEIE   CR2 = 0x01 << 7 //+ Tx buffer empty interrupt enable
)

const (
	RXDMAENn = 0
	TXDMAENn = 1
	SSOEn    = 2
	FRFn     = 4
	ERRIEn   = 5
	RXNEIEn  = 6
	TXEIEn   = 7
)

const (
	RXNE   SR = 0x01 << 0 //+ Receive buffer not empty
	TXE    SR = 0x01 << 1 //+ Transmit buffer empty
	CHSIDE SR = 0x01 << 2 //+ Channel side
	UDR    SR = 0x01 << 3 //+ Underrun flag
	CRCERR SR = 0x01 << 4 //+ CRC error flag
	MODF   SR = 0x01 << 5 //+ Mode fault
	OVR    SR = 0x01 << 6 //+ Overrun flag
	BSY    SR = 0x01 << 7 //+ Busy flag
	TIFRFE SR = 0x01 << 8 //+ TI frame format error
)

const (
	RXNEn   = 0
	TXEn    = 1
	CHSIDEn = 2
	UDRn    = 3
	CRCERRn = 4
	MODFn   = 5
	OVRn    = 6
	BSYn    = 7
	TIFRFEn = 8
)

const (
	CRCPOLY CRCPR = 0xFFFF << 0 //+ CRC polynomial register
)

const (
	CRCPOLYn = 0
)

const (
	RxCRC RXCRCR = 0xFFFF << 0 //+ Rx CRC register
)

const (
	RxCRCn = 0
)

const (
	TxCRC TXCRCR = 0xFFFF << 0 //+ Tx CRC register
)

const (
	TxCRCn = 0
)

const (
	CHLEN   I2SCFGR = 0x01 << 0  //+ Channel length (number of bits per audio channel)
	DATLEN  I2SCFGR = 0x03 << 1  //+ Data length to be transferred
	CKPOL   I2SCFGR = 0x01 << 3  //+ Steady state clock polarity
	I2SSTD  I2SCFGR = 0x03 << 4  //+ I2S standard selection
	PCMSYNC I2SCFGR = 0x01 << 7  //+ PCM frame synchronization
	I2SCFG  I2SCFGR = 0x03 << 8  //+ I2S configuration mode
	I2SE    I2SCFGR = 0x01 << 10 //+ I2S Enable
	I2SMOD  I2SCFGR = 0x01 << 11 //+ I2S mode selection
)

const (
	CHLENn   = 0
	DATLENn  = 1
	CKPOLn   = 3
	I2SSTDn  = 4
	PCMSYNCn = 7
	I2SCFGn  = 8
	I2SEn    = 10
	I2SMODn  = 11
)

const (
	I2SDIV I2SPR = 0xFF << 0 //+ I2S Linear prescaler
	ODD    I2SPR = 0x01 << 8 //+ Odd factor for the prescaler
	MCKOE  I2SPR = 0x01 << 9 //+ Master clock output enable
)

const (
	I2SDIVn = 0
	ODDn    = 8
	MCKOEn  = 9
)

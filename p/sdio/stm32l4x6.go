// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package sdio provides access to the registers of the SDMMC peripheral.
//
// Instances:
//  SDMMC1  SDMMC1_BASE  -  SDMMC1  Secure digital input/output interface
// Registers:
//  0x000 32  POWER    power control register
//  0x004 32  CLKCR    SDI clock control register
//  0x008 32  ARG      argument register
//  0x00C 32  CMD      command register
//  0x010 32  RESPCMD  command response register
//  0x014 32  RESP[4]  response registers
//  0x024 32  DTIMER   data timer register
//  0x028 32  DLEN     data length register
//  0x02C 32  DCTRL    data control register
//  0x030 32  DCOUNT   data counter register
//  0x034 32  STA      status register
//  0x038 32  ICR      interrupt clear register
//  0x03C 32  MASK     mask register
//  0x048 32  FIFOCNT  FIFO counter register
//  0x080 32  FIFO     data FIFO register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package sdio

const (
	PWRCTRL POWER = 0x03 << 0 //+ PWRCTRL
)

const (
	PWRCTRLn = 0
)

const (
	CLKDIV  CLKCR = 0xFF << 0  //+ Clock divide factor
	CLKEN   CLKCR = 0x01 << 8  //+ Clock enable bit
	PWRSAV  CLKCR = 0x01 << 9  //+ Power saving configuration bit
	BYPASS  CLKCR = 0x01 << 10 //+ Clock divider bypass enable bit
	WIDBUS  CLKCR = 0x03 << 11 //+ Wide bus mode enable bit
	NEGEDGE CLKCR = 0x01 << 13 //+ SDIO_CK dephasing selection bit
	HWFC_EN CLKCR = 0x01 << 14 //+ HW Flow Control enable
)

const (
	CLKDIVn  = 0
	CLKENn   = 8
	PWRSAVn  = 9
	BYPASSn  = 10
	WIDBUSn  = 11
	NEGEDGEn = 13
	HWFC_ENn = 14
)

const (
	CMDARG ARG = 0xFFFFFFFF << 0 //+ Command argument
)

const (
	CMDARGn = 0
)

const (
	CMDINDEX    CMD = 0x3F << 0  //+ Command index
	WAITRESP    CMD = 0x03 << 6  //+ Wait for response bits
	WAITINT     CMD = 0x01 << 8  //+ CPSM waits for interrupt request
	WAITPEND    CMD = 0x01 << 9  //+ CPSM Waits for ends of data transfer (CmdPend internal signal)
	CPSMEN      CMD = 0x01 << 10 //+ Command path state machine (CPSM) Enable bit
	SDIOSuspend CMD = 0x01 << 11 //+ SD I/O suspend command
	ENCMDcompl  CMD = 0x01 << 12 //+ Enable CMD completion
	nIEN        CMD = 0x01 << 13 //+ not Interrupt Enable
	CE_ATACMD   CMD = 0x01 << 14 //+ CE-ATA command
)

const (
	CMDINDEXn    = 0
	WAITRESPn    = 6
	WAITINTn     = 8
	WAITPENDn    = 9
	CPSMENn      = 10
	SDIOSuspendn = 11
	ENCMDcompln  = 12
	nIENn        = 13
	CE_ATACMDn   = 14
)

const (
	RESPCMD RESPCMD = 0x3F << 0 //+ Response command index
)

const (
	RESPCMDn = 0
)

const (
	CARDSTATUS1 RESP = 0xFFFFFFFF << 0 //+ see Table 132
)

const (
	CARDSTATUS1n = 0
)

const (
	DATATIME DTIMER = 0xFFFFFFFF << 0 //+ Data timeout period
)

const (
	DATATIMEn = 0
)

const (
	DATALENGTH DLEN = 0x1FFFFFF << 0 //+ Data length value
)

const (
	DATALENGTHn = 0
)

const (
	DTEN       DCTRL = 0x01 << 0  //+ DTEN
	DTDIR      DCTRL = 0x01 << 1  //+ Data transfer direction selection
	DTMODE     DCTRL = 0x01 << 2  //+ Data transfer mode selection 1: Stream or SDIO multibyte data transfer
	DMAEN      DCTRL = 0x01 << 3  //+ DMA enable bit
	DBLOCKSIZE DCTRL = 0x0F << 4  //+ Data block size
	RWSTART    DCTRL = 0x01 << 8  //+ Read wait start
	RWSTOP     DCTRL = 0x01 << 9  //+ Read wait stop
	RWMOD      DCTRL = 0x01 << 10 //+ Read wait mode
	SDIOEN     DCTRL = 0x01 << 11 //+ SD I/O enable functions
)

const (
	DTENn       = 0
	DTDIRn      = 1
	DTMODEn     = 2
	DMAENn      = 3
	DBLOCKSIZEn = 4
	RWSTARTn    = 8
	RWSTOPn     = 9
	RWMODn      = 10
	SDIOENn     = 11
)

const (
	DATACOUNT DCOUNT = 0x1FFFFFF << 0 //+ Data count value
)

const (
	DATACOUNTn = 0
)

const (
	CCRCFAIL STA = 0x01 << 0  //+ Command response received (CRC check failed)
	DCRCFAIL STA = 0x01 << 1  //+ Data block sent/received (CRC check failed)
	CTIMEOUT STA = 0x01 << 2  //+ Command response timeout
	DTIMEOUT STA = 0x01 << 3  //+ Data timeout
	TXUNDERR STA = 0x01 << 4  //+ Transmit FIFO underrun error
	RXOVERR  STA = 0x01 << 5  //+ Received FIFO overrun error
	CMDREND  STA = 0x01 << 6  //+ Command response received (CRC check passed)
	CMDSENT  STA = 0x01 << 7  //+ Command sent (no response required)
	DATAEND  STA = 0x01 << 8  //+ Data end (data counter, SDIDCOUNT, is zero)
	STBITERR STA = 0x01 << 9  //+ Start bit not detected on all data signals in wide bus mode
	DBCKEND  STA = 0x01 << 10 //+ Data block sent/received (CRC check passed)
	CMDACT   STA = 0x01 << 11 //+ Command transfer in progress
	TXACT    STA = 0x01 << 12 //+ Data transmit in progress
	RXACT    STA = 0x01 << 13 //+ Data receive in progress
	TXFIFOHE STA = 0x01 << 14 //+ Transmit FIFO half empty: at least 8 words can be written into the FIFO
	RXFIFOHF STA = 0x01 << 15 //+ Receive FIFO half full: there are at least 8 words in the FIFO
	TXFIFOF  STA = 0x01 << 16 //+ Transmit FIFO full
	RXFIFOF  STA = 0x01 << 17 //+ Receive FIFO full
	TXFIFOE  STA = 0x01 << 18 //+ Transmit FIFO empty
	RXFIFOE  STA = 0x01 << 19 //+ Receive FIFO empty
	TXDAVL   STA = 0x01 << 20 //+ Data available in transmit FIFO
	RXDAVL   STA = 0x01 << 21 //+ Data available in receive FIFO
	SDIOIT   STA = 0x01 << 22 //+ SDIO interrupt received
	CEATAEND STA = 0x01 << 23 //+ CE-ATA command completion signal received for CMD61
)

const (
	CCRCFAILn = 0
	DCRCFAILn = 1
	CTIMEOUTn = 2
	DTIMEOUTn = 3
	TXUNDERRn = 4
	RXOVERRn  = 5
	CMDRENDn  = 6
	CMDSENTn  = 7
	DATAENDn  = 8
	STBITERRn = 9
	DBCKENDn  = 10
	CMDACTn   = 11
	TXACTn    = 12
	RXACTn    = 13
	TXFIFOHEn = 14
	RXFIFOHFn = 15
	TXFIFOFn  = 16
	RXFIFOFn  = 17
	TXFIFOEn  = 18
	RXFIFOEn  = 19
	TXDAVLn   = 20
	RXDAVLn   = 21
	SDIOITn   = 22
	CEATAENDn = 23
)

const (
	CCRCFAILC ICR = 0x01 << 0  //+ CCRCFAIL flag clear bit
	DCRCFAILC ICR = 0x01 << 1  //+ DCRCFAIL flag clear bit
	CTIMEOUTC ICR = 0x01 << 2  //+ CTIMEOUT flag clear bit
	DTIMEOUTC ICR = 0x01 << 3  //+ DTIMEOUT flag clear bit
	TXUNDERRC ICR = 0x01 << 4  //+ TXUNDERR flag clear bit
	RXOVERRC  ICR = 0x01 << 5  //+ RXOVERR flag clear bit
	CMDRENDC  ICR = 0x01 << 6  //+ CMDREND flag clear bit
	CMDSENTC  ICR = 0x01 << 7  //+ CMDSENT flag clear bit
	DATAENDC  ICR = 0x01 << 8  //+ DATAEND flag clear bit
	STBITERRC ICR = 0x01 << 9  //+ STBITERR flag clear bit
	DBCKENDC  ICR = 0x01 << 10 //+ DBCKEND flag clear bit
	SDIOITC   ICR = 0x01 << 22 //+ SDIOIT flag clear bit
	CEATAENDC ICR = 0x01 << 23 //+ CEATAEND flag clear bit
)

const (
	CCRCFAILCn = 0
	DCRCFAILCn = 1
	CTIMEOUTCn = 2
	DTIMEOUTCn = 3
	TXUNDERRCn = 4
	RXOVERRCn  = 5
	CMDRENDCn  = 6
	CMDSENTCn  = 7
	DATAENDCn  = 8
	STBITERRCn = 9
	DBCKENDCn  = 10
	SDIOITCn   = 22
	CEATAENDCn = 23
)

const (
	CCRCFAILIE MASK = 0x01 << 0  //+ Command CRC fail interrupt enable
	DCRCFAILIE MASK = 0x01 << 1  //+ Data CRC fail interrupt enable
	CTIMEOUTIE MASK = 0x01 << 2  //+ Command timeout interrupt enable
	DTIMEOUTIE MASK = 0x01 << 3  //+ Data timeout interrupt enable
	TXUNDERRIE MASK = 0x01 << 4  //+ Tx FIFO underrun error interrupt enable
	RXOVERRIE  MASK = 0x01 << 5  //+ Rx FIFO overrun error interrupt enable
	CMDRENDIE  MASK = 0x01 << 6  //+ Command response received interrupt enable
	CMDSENTIE  MASK = 0x01 << 7  //+ Command sent interrupt enable
	DATAENDIE  MASK = 0x01 << 8  //+ Data end interrupt enable
	STBITERRIE MASK = 0x01 << 9  //+ Start bit error interrupt enable
	DBCKENDIE  MASK = 0x01 << 10 //+ Data block end interrupt enable
	CMDACTIE   MASK = 0x01 << 11 //+ Command acting interrupt enable
	TXACTIE    MASK = 0x01 << 12 //+ Data transmit acting interrupt enable
	RXACTIE    MASK = 0x01 << 13 //+ Data receive acting interrupt enable
	TXFIFOHEIE MASK = 0x01 << 14 //+ Tx FIFO half empty interrupt enable
	RXFIFOHFIE MASK = 0x01 << 15 //+ Rx FIFO half full interrupt enable
	TXFIFOFIE  MASK = 0x01 << 16 //+ Tx FIFO full interrupt enable
	RXFIFOFIE  MASK = 0x01 << 17 //+ Rx FIFO full interrupt enable
	TXFIFOEIE  MASK = 0x01 << 18 //+ Tx FIFO empty interrupt enable
	RXFIFOEIE  MASK = 0x01 << 19 //+ Rx FIFO empty interrupt enable
	TXDAVLIE   MASK = 0x01 << 20 //+ Data available in Tx FIFO interrupt enable
	RXDAVLIE   MASK = 0x01 << 21 //+ Data available in Rx FIFO interrupt enable
	SDIOITIE   MASK = 0x01 << 22 //+ SDIO mode interrupt received interrupt enable
	CEATAENDIE MASK = 0x01 << 23 //+ CE-ATA command completion signal received interrupt enable
)

const (
	CCRCFAILIEn = 0
	DCRCFAILIEn = 1
	CTIMEOUTIEn = 2
	DTIMEOUTIEn = 3
	TXUNDERRIEn = 4
	RXOVERRIEn  = 5
	CMDRENDIEn  = 6
	CMDSENTIEn  = 7
	DATAENDIEn  = 8
	STBITERRIEn = 9
	DBCKENDIEn  = 10
	CMDACTIEn   = 11
	TXACTIEn    = 12
	RXACTIEn    = 13
	TXFIFOHEIEn = 14
	RXFIFOHFIEn = 15
	TXFIFOFIEn  = 16
	RXFIFOFIEn  = 17
	TXFIFOEIEn  = 18
	RXFIFOEIEn  = 19
	TXDAVLIEn   = 20
	RXDAVLIEn   = 21
	SDIOITIEn   = 22
	CEATAENDIEn = 23
)

const (
	FIFOCOUNT FIFOCNT = 0xFFFFFF << 0 //+ Remaining number of words to be written to or read from the FIFO
)

const (
	FIFOCOUNTn = 0
)

const (
	FIFOData FIFO = 0xFFFFFFFF << 0 //+ Receive and transmit FIFO data
)

const (
	FIFODatan = 0
)

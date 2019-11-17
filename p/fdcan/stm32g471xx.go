// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package fdcan provides access to the registers of the FDCAN peripheral.
//
// Instances:
//  FDCAN   FDCAN_BASE   APB1  FDCAN2_intr0+,FDCAN2_intr1+  FDCAN
//  FDCAN1  FDCAN1_BASE  -     FDCAN2_intr0+,FDCAN2_intr1+  FDCAN
//  FDCAN2  FDCAN2_BASE  -     FDCAN2_intr0+,FDCAN2_intr1+  FDCAN
// Registers:
//  0x000 32  CREL    FDCAN Core Release Register
//  0x004 32  ENDN    FDCAN Core Release Register
//  0x00C 32  DBTP    This register is only writable if bits CCCR.CCE and CCCR.INIT are set. The CAN bit time may be programed in the range of 4 to 25 time quanta. The CAN time quantum may be programmed in the range of 1 to 1024 FDCAN clock periods. tq = (DBRP + 1) FDCAN clock period. DTSEG1 is the sum of Prop_Seg and Phase_Seg1. DTSEG2 is Phase_Seg2. Therefore the length of the bit time is (programmed values) [DTSEG1 + DTSEG2 + 3] tq or (functional values) [Sync_Seg + Prop_Seg + Phase_Seg1 + Phase_Seg2] tq. The Information Processing Time (IPT) is zero, meaning the data for the next bit is available at the first clock edge after the sample point.
//  0x010 32  TEST    Write access to the Test Register has to be enabled by setting bit CCCR[TEST] to 1 . All Test Register functions are set to their reset values when bit CCCR[TEST] is reset. Loop Back mode and software control of Tx pin FDCANx_TX are hardware test modes. Programming TX differently from 00 may disturb the message transfer on the CAN bus.
//  0x014 32  RWD     The RAM Watchdog monitors the READY output of the Message RAM. A Message RAM access starts the Message RAM Watchdog Counter with the value configured by the RWD[WDC] bits. The counter is reloaded with RWD[WDC] bits when the Message RAM signals successful completion by activating its READY output. In case there is no response from the Message RAM until the counter has counted down to 0, the counter stops and interrupt flag IR[WDI] bit is set. The RAM Watchdog Counter is clocked by the fdcan_pclk clock.
//  0x018 32  CCCR    For details about setting and resetting of single bits see Software initialization.
//  0x01C 32  NBTP    FDCAN_NBTP
//  0x020 32  TSCC    FDCAN Timestamp Counter Configuration Register
//  0x024 32  TSCV    FDCAN Timestamp Counter Value Register
//  0x028 32  TOCC    FDCAN Timeout Counter Configuration Register
//  0x02C 32  TOCV    FDCAN Timeout Counter Value Register
//  0x040 32  ECR     FDCAN Error Counter Register
//  0x044 32  PSR     FDCAN Protocol Status Register
//  0x048 32  TDCR    FDCAN Transmitter Delay Compensation Register
//  0x050 32  IR      The flags are set when one of the listed conditions is detected (edge-sensitive). The flags remain set until the Host clears them. A flag is cleared by writing a 1 to the corresponding bit position. Writing a 0 has no effect. A hard reset will clear the register. The configuration of IE controls whether an interrupt is generated. The configuration of ILS controls on which interrupt line an interrupt is signaled.
//  0x054 32  IE      The settings in the Interrupt Enable register determine which status changes in the Interrupt Register will be signaled on an interrupt line.
//  0x058 32  ILS     The Interrupt Line Select register assigns an interrupt generated by a specific interrupt flag from the Interrupt Register to one of the two module interrupt lines. For interrupt generation the respective interrupt line has to be enabled via ILE[EINT0] and ILE[EINT1].
//  0x05C 32  ILE     Each of the two interrupt lines to the CPU can be enabled/disabled separately by programming bits EINT0 and EINT1.
//  0x080 32  RXGFC   Global settings for Message ID filtering. The Global Filter Configuration controls the filter path for standard and extended messages as described in Figure706: Standard Message ID filter path and Figure707: Extended Message ID filter path.
//  0x084 32  XIDAM   FDCAN Extended ID and Mask Register
//  0x088 32  HPMS    This register is updated every time a Message ID filter element configured to generate a priority event match. This can be used to monitor the status of incoming high priority messages and to enable fast access to these messages.
//  0x090 32  RXF0S   FDCAN Rx FIFO 0 Status Register
//  0x094 32  RXF0A   CAN Rx FIFO 0 Acknowledge Register
//  0x098 32  RXF1S   FDCAN Rx FIFO 1 Status Register
//  0x09C 32  RXF1A   FDCAN Rx FIFO 1 Acknowledge Register
//  0x0C0 32  TXBC    FDCAN Tx Buffer Configuration Register
//  0x0C4 32  TXFQS   The Tx FIFO/Queue status is related to the pending Tx requests listed in register TXBRP. Therefore the effect of Add/Cancellation requests may be delayed due to a running Tx scan (TXBRP not yet updated).
//  0x0C8 32  TXBRP   FDCAN Tx Buffer Request Pending Register
//  0x0CC 32  TXBAR   FDCAN Tx Buffer Add Request Register
//  0x0D0 32  TXBCR   FDCAN Tx Buffer Cancellation Request Register
//  0x0D4 32  TXBTO   FDCAN Tx Buffer Transmission Occurred Register
//  0x0D8 32  TXBCF   FDCAN Tx Buffer Cancellation Finished Register
//  0x0DC 32  TXBTIE  FDCAN Tx Buffer Transmission Interrupt Enable Register
//  0x0E0 32  TXBCIE  FDCAN Tx Buffer Cancellation Finished Interrupt Enable Register
//  0x0E4 32  TXEFS   FDCAN Tx Event FIFO Status Register
//  0x0E8 32  TXEFA   FDCAN Tx Event FIFO Acknowledge Register
//  0x100 32  CKDIV   FDCAN CFG clock divider register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package fdcan

const (
	DAY     CREL = 0xFF << 0  //+ DAY
	MON     CREL = 0xFF << 8  //+ MON
	YEAR    CREL = 0x0F << 16 //+ YEAR
	SUBSTEP CREL = 0x0F << 20 //+ SUBSTEP
	STEP    CREL = 0x0F << 24 //+ STEP
	REL     CREL = 0x0F << 28 //+ REL
)

const (
	DAYn     = 0
	MONn     = 8
	YEARn    = 16
	SUBSTEPn = 20
	STEPn    = 24
	RELn     = 28
)

const (
	ETV ENDN = 0xFFFFFFFF << 0 //+ ETV
)

const (
	ETVn = 0
)

const (
	DSJW   DBTP = 0x0F << 0  //+ DSJW
	DTSEG2 DBTP = 0x0F << 4  //+ DTSEG2
	DTSEG1 DBTP = 0x1F << 8  //+ DTSEG1
	DBRP   DBTP = 0x1F << 16 //+ DBRP
	TDC    DBTP = 0x01 << 23 //+ TDC
)

const (
	DSJWn   = 0
	DTSEG2n = 4
	DTSEG1n = 8
	DBRPn   = 16
	TDCn    = 23
)

const (
	LBCK TEST = 0x01 << 4 //+ LBCK
	TX   TEST = 0x03 << 5 //+ TX
	RX   TEST = 0x01 << 7 //+ RX
)

const (
	LBCKn = 4
	TXn   = 5
	RXn   = 7
)

const (
	WDC RWD = 0xFF << 0 //+ WDC
	WDV RWD = 0xFF << 8 //+ WDV
)

const (
	WDCn = 0
	WDVn = 8
)

const (
	INIT CCCR = 0x01 << 0  //+ INIT
	CCE  CCCR = 0x01 << 1  //+ CCE
	ASM  CCCR = 0x01 << 2  //+ ASM
	CSA  CCCR = 0x01 << 3  //+ CSA
	CSR  CCCR = 0x01 << 4  //+ CSR
	MON  CCCR = 0x01 << 5  //+ MON
	DAR  CCCR = 0x01 << 6  //+ DAR
	TEST CCCR = 0x01 << 7  //+ TEST
	FDOE CCCR = 0x01 << 8  //+ FDOE
	BRSE CCCR = 0x01 << 9  //+ BRSE
	PXHD CCCR = 0x01 << 12 //+ PXHD
	EFBI CCCR = 0x01 << 13 //+ EFBI
	TXP  CCCR = 0x01 << 14 //+ TXP
	NISO CCCR = 0x01 << 15 //+ NISO
)

const (
	INITn = 0
	CCEn  = 1
	ASMn  = 2
	CSAn  = 3
	CSRn  = 4
	MONn  = 5
	DARn  = 6
	TESTn = 7
	FDOEn = 8
	BRSEn = 9
	PXHDn = 12
	EFBIn = 13
	TXPn  = 14
	NISOn = 15
)

const (
	TSEG2  NBTP = 0x7F << 0   //+ TSEG2
	NTSEG1 NBTP = 0xFF << 8   //+ NTSEG1
	NBRP   NBTP = 0x1FF << 16 //+ NBRP
	NSJW   NBTP = 0x7F << 25  //+ NSJW
)

const (
	TSEG2n  = 0
	NTSEG1n = 8
	NBRPn   = 16
	NSJWn   = 25
)

const (
	TSS TSCC = 0x03 << 0  //+ TSS
	TCP TSCC = 0x0F << 16 //+ TCP
)

const (
	TSSn = 0
	TCPn = 16
)

const (
	TSC TSCV = 0xFFFF << 0 //+ TSC
)

const (
	TSCn = 0
)

const (
	ETOC TOCC = 0x01 << 0    //+ ETOC
	TOS  TOCC = 0x03 << 1    //+ TOS
	TOP  TOCC = 0xFFFF << 16 //+ TOP
)

const (
	ETOCn = 0
	TOSn  = 1
	TOPn  = 16
)

const (
	TOC TOCV = 0xFFFF << 0 //+ TOC
)

const (
	TOCn = 0
)

const (
	TEC  ECR = 0xFF << 0  //+ TEC
	TREC ECR = 0x7F << 8  //+ TREC
	RP   ECR = 0x01 << 15 //+ RP
	CEL  ECR = 0xFF << 16 //+ CEL
)

const (
	TECn  = 0
	TRECn = 8
	RPn   = 15
	CELn  = 16
)

const (
	LEC  PSR = 0x07 << 0  //+ LEC
	ACT  PSR = 0x03 << 3  //+ ACT
	EP   PSR = 0x01 << 5  //+ EP
	EW   PSR = 0x01 << 6  //+ EW
	BO   PSR = 0x01 << 7  //+ BO
	DLEC PSR = 0x07 << 8  //+ DLEC
	RESI PSR = 0x01 << 11 //+ RESI
	RBRS PSR = 0x01 << 12 //+ RBRS
	REDL PSR = 0x01 << 13 //+ REDL
	PXE  PSR = 0x01 << 14 //+ PXE
	TDCV PSR = 0x7F << 16 //+ TDCV
)

const (
	LECn  = 0
	ACTn  = 3
	EPn   = 5
	EWn   = 6
	BOn   = 7
	DLECn = 8
	RESIn = 11
	RBRSn = 12
	REDLn = 13
	PXEn  = 14
	TDCVn = 16
)

const (
	TDCF TDCR = 0x7F << 0 //+ TDCF
	TDCO TDCR = 0x7F << 8 //+ TDCO
)

const (
	TDCFn = 0
	TDCOn = 8
)

const (
	RF0N IR = 0x01 << 0  //+ RF0N
	RF0W IR = 0x01 << 1  //+ RF0W
	RF0F IR = 0x01 << 2  //+ RF0F
	RF0L IR = 0x01 << 3  //+ RF0L
	RF1N IR = 0x01 << 4  //+ RF1N
	RF1W IR = 0x01 << 5  //+ RF1W
	RF1F IR = 0x01 << 6  //+ RF1F
	RF1L IR = 0x01 << 7  //+ RF1L
	HPM  IR = 0x01 << 8  //+ HPM
	TC   IR = 0x01 << 9  //+ TC
	TCF  IR = 0x01 << 10 //+ TCF
	TFE  IR = 0x01 << 11 //+ TFE
	TEFN IR = 0x01 << 12 //+ TEFN
	TEFW IR = 0x01 << 13 //+ TEFW
	TEFF IR = 0x01 << 14 //+ TEFF
	TEFL IR = 0x01 << 15 //+ TEFL
	TSW  IR = 0x01 << 16 //+ TSW
	MRAF IR = 0x01 << 17 //+ MRAF
	TOO  IR = 0x01 << 18 //+ TOO
	DRX  IR = 0x01 << 19 //+ DRX
	ELO  IR = 0x01 << 22 //+ ELO
	EP   IR = 0x01 << 23 //+ EP
	EW   IR = 0x01 << 24 //+ EW
	BO   IR = 0x01 << 25 //+ BO
	WDI  IR = 0x01 << 26 //+ WDI
	PEA  IR = 0x01 << 27 //+ PEA
	PED  IR = 0x01 << 28 //+ PED
	ARA  IR = 0x01 << 29 //+ ARA
)

const (
	RF0Nn = 0
	RF0Wn = 1
	RF0Fn = 2
	RF0Ln = 3
	RF1Nn = 4
	RF1Wn = 5
	RF1Fn = 6
	RF1Ln = 7
	HPMn  = 8
	TCn   = 9
	TCFn  = 10
	TFEn  = 11
	TEFNn = 12
	TEFWn = 13
	TEFFn = 14
	TEFLn = 15
	TSWn  = 16
	MRAFn = 17
	TOOn  = 18
	DRXn  = 19
	ELOn  = 22
	EPn   = 23
	EWn   = 24
	BOn   = 25
	WDIn  = 26
	PEAn  = 27
	PEDn  = 28
	ARAn  = 29
)

const (
	RF0NE IE = 0x01 << 0  //+ RF0NE
	RF0WE IE = 0x01 << 1  //+ RF0WE
	RF0FE IE = 0x01 << 2  //+ RF0FE
	RF0LE IE = 0x01 << 3  //+ RF0LE
	RF1NE IE = 0x01 << 4  //+ RF1NE
	RF1WE IE = 0x01 << 5  //+ RF1WE
	RF1FE IE = 0x01 << 6  //+ RF1FE
	RF1LE IE = 0x01 << 7  //+ RF1LE
	HPME  IE = 0x01 << 8  //+ HPME
	TCE   IE = 0x01 << 9  //+ TCE
	TCFE  IE = 0x01 << 10 //+ TCFE
	TFEE  IE = 0x01 << 11 //+ TFEE
	TEFNE IE = 0x01 << 12 //+ TEFNE
	TEFWE IE = 0x01 << 13 //+ TEFWE
	TEFFE IE = 0x01 << 14 //+ TEFFE
	TEFLE IE = 0x01 << 15 //+ TEFLE
	TSWE  IE = 0x01 << 16 //+ TSWE
	MRAFE IE = 0x01 << 17 //+ MRAFE
	TOOE  IE = 0x01 << 18 //+ TOOE
	DRX   IE = 0x01 << 19 //+ DRX
	BECE  IE = 0x01 << 20 //+ BECE
	BEUE  IE = 0x01 << 21 //+ BEUE
	ELOE  IE = 0x01 << 22 //+ ELOE
	EPE   IE = 0x01 << 23 //+ EPE
	EWE   IE = 0x01 << 24 //+ EWE
	BOE   IE = 0x01 << 25 //+ BOE
	WDIE  IE = 0x01 << 26 //+ WDIE
	PEAE  IE = 0x01 << 27 //+ PEAE
	PEDE  IE = 0x01 << 28 //+ PEDE
	ARAE  IE = 0x01 << 29 //+ ARAE
)

const (
	RF0NEn = 0
	RF0WEn = 1
	RF0FEn = 2
	RF0LEn = 3
	RF1NEn = 4
	RF1WEn = 5
	RF1FEn = 6
	RF1LEn = 7
	HPMEn  = 8
	TCEn   = 9
	TCFEn  = 10
	TFEEn  = 11
	TEFNEn = 12
	TEFWEn = 13
	TEFFEn = 14
	TEFLEn = 15
	TSWEn  = 16
	MRAFEn = 17
	TOOEn  = 18
	DRXn   = 19
	BECEn  = 20
	BEUEn  = 21
	ELOEn  = 22
	EPEn   = 23
	EWEn   = 24
	BOEn   = 25
	WDIEn  = 26
	PEAEn  = 27
	PEDEn  = 28
	ARAEn  = 29
)

const (
	RF0NL ILS = 0x01 << 0  //+ RF0NL
	RF0WL ILS = 0x01 << 1  //+ RF0WL
	RF0FL ILS = 0x01 << 2  //+ RF0FL
	RF0LL ILS = 0x01 << 3  //+ RF0LL
	RF1NL ILS = 0x01 << 4  //+ RF1NL
	RF1WL ILS = 0x01 << 5  //+ RF1WL
	RF1FL ILS = 0x01 << 6  //+ RF1FL
	RF1LL ILS = 0x01 << 7  //+ RF1LL
	HPML  ILS = 0x01 << 8  //+ HPML
	TCL   ILS = 0x01 << 9  //+ TCL
	TCFL  ILS = 0x01 << 10 //+ TCFL
	TFEL  ILS = 0x01 << 11 //+ TFEL
	TEFNL ILS = 0x01 << 12 //+ TEFNL
	TEFWL ILS = 0x01 << 13 //+ TEFWL
	TEFFL ILS = 0x01 << 14 //+ TEFFL
	TEFLL ILS = 0x01 << 15 //+ TEFLL
	TSWL  ILS = 0x01 << 16 //+ TSWL
	MRAFL ILS = 0x01 << 17 //+ MRAFL
	TOOL  ILS = 0x01 << 18 //+ TOOL
	DRXL  ILS = 0x01 << 19 //+ DRXL
	BECL  ILS = 0x01 << 20 //+ BECL
	BEUL  ILS = 0x01 << 21 //+ BEUL
	ELOL  ILS = 0x01 << 22 //+ ELOL
	EPL   ILS = 0x01 << 23 //+ EPL
	EWL   ILS = 0x01 << 24 //+ EWL
	BOL   ILS = 0x01 << 25 //+ BOL
	WDIL  ILS = 0x01 << 26 //+ WDIL
	PEAL  ILS = 0x01 << 27 //+ PEAL
	PEDL  ILS = 0x01 << 28 //+ PEDL
	ARAL  ILS = 0x01 << 29 //+ ARAL
)

const (
	RF0NLn = 0
	RF0WLn = 1
	RF0FLn = 2
	RF0LLn = 3
	RF1NLn = 4
	RF1WLn = 5
	RF1FLn = 6
	RF1LLn = 7
	HPMLn  = 8
	TCLn   = 9
	TCFLn  = 10
	TFELn  = 11
	TEFNLn = 12
	TEFWLn = 13
	TEFFLn = 14
	TEFLLn = 15
	TSWLn  = 16
	MRAFLn = 17
	TOOLn  = 18
	DRXLn  = 19
	BECLn  = 20
	BEULn  = 21
	ELOLn  = 22
	EPLn   = 23
	EWLn   = 24
	BOLn   = 25
	WDILn  = 26
	PEALn  = 27
	PEDLn  = 28
	ARALn  = 29
)

const (
	EINT0 ILE = 0x01 << 0 //+ EINT0
	EINT1 ILE = 0x01 << 1 //+ EINT1
)

const (
	EINT0n = 0
	EINT1n = 1
)

const (
	RRFE RXGFC = 0x01 << 0 //+ RRFE
	RRFS RXGFC = 0x01 << 1 //+ RRFS
	ANFE RXGFC = 0x03 << 2 //+ ANFE
	ANFS RXGFC = 0x03 << 4 //+ ANFS
)

const (
	RRFEn = 0
	RRFSn = 1
	ANFEn = 2
	ANFSn = 4
)

const (
	EIDM XIDAM = 0x1FFFFFFF << 0 //+ EIDM
)

const (
	EIDMn = 0
)

const (
	BIDX HPMS = 0x3F << 0  //+ BIDX
	MSI  HPMS = 0x03 << 6  //+ MSI
	FIDX HPMS = 0x7F << 8  //+ FIDX
	FLST HPMS = 0x01 << 15 //+ FLST
)

const (
	BIDXn = 0
	MSIn  = 6
	FIDXn = 8
	FLSTn = 15
)

const (
	F0FL RXF0S = 0x7F << 0  //+ F0FL
	F0GI RXF0S = 0x3F << 8  //+ F0GI
	F0PI RXF0S = 0x3F << 16 //+ F0PI
	F0F  RXF0S = 0x01 << 24 //+ F0F
	RF0L RXF0S = 0x01 << 25 //+ RF0L
)

const (
	F0FLn = 0
	F0GIn = 8
	F0PIn = 16
	F0Fn  = 24
	RF0Ln = 25
)

const (
	F0AI RXF0A = 0x3F << 0 //+ F0AI
)

const (
	F0AIn = 0
)

const (
	F1FL RXF1S = 0x7F << 0  //+ F1FL
	F1GI RXF1S = 0x3F << 8  //+ F1GI
	F1PI RXF1S = 0x3F << 16 //+ F1PI
	F1F  RXF1S = 0x01 << 24 //+ F1F
	RF1L RXF1S = 0x01 << 25 //+ RF1L
	DMS  RXF1S = 0x03 << 30 //+ DMS
)

const (
	F1FLn = 0
	F1GIn = 8
	F1PIn = 16
	F1Fn  = 24
	RF1Ln = 25
	DMSn  = 30
)

const (
	F1AI RXF1A = 0x3F << 0 //+ F1AI
)

const (
	F1AIn = 0
)

const (
	TBSA TXBC = 0x3FFF << 2 //+ TBSA
	NDTB TXBC = 0x3F << 16  //+ NDTB
	TFQS TXBC = 0x3F << 24  //+ TFQS
	TFQM TXBC = 0x01 << 30  //+ TFQM
)

const (
	TBSAn = 2
	NDTBn = 16
	TFQSn = 24
	TFQMn = 30
)

const (
	TFFL  TXFQS = 0x3F << 0  //+ TFFL
	TFGI  TXFQS = 0x1F << 8  //+ TFGI
	TFQPI TXFQS = 0x1F << 16 //+ TFQPI
	TFQF  TXFQS = 0x01 << 21 //+ TFQF
)

const (
	TFFLn  = 0
	TFGIn  = 8
	TFQPIn = 16
	TFQFn  = 21
)

const (
	TRP TXBRP = 0xFFFFFFFF << 0 //+ TRP
)

const (
	TRPn = 0
)

const (
	AR TXBAR = 0xFFFFFFFF << 0 //+ AR
)

const (
	ARn = 0
)

const (
	CR TXBCR = 0xFFFFFFFF << 0 //+ CR
)

const (
	CRn = 0
)

const (
	TO TXBTO = 0xFFFFFFFF << 0 //+ TO
)

const (
	TOn = 0
)

const (
	CF TXBCF = 0xFFFFFFFF << 0 //+ CF
)

const (
	CFn = 0
)

const (
	TIE TXBTIE = 0xFFFFFFFF << 0 //+ TIE
)

const (
	TIEn = 0
)

const (
	CFIE TXBCIE = 0xFFFFFFFF << 0 //+ CFIE
)

const (
	CFIEn = 0
)

const (
	EFFL TXEFS = 0x3F << 0  //+ EFFL
	EFGI TXEFS = 0x1F << 8  //+ EFGI
	EFPI TXEFS = 0x1F << 16 //+ EFPI
	EFF  TXEFS = 0x01 << 24 //+ EFF
	TEFL TXEFS = 0x01 << 25 //+ TEFL
)

const (
	EFFLn = 0
	EFGIn = 8
	EFPIn = 16
	EFFn  = 24
	TEFLn = 25
)

const (
	EFAI TXEFA = 0x1F << 0 //+ EFAI
)

const (
	EFAIn = 0
)

const (
	PDIV CKDIV = 0x0F << 0 //+ input clock divider. the APB clock could be divided prior to be used by the CAN sub
)

const (
	PDIVn = 0
)
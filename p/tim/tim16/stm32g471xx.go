// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package tim16 provides access to the registers of the TIM peripheral.
//
// Instances:
//  TIM16  TIM16_BASE  APB2  -  General purpose timers
//  TIM17  TIM17_BASE  APB2  -  General purpose timers
// Registers:
//  0x000 32  CR1           control register 1
//  0x004 32  CR2           control register 2
//  0x00C 32  DIER          DMA/Interrupt enable register
//  0x010 32  SR            status register
//  0x014 32  EGR           event generation register
//  0x018 32  CCMR1_Output  capture/compare mode register (output mode)
//  0x018 32  CCMR1_Input   capture/compare mode register 1 (input mode)
//  0x020 32  CCER          capture/compare enable register
//  0x024 32  CNT           counter
//  0x028 32  PSC           prescaler
//  0x02C 32  ARR           auto-reload register
//  0x030 32  RCR           repetition counter register
//  0x034 32  CCR1          capture/compare register 1
//  0x044 32  BDTR          break and dead-time register
//  0x054 32  DTR2          timer Deadtime Register 2
//  0x05C 32  TISEL         TIM timer input selection register
//  0x060 32  AF1           TIM alternate function option register 1
//  0x064 32  AF2           TIM alternate function option register 2
//  0x3DC 32  DCR           DMA control register
//  0x3E0 32  DMAR          DMA address for full transfer
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package tim16

const (
	CEN      CR1 = 0x01 << 0  //+ Counter enable
	UDIS     CR1 = 0x01 << 1  //+ Update disable
	URS      CR1 = 0x01 << 2  //+ Update request source
	OPM      CR1 = 0x01 << 3  //+ One-pulse mode
	ARPE     CR1 = 0x01 << 7  //+ Auto-reload preload enable
	CKD      CR1 = 0x03 << 8  //+ Clock division
	UIFREMAP CR1 = 0x01 << 11 //+ UIF status bit remapping
	DITHEN   CR1 = 0x01 << 12 //+ Dithering Enable
)

const (
	CENn      = 0
	UDISn     = 1
	URSn      = 2
	OPMn      = 3
	ARPEn     = 7
	CKDn      = 8
	UIFREMAPn = 11
	DITHENn   = 12
)

const (
	CCPC  CR2 = 0x01 << 0 //+ Capture/compare preloaded control
	CCUS  CR2 = 0x01 << 2 //+ Capture/compare control update selection
	CCDS  CR2 = 0x01 << 3 //+ Capture/compare DMA selection
	OIS1  CR2 = 0x01 << 8 //+ Output Idle state 1
	OIS1N CR2 = 0x01 << 9 //+ Output Idle state 1
)

const (
	CCPCn  = 0
	CCUSn  = 2
	CCDSn  = 3
	OIS1n  = 8
	OIS1Nn = 9
)

const (
	UIE   DIER = 0x01 << 0  //+ Update interrupt enable
	CC1IE DIER = 0x01 << 1  //+ Capture/Compare 1 interrupt enable
	COMIE DIER = 0x01 << 5  //+ COM interrupt enable
	BIE   DIER = 0x01 << 7  //+ Break interrupt enable
	UDE   DIER = 0x01 << 8  //+ Update DMA request enable
	CC1DE DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable
	COMDE DIER = 0x01 << 13 //+ COM DMA request enable
)

const (
	UIEn   = 0
	CC1IEn = 1
	COMIEn = 5
	BIEn   = 7
	UDEn   = 8
	CC1DEn = 9
	COMDEn = 13
)

const (
	UIF   SR = 0x01 << 0 //+ Update interrupt flag
	CC1IF SR = 0x01 << 1 //+ Capture/compare 1 interrupt flag
	COMIF SR = 0x01 << 5 //+ COM interrupt flag
	BIF   SR = 0x01 << 7 //+ Break interrupt flag
	CC1OF SR = 0x01 << 9 //+ Capture/Compare 1 overcapture flag
)

const (
	UIFn   = 0
	CC1IFn = 1
	COMIFn = 5
	BIFn   = 7
	CC1OFn = 9
)

const (
	UG   EGR = 0x01 << 0 //+ Update generation
	CC1G EGR = 0x01 << 1 //+ Capture/compare 1 generation
	COMG EGR = 0x01 << 5 //+ Capture/Compare control update generation
	BG   EGR = 0x01 << 7 //+ Break generation
)

const (
	UGn   = 0
	CC1Gn = 1
	COMGn = 5
	BGn   = 7
)

const (
	CC1S   CCMR1_Output = 0x03 << 0  //+ Capture/Compare 1 selection
	OC1FE  CCMR1_Output = 0x01 << 2  //+ Output Compare 1 fast enable
	OC1PE  CCMR1_Output = 0x01 << 3  //+ Output Compare 1 preload enable
	OC1M   CCMR1_Output = 0x07 << 4  //+ Output Compare 1 mode
	OC1M_3 CCMR1_Output = 0x01 << 16 //+ Output Compare 1 mode
)

const (
	CC1Sn   = 0
	OC1FEn  = 2
	OC1PEn  = 3
	OC1Mn   = 4
	OC1M_3n = 16
)

const (
	CC1S   CCMR1_Input = 0x03 << 0 //+ Capture/Compare 1 selection
	IC1PSC CCMR1_Input = 0x03 << 2 //+ Input capture 1 prescaler
	IC1F   CCMR1_Input = 0x0F << 4 //+ Input capture 1 filter
)

const (
	CC1Sn   = 0
	IC1PSCn = 2
	IC1Fn   = 4
)

const (
	CC1E  CCER = 0x01 << 0 //+ Capture/Compare 1 output enable
	CC1P  CCER = 0x01 << 1 //+ Capture/Compare 1 output Polarity
	CC1NE CCER = 0x01 << 2 //+ Capture/Compare 1 complementary output enable
	CC1NP CCER = 0x01 << 3 //+ Capture/Compare 1 output Polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NEn = 2
	CC1NPn = 3
)

const (
	CNT    CNT = 0xFFFF << 0 //+ counter value
	UIFCPY CNT = 0x01 << 31  //+ UIF Copy
)

const (
	CNTn    = 0
	UIFCPYn = 31
)

const (
	PSC PSC = 0xFFFF << 0 //+ Prescaler value
)

const (
	PSCn = 0
)

const (
	ARR ARR = 0xFFFF << 0 //+ Auto-reload value
)

const (
	ARRn = 0
)

const (
	REP RCR = 0xFF << 0 //+ Repetition counter value
)

const (
	REPn = 0
)

const (
	CCR1 CCR1 = 0xFFFF << 0 //+ Capture/Compare 1 value
)

const (
	CCR1n = 0
)

const (
	DTG    BDTR = 0xFF << 0  //+ Dead-time generator setup
	LOCK   BDTR = 0x03 << 8  //+ Lock configuration
	OSSI   BDTR = 0x01 << 10 //+ Off-state selection for Idle mode
	OSSR   BDTR = 0x01 << 11 //+ Off-state selection for Run mode
	BKE    BDTR = 0x01 << 12 //+ Break enable
	BKP    BDTR = 0x01 << 13 //+ Break polarity
	AOE    BDTR = 0x01 << 14 //+ Automatic output enable
	MOE    BDTR = 0x01 << 15 //+ Main output enable
	BKF    BDTR = 0x0F << 16 //+ Break filter
	BKDSRM BDTR = 0x01 << 26 //+ BKDSRM
	BKBID  BDTR = 0x01 << 28 //+ BKBID
)

const (
	DTGn    = 0
	LOCKn   = 8
	OSSIn   = 10
	OSSRn   = 11
	BKEn    = 12
	BKPn    = 13
	AOEn    = 14
	MOEn    = 15
	BKFn    = 16
	BKDSRMn = 26
	BKBIDn  = 28
)

const (
	DTGF DTR2 = 0xFF << 0  //+ Dead-time generator setup
	DTAE DTR2 = 0x01 << 16 //+ Deadtime Asymmetric Enable
	DTPE DTR2 = 0x01 << 17 //+ Deadtime Preload Enable
)

const (
	DTGFn = 0
	DTAEn = 16
	DTPEn = 17
)

const (
	TI1SEL TISEL = 0x0F << 0 //+ TI1[0] to TI1[15] input selection
)

const (
	TI1SELn = 0
)

const (
	BKINE   AF1 = 0x01 << 0  //+ BRK BKIN input enable
	BKCMP1E AF1 = 0x01 << 1  //+ BRK COMP1 enable
	BKCMP2E AF1 = 0x01 << 2  //+ BRK COMP2 enable
	BKCMP3E AF1 = 0x01 << 3  //+ BRK COMP3 enable
	BKCMP4E AF1 = 0x01 << 4  //+ BRK COMP4 enable
	BKCMP5E AF1 = 0x01 << 5  //+ BRK COMP5 enable
	BKCMP6E AF1 = 0x01 << 6  //+ BRK COMP6 enable
	BKCMP7E AF1 = 0x01 << 7  //+ BRK COMP7 enable
	BKINP   AF1 = 0x01 << 9  //+ BRK BKIN input polarity
	BKCMP1P AF1 = 0x01 << 10 //+ BRK COMP1 input polarity
	BKCMP2P AF1 = 0x01 << 11 //+ BRK COMP2 input polarity
	BKCMP3P AF1 = 0x01 << 12 //+ BRK COMP3 input polarity
	BKCMP4P AF1 = 0x01 << 13 //+ BRK COMP4 input polarity
)

const (
	BKINEn   = 0
	BKCMP1En = 1
	BKCMP2En = 2
	BKCMP3En = 3
	BKCMP4En = 4
	BKCMP5En = 5
	BKCMP6En = 6
	BKCMP7En = 7
	BKINPn   = 9
	BKCMP1Pn = 10
	BKCMP2Pn = 11
	BKCMP3Pn = 12
	BKCMP4Pn = 13
)

const (
	OCRSEL AF2 = 0x07 << 16 //+ OCREF_CLR source selection
)

const (
	OCRSELn = 16
)

const (
	DBA DCR = 0x1F << 0 //+ DMA base address
	DBL DCR = 0x1F << 8 //+ DMA burst length
)

const (
	DBAn = 0
	DBLn = 8
)

const (
	DMAB DMAR = 0xFFFFFFFF << 0 //+ DMA register for burst accesses
)

const (
	DMABn = 0
)

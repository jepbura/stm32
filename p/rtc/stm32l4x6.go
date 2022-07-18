// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package rtc provides access to the registers of the RTC peripheral.
//
// Instances:
//  RTC  RTC_BASE  -  TAMP_STAMP,RTC_WKUP,RTC_ALARM  Real-time clock
// Registers:
//  0x000 32  TR                 time register
//  0x004 32  DR                 date register
//  0x008 32  CR                 control register
//  0x00C 32  ISR                initialization and status register
//  0x010 32  PRER               prescaler register
//  0x014 32  WUTR               wakeup timer register
//  0x01C 32  ALRMAR(ALRMR)      alarm A register
//  0x020 32  ALRMBR(ALRMR)      alarm B register
//  0x024 32  WPR                write protection register
//  0x028 32  SSR                sub second register
//  0x02C 32  SHIFTR             shift control register
//  0x030 32  TSTR(TR)           time stamp time register
//  0x034 32  TSDR(DR)           time stamp date register
//  0x038 32  TSSSR(SSR)         timestamp sub second register
//  0x03C 32  CALR               calibration register
//  0x040 32  TAMPCR             tamper configuration register
//  0x044 32  ALRMASSR(ALRMSSR)  alarm A sub second register
//  0x048 32  ALRMBSSR(ALRMSSR)  alarm B sub second register
//  0x04C 32  OR                 option register
//  0x050 32  BKPR[32]           backup register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package rtc

const (
	SU  TR = 0x0F << 0  //+ Second units in BCD format
	ST  TR = 0x07 << 4  //+ Second tens in BCD format
	MNU TR = 0x0F << 8  //+ Minute units in BCD format
	MNT TR = 0x07 << 12 //+ Minute tens in BCD format
	HU  TR = 0x0F << 16 //+ Hour units in BCD format
	HT  TR = 0x03 << 20 //+ Hour tens in BCD format
	PM  TR = 0x01 << 22 //+ AM/PM notation
)

const (
	SUn  = 0
	STn  = 4
	MNUn = 8
	MNTn = 12
	HUn  = 16
	HTn  = 20
	PMn  = 22
)

const (
	DU  DR = 0x0F << 0  //+ Date units in BCD format
	DT  DR = 0x03 << 4  //+ Date tens in BCD format
	MU  DR = 0x0F << 8  //+ Month units in BCD format
	MT  DR = 0x01 << 12 //+ Month tens in BCD format
	WDU DR = 0x07 << 13 //+ Week day units
	YU  DR = 0x0F << 16 //+ Year units in BCD format
	YT  DR = 0x0F << 20 //+ Year tens in BCD format
)

const (
	DUn  = 0
	DTn  = 4
	MUn  = 8
	MTn  = 12
	WDUn = 13
	YUn  = 16
	YTn  = 20
)

const (
	WCKSEL  CR = 0x07 << 0  //+ Wakeup clock selection
	TSEDGE  CR = 0x01 << 3  //+ Time-stamp event active edge
	REFCKON CR = 0x01 << 4  //+ Reference clock detection enable (50 or 60 Hz)
	BYPSHAD CR = 0x01 << 5  //+ Bypass the shadow registers
	FMT     CR = 0x01 << 6  //+ Hour format
	ALRAE   CR = 0x01 << 8  //+ Alarm A enable
	ALRBE   CR = 0x01 << 9  //+ Alarm B enable
	WUTE    CR = 0x01 << 10 //+ Wakeup timer enable
	TSE     CR = 0x01 << 11 //+ Time stamp enable
	ALRAIE  CR = 0x01 << 12 //+ Alarm A interrupt enable
	ALRBIE  CR = 0x01 << 13 //+ Alarm B interrupt enable
	WUTIE   CR = 0x01 << 14 //+ Wakeup timer interrupt enable
	TSIE    CR = 0x01 << 15 //+ Time-stamp interrupt enable
	ADD1H   CR = 0x01 << 16 //+ Add 1 hour (summer time change)
	SUB1H   CR = 0x01 << 17 //+ Subtract 1 hour (winter time change)
	BKP     CR = 0x01 << 18 //+ Backup
	COSEL   CR = 0x01 << 19 //+ Calibration output selection
	POL     CR = 0x01 << 20 //+ Output polarity
	OSEL    CR = 0x03 << 21 //+ Output selection
	COE     CR = 0x01 << 23 //+ Calibration output enable
	ITSE    CR = 0x01 << 24 //+ timestamp on internal event enable
)

const (
	WCKSELn  = 0
	TSEDGEn  = 3
	REFCKONn = 4
	BYPSHADn = 5
	FMTn     = 6
	ALRAEn   = 8
	ALRBEn   = 9
	WUTEn    = 10
	TSEn     = 11
	ALRAIEn  = 12
	ALRBIEn  = 13
	WUTIEn   = 14
	TSIEn    = 15
	ADD1Hn   = 16
	SUB1Hn   = 17
	BKPn     = 18
	COSELn   = 19
	POLn     = 20
	OSELn    = 21
	COEn     = 23
	ITSEn    = 24
)

const (
	ALRAWF  ISR = 0x01 << 0  //+ Alarm A write flag
	ALRBWF  ISR = 0x01 << 1  //+ Alarm B write flag
	WUTWF   ISR = 0x01 << 2  //+ Wakeup timer write flag
	SHPF    ISR = 0x01 << 3  //+ Shift operation pending
	INITS   ISR = 0x01 << 4  //+ Initialization status flag
	RSF     ISR = 0x01 << 5  //+ Registers synchronization flag
	INITF   ISR = 0x01 << 6  //+ Initialization flag
	INIT    ISR = 0x01 << 7  //+ Initialization mode
	ALRAF   ISR = 0x01 << 8  //+ Alarm A flag
	ALRBF   ISR = 0x01 << 9  //+ Alarm B flag
	WUTF    ISR = 0x01 << 10 //+ Wakeup timer flag
	TSF     ISR = 0x01 << 11 //+ Time-stamp flag
	TSOVF   ISR = 0x01 << 12 //+ Time-stamp overflow flag
	TAMP1F  ISR = 0x01 << 13 //+ Tamper detection flag
	TAMP2F  ISR = 0x01 << 14 //+ RTC_TAMP2 detection flag
	TAMP3F  ISR = 0x01 << 15 //+ RTC_TAMP3 detection flag
	RECALPF ISR = 0x01 << 16 //+ Recalibration pending Flag
)

const (
	ALRAWFn  = 0
	ALRBWFn  = 1
	WUTWFn   = 2
	SHPFn    = 3
	INITSn   = 4
	RSFn     = 5
	INITFn   = 6
	INITn    = 7
	ALRAFn   = 8
	ALRBFn   = 9
	WUTFn    = 10
	TSFn     = 11
	TSOVFn   = 12
	TAMP1Fn  = 13
	TAMP2Fn  = 14
	TAMP3Fn  = 15
	RECALPFn = 16
)

const (
	PREDIV_S PRER = 0x7FFF << 0 //+ Synchronous prescaler factor
	PREDIV_A PRER = 0x7F << 16  //+ Asynchronous prescaler factor
)

const (
	PREDIV_Sn = 0
	PREDIV_An = 16
)

const (
	WUT WUTR = 0xFFFF << 0 //+ Wakeup auto-reload value bits
)

const (
	WUTn = 0
)

const (
	ASU    ALRMR = 0x0F << 0  //+ Second units in BCD format
	AST    ALRMR = 0x07 << 4  //+ Second tens in BCD format
	AMSK1  ALRMR = 0x01 << 7  //+ Alarm A seconds mask
	AMNU   ALRMR = 0x0F << 8  //+ Minute units in BCD format
	AMNT   ALRMR = 0x07 << 12 //+ Minute tens in BCD format
	AMSK2  ALRMR = 0x01 << 15 //+ Alarm A minutes mask
	AHU    ALRMR = 0x0F << 16 //+ Hour units in BCD format
	AHT    ALRMR = 0x03 << 20 //+ Hour tens in BCD format
	APM    ALRMR = 0x01 << 22 //+ AM/PM notation
	AMSK3  ALRMR = 0x01 << 23 //+ Alarm A hours mask
	ADU    ALRMR = 0x0F << 24 //+ Date units or day in BCD format
	ADT    ALRMR = 0x03 << 28 //+ Date tens in BCD format
	AWDSEL ALRMR = 0x01 << 30 //+ Week day selection
	AMSK4  ALRMR = 0x01 << 31 //+ Alarm A date mask
)

const (
	ASUn    = 0
	ASTn    = 4
	AMSK1n  = 7
	AMNUn   = 8
	AMNTn   = 12
	AMSK2n  = 15
	AHUn    = 16
	AHTn    = 20
	APMn    = 22
	AMSK3n  = 23
	ADUn    = 24
	ADTn    = 28
	AWDSELn = 30
	AMSK4n  = 31
)

const (
	KEY WPR = 0xFF << 0 //+ Write protection key
)

const (
	KEYn = 0
)

const (
	SS SSR = 0xFFFF << 0 //+ Sub second value
)

const (
	SSn = 0
)

const (
	SUBFS SHIFTR = 0x7FFF << 0 //+ Subtract a fraction of a second
	ADD1S SHIFTR = 0x01 << 31  //+ Add one second
)

const (
	SUBFSn = 0
	ADD1Sn = 31
)

const (
	CALM   CALR = 0x1FF << 0 //+ Calibration minus
	CALW16 CALR = 0x01 << 13 //+ Use a 16-second calibration cycle period
	CALW8  CALR = 0x01 << 14 //+ Use an 8-second calibration cycle period
	CALP   CALR = 0x01 << 15 //+ Increase frequency of RTC by 488.5 ppm
)

const (
	CALMn   = 0
	CALW16n = 13
	CALW8n  = 14
	CALPn   = 15
)

const (
	TAMP1E       TAMPCR = 0x01 << 0  //+ Tamper 1 detection enable
	TAMP1TRG     TAMPCR = 0x01 << 1  //+ Active level for tamper 1
	TAMPIE       TAMPCR = 0x01 << 2  //+ Tamper interrupt enable
	TAMP2E       TAMPCR = 0x01 << 3  //+ Tamper 2 detection enable
	TAMP2TRG     TAMPCR = 0x01 << 4  //+ Active level for tamper 2
	TAMP3E       TAMPCR = 0x01 << 5  //+ Tamper 3 detection enable
	TAMP3TRG     TAMPCR = 0x01 << 6  //+ Active level for tamper 3
	TAMPTS       TAMPCR = 0x01 << 7  //+ Activate timestamp on tamper detection event
	TAMPFREQ     TAMPCR = 0x07 << 8  //+ Tamper sampling frequency
	TAMPFLT      TAMPCR = 0x03 << 11 //+ Tamper filter count
	TAMPPRCH     TAMPCR = 0x03 << 13 //+ Tamper precharge duration
	TAMPPUDIS    TAMPCR = 0x01 << 15 //+ TAMPER pull-up disable
	TAMP1IE      TAMPCR = 0x01 << 16 //+ Tamper 1 interrupt enable
	TAMP1NOERASE TAMPCR = 0x01 << 17 //+ Tamper 1 no erase
	TAMP1MF      TAMPCR = 0x01 << 18 //+ Tamper 1 mask flag
	TAMP2IE      TAMPCR = 0x01 << 19 //+ Tamper 2 interrupt enable
	TAMP2NOERASE TAMPCR = 0x01 << 20 //+ Tamper 2 no erase
	TAMP2MF      TAMPCR = 0x01 << 21 //+ Tamper 2 mask flag
	TAMP3IE      TAMPCR = 0x01 << 22 //+ Tamper 3 interrupt enable
	TAMP3NOERASE TAMPCR = 0x01 << 23 //+ Tamper 3 no erase
	TAMP3MF      TAMPCR = 0x01 << 24 //+ Tamper 3 mask flag
)

const (
	TAMP1En       = 0
	TAMP1TRGn     = 1
	TAMPIEn       = 2
	TAMP2En       = 3
	TAMP2TRGn     = 4
	TAMP3En       = 5
	TAMP3TRGn     = 6
	TAMPTSn       = 7
	TAMPFREQn     = 8
	TAMPFLTn      = 11
	TAMPPRCHn     = 13
	TAMPPUDISn    = 15
	TAMP1IEn      = 16
	TAMP1NOERASEn = 17
	TAMP1MFn      = 18
	TAMP2IEn      = 19
	TAMP2NOERASEn = 20
	TAMP2MFn      = 21
	TAMP3IEn      = 22
	TAMP3NOERASEn = 23
	TAMP3MFn      = 24
)

const (
	ASS     ALRMSSR = 0x7FFF << 0 //+ Sub seconds value
	AMASKSS ALRMSSR = 0x0F << 24  //+ Mask the most-significant bits starting at this bit
)

const (
	ASSn     = 0
	AMASKSSn = 24
)

const (
	RTC_ALARM_TYPE OR = 0x01 << 0 //+ RTC_ALARM on PC13 output type
	RTC_OUT_RMP    OR = 0x01 << 1 //+ RTC_OUT remap
)

const (
	RTC_ALARM_TYPEn = 0
	RTC_OUT_RMPn    = 1
)

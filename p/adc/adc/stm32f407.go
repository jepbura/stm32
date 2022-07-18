// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f407

// Package adc provides access to the registers of the ADC peripheral.
//
// Instances:
//  ADC1  ADC1_BASE  APB2  ADC+  Analog-to-digital converter
//  ADC2  ADC2_BASE  APB2  ADC+  Analog-to-digital converter
//  ADC3  ADC3_BASE  APB2  ADC+  Analog-to-digital converter
// Registers:
//  0x000 32  SR     status register
//  0x004 32  CR1    control register 1
//  0x008 32  CR2    control register 2
//  0x00C 32  SMPR1  sample time register 1
//  0x010 32  SMPR2  sample time register 2
//  0x014 32  JOFR1  injected channel data offset register x
//  0x018 32  JOFR2  injected channel data offset register x
//  0x01C 32  JOFR3  injected channel data offset register x
//  0x020 32  JOFR4  injected channel data offset register x
//  0x024 32  HTR    watchdog higher threshold register
//  0x028 32  LTR    watchdog lower threshold register
//  0x02C 32  SQR1   regular sequence register 1
//  0x030 32  SQR2   regular sequence register 2
//  0x034 32  SQR3   regular sequence register 3
//  0x038 32  JSQR   injected sequence register
//  0x03C 32  JDR1   injected data register x
//  0x040 32  JDR2   injected data register x
//  0x044 32  JDR3   injected data register x
//  0x048 32  JDR4   injected data register x
//  0x04C 32  DR     regular data register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package adc

const (
	AWD   SR = 0x01 << 0 //+ Analog watchdog flag
	EOC   SR = 0x01 << 1 //+ Regular channel end of conversion
	JEOC  SR = 0x01 << 2 //+ Injected channel end of conversion
	JSTRT SR = 0x01 << 3 //+ Injected channel start flag
	STRT  SR = 0x01 << 4 //+ Regular channel start flag
	OVR   SR = 0x01 << 5 //+ Overrun
)

const (
	AWDn   = 0
	EOCn   = 1
	JEOCn  = 2
	JSTRTn = 3
	STRTn  = 4
	OVRn   = 5
)

const (
	AWDCH   CR1 = 0x1F << 0  //+ Analog watchdog channel select bits
	EOCIE   CR1 = 0x01 << 5  //+ Interrupt enable for EOC
	AWDIE   CR1 = 0x01 << 6  //+ Analog watchdog interrupt enable
	JEOCIE  CR1 = 0x01 << 7  //+ Interrupt enable for injected channels
	SCAN    CR1 = 0x01 << 8  //+ Scan mode
	AWDSGL  CR1 = 0x01 << 9  //+ Enable the watchdog on a single channel in scan mode
	JAUTO   CR1 = 0x01 << 10 //+ Automatic injected group conversion
	DISCEN  CR1 = 0x01 << 11 //+ Discontinuous mode on regular channels
	JDISCEN CR1 = 0x01 << 12 //+ Discontinuous mode on injected channels
	DISCNUM CR1 = 0x07 << 13 //+ Discontinuous mode channel count
	JAWDEN  CR1 = 0x01 << 22 //+ Analog watchdog enable on injected channels
	AWDEN   CR1 = 0x01 << 23 //+ Analog watchdog enable on regular channels
	RES     CR1 = 0x03 << 24 //+ Resolution
	OVRIE   CR1 = 0x01 << 26 //+ Overrun interrupt enable
)

const (
	AWDCHn   = 0
	EOCIEn   = 5
	AWDIEn   = 6
	JEOCIEn  = 7
	SCANn    = 8
	AWDSGLn  = 9
	JAUTOn   = 10
	DISCENn  = 11
	JDISCENn = 12
	DISCNUMn = 13
	JAWDENn  = 22
	AWDENn   = 23
	RESn     = 24
	OVRIEn   = 26
)

const (
	ADON     CR2 = 0x01 << 0  //+ A/D Converter ON / OFF
	CONT     CR2 = 0x01 << 1  //+ Continuous conversion
	DMA      CR2 = 0x01 << 8  //+ Direct memory access mode (for single ADC mode)
	DDS      CR2 = 0x01 << 9  //+ DMA disable selection (for single ADC mode)
	EOCS     CR2 = 0x01 << 10 //+ End of conversion selection
	ALIGN    CR2 = 0x01 << 11 //+ Data alignment
	JEXTSEL  CR2 = 0x0F << 16 //+ External event select for injected group
	JEXTEN   CR2 = 0x03 << 20 //+ External trigger enable for injected channels
	JSWSTART CR2 = 0x01 << 22 //+ Start conversion of injected channels
	EXTSEL   CR2 = 0x0F << 24 //+ External event select for regular group
	EXTEN    CR2 = 0x03 << 28 //+ External trigger enable for regular channels
	SWSTART  CR2 = 0x01 << 30 //+ Start conversion of regular channels
)

const (
	ADONn     = 0
	CONTn     = 1
	DMAn      = 8
	DDSn      = 9
	EOCSn     = 10
	ALIGNn    = 11
	JEXTSELn  = 16
	JEXTENn   = 20
	JSWSTARTn = 22
	EXTSELn   = 24
	EXTENn    = 28
	SWSTARTn  = 30
)

const (
	SMPx_x SMPR1 = 0xFFFFFFFF << 0 //+ Sample time bits
)

const (
	SMPx_xn = 0
)

const (
	SMPx_x SMPR2 = 0xFFFFFFFF << 0 //+ Sample time bits
)

const (
	SMPx_xn = 0
)

const (
	JOFFSET1 JOFR1 = 0xFFF << 0 //+ Data offset for injected channel x
)

const (
	JOFFSET1n = 0
)

const (
	JOFFSET2 JOFR2 = 0xFFF << 0 //+ Data offset for injected channel x
)

const (
	JOFFSET2n = 0
)

const (
	JOFFSET3 JOFR3 = 0xFFF << 0 //+ Data offset for injected channel x
)

const (
	JOFFSET3n = 0
)

const (
	JOFFSET4 JOFR4 = 0xFFF << 0 //+ Data offset for injected channel x
)

const (
	JOFFSET4n = 0
)

const (
	HT HTR = 0xFFF << 0 //+ Analog watchdog higher threshold
)

const (
	HTn = 0
)

const (
	LT LTR = 0xFFF << 0 //+ Analog watchdog lower threshold
)

const (
	LTn = 0
)

const (
	SQ13 SQR1 = 0x1F << 0  //+ 13th conversion in regular sequence
	SQ14 SQR1 = 0x1F << 5  //+ 14th conversion in regular sequence
	SQ15 SQR1 = 0x1F << 10 //+ 15th conversion in regular sequence
	SQ16 SQR1 = 0x1F << 15 //+ 16th conversion in regular sequence
	L    SQR1 = 0x0F << 20 //+ Regular channel sequence length
)

const (
	SQ13n = 0
	SQ14n = 5
	SQ15n = 10
	SQ16n = 15
	Ln    = 20
)

const (
	SQ7  SQR2 = 0x1F << 0  //+ 7th conversion in regular sequence
	SQ8  SQR2 = 0x1F << 5  //+ 8th conversion in regular sequence
	SQ9  SQR2 = 0x1F << 10 //+ 9th conversion in regular sequence
	SQ10 SQR2 = 0x1F << 15 //+ 10th conversion in regular sequence
	SQ11 SQR2 = 0x1F << 20 //+ 11th conversion in regular sequence
	SQ12 SQR2 = 0x1F << 25 //+ 12th conversion in regular sequence
)

const (
	SQ7n  = 0
	SQ8n  = 5
	SQ9n  = 10
	SQ10n = 15
	SQ11n = 20
	SQ12n = 25
)

const (
	SQ1 SQR3 = 0x1F << 0  //+ 1st conversion in regular sequence
	SQ2 SQR3 = 0x1F << 5  //+ 2nd conversion in regular sequence
	SQ3 SQR3 = 0x1F << 10 //+ 3rd conversion in regular sequence
	SQ4 SQR3 = 0x1F << 15 //+ 4th conversion in regular sequence
	SQ5 SQR3 = 0x1F << 20 //+ 5th conversion in regular sequence
	SQ6 SQR3 = 0x1F << 25 //+ 6th conversion in regular sequence
)

const (
	SQ1n = 0
	SQ2n = 5
	SQ3n = 10
	SQ4n = 15
	SQ5n = 20
	SQ6n = 25
)

const (
	JSQ1 JSQR = 0x1F << 0  //+ 1st conversion in injected sequence
	JSQ2 JSQR = 0x1F << 5  //+ 2nd conversion in injected sequence
	JSQ3 JSQR = 0x1F << 10 //+ 3rd conversion in injected sequence
	JSQ4 JSQR = 0x1F << 15 //+ 4th conversion in injected sequence
	JL   JSQR = 0x03 << 20 //+ Injected sequence length
)

const (
	JSQ1n = 0
	JSQ2n = 5
	JSQ3n = 10
	JSQ4n = 15
	JLn   = 20
)

const (
	JDATA JDR1 = 0xFFFF << 0 //+ Injected data
)

const (
	JDATAn = 0
)

const (
	JDATA JDR2 = 0xFFFF << 0 //+ Injected data
)

const (
	JDATAn = 0
)

const (
	JDATA JDR3 = 0xFFFF << 0 //+ Injected data
)

const (
	JDATAn = 0
)

const (
	JDATA JDR4 = 0xFFFF << 0 //+ Injected data
)

const (
	JDATAn = 0
)

const (
	DATA DR = 0xFFFF << 0 //+ Regular data
)

const (
	DATAn = 0
)

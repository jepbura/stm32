// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f407

// Package tim10 provides access to the registers of the TIM peripheral.
//
// Instances:
//  TIM10  TIM10_BASE  APB2  TIM1_UP_TIM10*       General-purpose-timers
//  TIM13  TIM13_BASE  APB1  TIM8_UP_TIM13*       General-purpose-timers
//  TIM14  TIM14_BASE  APB1  TIM8_TRG_COM_TIM14*  General-purpose-timers
// Registers:
//  0x000 32  CR1        control register 1
//  0x00C 32  DIER       DMA/Interrupt enable register
//  0x010 32  SR         status register
//  0x014 32  EGR        event generation register
//  0x018 32  CCMR1      capture/compare mode register 1
//  0x020 32  CCER       capture/compare enable register
//  0x024 32  CNT        counter
//  0x028 32  PSC        prescaler
//  0x02C 32  ARR        auto-reload register
//  0x034 32  CCR1(CCR)  capture/compare register 1
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package tim10

const (
	CEN  CR1 = 0x01 << 0 //+ Counter enable
	UDIS CR1 = 0x01 << 1 //+ Update disable
	URS  CR1 = 0x01 << 2 //+ Update request source
	ARPE CR1 = 0x01 << 7 //+ Auto-reload preload enable
	CKD  CR1 = 0x03 << 8 //+ Clock division
)

const (
	CENn  = 0
	UDISn = 1
	URSn  = 2
	ARPEn = 7
	CKDn  = 8
)

const (
	UIE   DIER = 0x01 << 0 //+ Update interrupt enable
	CC1IE DIER = 0x01 << 1 //+ Capture/Compare 1 interrupt enable
)

const (
	UIEn   = 0
	CC1IEn = 1
)

const (
	UIF   SR = 0x01 << 0 //+ Update interrupt flag
	CC1IF SR = 0x01 << 1 //+ Capture/compare 1 interrupt flag
	CC1OF SR = 0x01 << 9 //+ Capture/Compare 1 overcapture flag
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC1OFn = 9
)

const (
	UG   EGR = 0x01 << 0 //+ Update generation
	CC1G EGR = 0x01 << 1 //+ Capture/compare 1 generation
)

const (
	UGn   = 0
	CC1Gn = 1
)

const (
	CC1S  CCMR1 = 0x03 << 0 //+ Capture/Compare 1 selection
	OC1FE CCMR1 = 0x01 << 2 //+ Output Compare 1 fast enable
	OC1PE CCMR1 = 0x01 << 3 //+ Output Compare 1 preload enable
	OC1M  CCMR1 = 0x07 << 4 //+ Output Compare 1 mode
	ICPCS CCMR1 = 0x03 << 2 //+ Input capture 1 prescaler
	IC1F  CCMR1 = 0x0F << 4 //+ Input capture 1 filter
)

const (
	CC1Sn  = 0
	OC1FEn = 2
	OC1PEn = 3
	OC1Mn  = 4
	ICPCSn = 2
	IC1Fn  = 4
)

const (
	CC1E  CCER = 0x01 << 0 //+ Capture/Compare 1 output enable
	CC1P  CCER = 0x01 << 1 //+ Capture/Compare 1 output Polarity
	CC1NP CCER = 0x01 << 3 //+ Capture/Compare 1 output Polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NPn = 3
)

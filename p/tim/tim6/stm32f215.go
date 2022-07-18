// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

// Package tim6 provides access to the registers of the TIM peripheral.
//
// Instances:
//  TIM6  TIM6_BASE  APB1  TIM6_DAC*  Basic-timers
//  TIM7  TIM7_BASE  APB1  TIM7       Basic-timers
// Registers:
//  0x000 32  CR1   control register 1
//  0x004 32  CR2   control register 2
//  0x00C 32  DIER  DMA/Interrupt enable register
//  0x010 32  SR    status register
//  0x014 32  EGR   event generation register
//  0x024 32  CNT   counter
//  0x028 32  PSC   prescaler
//  0x02C 32  ARR   auto-reload register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package tim6

const (
	CEN  CR1 = 0x01 << 0 //+ Counter enable
	UDIS CR1 = 0x01 << 1 //+ Update disable
	URS  CR1 = 0x01 << 2 //+ Update request source
	OPM  CR1 = 0x01 << 3 //+ One-pulse mode
	ARPE CR1 = 0x01 << 7 //+ Auto-reload preload enable
)

const (
	CENn  = 0
	UDISn = 1
	URSn  = 2
	OPMn  = 3
	ARPEn = 7
)

const (
	MMS CR2 = 0x07 << 4 //+ Master mode selection
)

const (
	MMSn = 4
)

const (
	UIE DIER = 0x01 << 0 //+ Update interrupt enable
	UDE DIER = 0x01 << 8 //+ Update DMA request enable
)

const (
	UIEn = 0
	UDEn = 8
)

const (
	UIF SR = 0x01 << 0 //+ Update interrupt flag
)

const (
	UIFn = 0
)

const (
	UG EGR = 0x01 << 0 //+ Update generation
)

const (
	UGn = 0
)

const (
	CNT CNT = 0xFFFF << 0 //+ Low counter value
)

const (
	CNTn = 0
)

const (
	PSC PSC = 0xFFFF << 0 //+ Prescaler value
)

const (
	PSCn = 0
)

const (
	ARR ARR = 0xFFFF << 0 //+ Low Auto-reload value
)

const (
	ARRn = 0
)

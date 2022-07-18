// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package iwdg provides access to the registers of the IWDG peripheral.
//
// Instances:
//  IWDG  IWDG_BASE  -  -  Independent watchdog
// Registers:
//  0x000 32  KR   Key register
//  0x004 32  PR   Prescaler register
//  0x008 32  RLR  Reload register
//  0x00C 32  SR   Status register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package iwdg

const (
	KEY KR = 0xFFFF << 0 //+ Key value
)

const (
	KEYn = 0
)

const (
	PR PR = 0x07 << 0 //+ Prescaler divider
)

const (
	PRn = 0
)

const (
	RL RLR = 0xFFF << 0 //+ Watchdog counter reload value
)

const (
	RLn = 0
)

const (
	PVU SR = 0x01 << 0 //+ Watchdog prescaler value update
	RVU SR = 0x01 << 1 //+ Watchdog counter reload value update
)

const (
	PVUn = 0
	RVUn = 1
)

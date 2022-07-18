// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package dbg provides access to the registers of the DBG peripheral.
//
// Instances:
//  DBG  DBG_BASE  -  -  Debug support
// Registers:
//  0x000 32  DBGMCU_IDCODE   IDCODE
//  0x004 32  DBGMCU_CR       Control Register
//  0x008 32  DBGMCU_APB1_FZ  Debug MCU APB1 Freeze registe
//  0x00C 32  DBGMCU_APB2_FZ  Debug MCU APB2 Freeze registe
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package dbg

const (
	DEV_ID DBGMCU_IDCODE = 0xFFF << 0   //+ DEV_ID
	REV_ID DBGMCU_IDCODE = 0xFFFF << 16 //+ REV_ID
)

const (
	DEV_IDn = 0
	REV_IDn = 16
)

const (
	DBG_SLEEP   DBGMCU_CR = 0x01 << 0 //+ DBG_SLEEP
	DBG_STOP    DBGMCU_CR = 0x01 << 1 //+ DBG_STOP
	DBG_STANDBY DBGMCU_CR = 0x01 << 2 //+ DBG_STANDBY
	TRACE_IOEN  DBGMCU_CR = 0x01 << 5 //+ TRACE_IOEN
	TRACE_MODE  DBGMCU_CR = 0x03 << 6 //+ TRACE_MODE
)

const (
	DBG_SLEEPn   = 0
	DBG_STOPn    = 1
	DBG_STANDBYn = 2
	TRACE_IOENn  = 5
	TRACE_MODEn  = 6
)

const (
	DBG_TIM2_STOP          DBGMCU_APB1_FZ = 0x01 << 0  //+ DBG_TIM2_STOP
	DBG_TIM3_STOP          DBGMCU_APB1_FZ = 0x01 << 1  //+ DBG_TIM3 _STOP
	DBG_TIM4_STOP          DBGMCU_APB1_FZ = 0x01 << 2  //+ DBG_TIM4_STOP
	DBG_TIM5_STOP          DBGMCU_APB1_FZ = 0x01 << 3  //+ DBG_TIM5_STOP
	DBG_RTC_Stop           DBGMCU_APB1_FZ = 0x01 << 10 //+ RTC stopped when Core is halted
	DBG_WWDG_STOP          DBGMCU_APB1_FZ = 0x01 << 11 //+ DBG_WWDG_STOP
	DBG_IWDEG_STOP         DBGMCU_APB1_FZ = 0x01 << 12 //+ DBG_IWDEG_STOP
	DBG_I2C1_SMBUS_TIMEOUT DBGMCU_APB1_FZ = 0x01 << 21 //+ DBG_J2C1_SMBUS_TIMEOUT
	DBG_I2C2_SMBUS_TIMEOUT DBGMCU_APB1_FZ = 0x01 << 22 //+ DBG_J2C2_SMBUS_TIMEOUT
	DBG_I2C3SMBUS_TIMEOUT  DBGMCU_APB1_FZ = 0x01 << 23 //+ DBG_J2C3SMBUS_TIMEOUT
)

const (
	DBG_TIM2_STOPn          = 0
	DBG_TIM3_STOPn          = 1
	DBG_TIM4_STOPn          = 2
	DBG_TIM5_STOPn          = 3
	DBG_RTC_Stopn           = 10
	DBG_WWDG_STOPn          = 11
	DBG_IWDEG_STOPn         = 12
	DBG_I2C1_SMBUS_TIMEOUTn = 21
	DBG_I2C2_SMBUS_TIMEOUTn = 22
	DBG_I2C3SMBUS_TIMEOUTn  = 23
)

const (
	DBG_TIM1_STOP  DBGMCU_APB2_FZ = 0x01 << 0  //+ TIM1 counter stopped when core is halted
	DBG_TIM9_STOP  DBGMCU_APB2_FZ = 0x01 << 16 //+ TIM9 counter stopped when core is halted
	DBG_TIM10_STOP DBGMCU_APB2_FZ = 0x01 << 17 //+ TIM10 counter stopped when core is halted
	DBG_TIM11_STOP DBGMCU_APB2_FZ = 0x01 << 18 //+ TIM11 counter stopped when core is halted
)

const (
	DBG_TIM1_STOPn  = 0
	DBG_TIM9_STOPn  = 16
	DBG_TIM10_STOPn = 17
	DBG_TIM11_STOPn = 18
)

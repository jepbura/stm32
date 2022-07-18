// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package exti provides access to the registers of the EXTI peripheral.
//
// Instances:
//  EXTI  EXTI_BASE  -  TAMP_STAMP,EXTI0,EXTI1,EXTI2,EXTI3,EXTI4,EXTI9_5,EXTI15_10  External interrupt/event controller
// Registers:
//  0x000 32  IMR    Interrupt mask register (EXTI_IMR)
//  0x004 32  EMR    Event mask register (EXTI_EMR)
//  0x008 32  RTSR   Rising Trigger selection register (EXTI_RTSR)
//  0x00C 32  FTSR   Falling Trigger selection register (EXTI_FTSR)
//  0x010 32  SWIER  Software interrupt event register (EXTI_SWIER)
//  0x014 32  PR     Pending register (EXTI_PR)
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package exti

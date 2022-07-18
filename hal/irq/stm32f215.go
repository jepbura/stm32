// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

// Package irq provides the list of supported external interrupts.
package irq

import "embedded/rtos"

const (
	WWDG               rtos.IRQ = 0  // Window Watchdog interrupt
	PVD                rtos.IRQ = 1  // PVD through EXTI line detection interrupt
	TAMP_STAMP         rtos.IRQ = 2  // Tamper and TimeStamp interrupts through the EXTI line
	RTC_WKUP           rtos.IRQ = 3  // RTC Wakeup interrupt through the EXTI line
	FLASH              rtos.IRQ = 4  // FlASH global interrupt
	RCC                rtos.IRQ = 5  // RCC global interrupt
	EXTI0              rtos.IRQ = 6  // EXTI Line0 interrupt
	EXTI1              rtos.IRQ = 7  // EXTI Line1 interrupt
	EXTI2              rtos.IRQ = 8  // EXTI Line2 interrupt
	EXTI3              rtos.IRQ = 9  // EXTI Line3 interrupt
	EXTI4              rtos.IRQ = 10 // EXTI Line4 interrupt
	DMA1_STREAM0       rtos.IRQ = 11 // DMA1 Stream0 global interrupt
	DMA1_STREAM1       rtos.IRQ = 12 // DMA1 Stream1 global interrupt
	DMA1_STREAM2       rtos.IRQ = 13 // DMA1 Stream2 global interrupt
	DMA1_STREAM3       rtos.IRQ = 14 // DMA1 Stream3 global interrupt
	DMA1_STREAM4       rtos.IRQ = 15 // DMA1 Stream4 global interrupt
	DMA1_STREAM5       rtos.IRQ = 16 // DMA1 Stream5 global interrupt
	DMA1_STREAM6       rtos.IRQ = 17 // DMA1 Stream6 global interrupt
	ADC                rtos.IRQ = 18 // ADC1 global interrupt
	CAN1_TX            rtos.IRQ = 19 // CAN1 TX interrupts
	CAN1_RX0           rtos.IRQ = 20 // CAN1 RX0 interrupts
	CAN1_RX1           rtos.IRQ = 21 // CAN1 RX1 interrupts
	CAN1_SCE           rtos.IRQ = 22 // CAN1 SCE interrupt
	EXTI9_5            rtos.IRQ = 23 // EXTI Line[9:5] interrupts
	TIM1_BRK_TIM9      rtos.IRQ = 24 // TIM1 Break interrupt and TIM9 global interrupt
	TIM1_UP_TIM10      rtos.IRQ = 25 // TIM1 Update interrupt and TIM10 global interrupt
	TIM1_TRG_COM_TIM11 rtos.IRQ = 26 // TIM1 Trigger and Commutation interrupts and TIM11 global interrupt
	TIM1_CC            rtos.IRQ = 27 // TIM1 Capture Compare interrupt
	TIM2               rtos.IRQ = 28 // TIM2 global interrupt
	TIM3               rtos.IRQ = 29 // TIM3 global interrupt
	TIM4               rtos.IRQ = 30 // TIM4 global interrupt
	I2C1_EV            rtos.IRQ = 31 // I2C1 event interrupt
	I2C1_ER            rtos.IRQ = 32 // I2C1 error interrupt
	I2C2_EV            rtos.IRQ = 33 // I2C2 event interrupt
	I2C2_ER            rtos.IRQ = 34 // I2C2 error interrupt
	SPI1               rtos.IRQ = 35 // SPI1 global interrupt
	SPI2               rtos.IRQ = 36 // SPI2 global interrupt
	USART1             rtos.IRQ = 37 // USART1 global interrupt
	USART2             rtos.IRQ = 38 // USART2 global interrupt
	USART3             rtos.IRQ = 39 // USART3 global interrupt
	EXTI15_10          rtos.IRQ = 40 // EXTI Line[15:10] interrupts
	RTC_ALARM          rtos.IRQ = 41 // RTC Alarms (A and B) through EXTI line interrupt
	OTG_FS_WKUP        rtos.IRQ = 42 // USB On-The-Go FS Wakeup through EXTI line interrupt
	TIM8_BRK_TIM12     rtos.IRQ = 43 // TIM8 Break interrupt and TIM12 global interrupt
	TIM8_UP_TIM13      rtos.IRQ = 44 // TIM8 Update interrupt and TIM13 global interrupt
	TIM8_TRG_COM_TIM14 rtos.IRQ = 45 // TIM8 Trigger and Commutation interrupts and TIM14 global interrupt
	TIM8_CC            rtos.IRQ = 46 // TIM8 Capture Compare interrupt
	DMA1_STREAM7       rtos.IRQ = 47 // DMA1 Stream7 global interrupt
	FSMC               rtos.IRQ = 48 // FSMC global interrupt
	SDIO               rtos.IRQ = 49 // SDIO global interrupt
	TIM5               rtos.IRQ = 50 // TIM5 global interrupt
	SPI3               rtos.IRQ = 51 // SPI3 global interrupt
	UART4              rtos.IRQ = 52 // UART4 global interrupt
	UART5              rtos.IRQ = 53 // UART5 global interrupt
	TIM6_DAC           rtos.IRQ = 54 // TIM6 global interrupt, DAC1 and DAC2 underrun error interrupt
	TIM7               rtos.IRQ = 55 // TIM7 global interrupt
	DMA2_STREAM0       rtos.IRQ = 56 // DMA2 Stream0 global interrupt
	DMA2_STREAM1       rtos.IRQ = 57 // DMA2 Stream1 global interrupt
	DMA2_STREAM2       rtos.IRQ = 58 // DMA2 Stream2 global interrupt
	DMA2_STREAM3       rtos.IRQ = 59 // DMA2 Stream3 global interrupt
	DMA2_STREAM4       rtos.IRQ = 60 // DMA2 Stream4 global interrupt
	ETH                rtos.IRQ = 61 // Ethernet global interrupt
	ETH_WKUP           rtos.IRQ = 62 // Ethernet Wakeup through EXTI line interrupt
	CAN2_TX            rtos.IRQ = 63 // CAN2 TX interrupts
	CAN2_RX0           rtos.IRQ = 64 // CAN2 RX0 interrupts
	CAN2_RX1           rtos.IRQ = 65 // CAN2 RX1 interrupts
	CAN2_SCE           rtos.IRQ = 66 // CAN2 SCE interrupt
	OTG_FS             rtos.IRQ = 67 // USB On The Go FS global interrupt
	DMA2_STREAM5       rtos.IRQ = 68 // DMA2 Stream5 global interrupt
	DMA2_STREAM6       rtos.IRQ = 69 // DMA2 Stream6 global interrupt
	DMA2_STREAM7       rtos.IRQ = 70 // DMA2 Stream7 global interrupt
	USART6             rtos.IRQ = 71 // USART6 global interrupt
	I2C3_EV            rtos.IRQ = 72 // I2C3 event interrupt
	I2C3_ER            rtos.IRQ = 73 // I2C3 error interrupt
	OTG_HS_EP1_OUT     rtos.IRQ = 74 // USB On The Go HS End Point 1 Out global interrupt
	OTG_HS_EP1_IN      rtos.IRQ = 75 // USB On The Go HS End Point 1 In global interrupt
	OTG_HS_WKUP        rtos.IRQ = 76 // USB On The Go HS Wakeup through EXTI interrupt
	OTG_HS             rtos.IRQ = 77 // USB On The Go HS global interrupt
	DCMI               rtos.IRQ = 78 // DCMI global interrupt
	CRYP               rtos.IRQ = 79 // CRYP crypto global interrupt
	HASH_RNG           rtos.IRQ = 80 // Hash and Rng global interrupt
)

// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package dma provides access to the registers of the DMA1 peripheral.
//
// Instances:
//  DMA1  DMA1_BASE  AHB1  DMA1_CH1,DMA1_CH2,DMA1_CH3,DMA1_CH4,DMA1_CH5,DMA1_CH6,DMA1_CH7  Direct memory access controller
//  DMA2  DMA2_BASE  AHB1  DMA2_CH1,DMA2_CH2,DMA2_CH3,DMA2_CH4,DMA2_CH5,DMA2_CH6,DMA2_CH7  Direct memory access controller
// Registers:
//  0x000 32  ISR     interrupt status register
//  0x004 32  IFCR    interrupt flag clear register
//  0x008 32  CCR1    channel x configuration register
//  0x00C 32  CNDTR1  channel x number of data register
//  0x010 32  CPAR1   channel x peripheral address register
//  0x014 32  CMAR1   channel x memory address register
//  0x01C 32  CCR2    channel x configuration register
//  0x020 32  CNDTR2  channel x number of data register
//  0x024 32  CPAR2   channel x peripheral address register
//  0x028 32  CMAR2   channel x memory address register
//  0x030 32  CCR3    channel x configuration register
//  0x034 32  CNDTR3  channel x number of data register
//  0x038 32  CPAR3   channel x peripheral address register
//  0x03C 32  CMAR3   channel x memory address register
//  0x044 32  CCR4    channel x configuration register
//  0x048 32  CNDTR4  channel x number of data register
//  0x04C 32  CPAR4   channel x peripheral address register
//  0x050 32  CMAR4   channel x memory address register
//  0x058 32  CCR5    channel x configuration register
//  0x05C 32  CNDTR5  channel x number of data register
//  0x060 32  CPAR5   channel x peripheral address register
//  0x064 32  CMAR5   channel x memory address register
//  0x06C 32  CCR6    channel x configuration register
//  0x070 32  CNDTR6  channel x number of data register
//  0x074 32  CPAR6   channel x peripheral address register
//  0x078 32  CMAR6   channel x memory address register
//  0x080 32  CCR7    channel x configuration register
//  0x084 32  CNDTR7  channel x number of data register
//  0x088 32  CPAR7   channel x peripheral address register
//  0x08C 32  CMAR7   channel x memory address register
//  0x0A8 32  CSELR   channel selection register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package dma

const (
	GIF1  ISR = 0x01 << 0  //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF1 ISR = 0x01 << 1  //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF1 ISR = 0x01 << 2  //+ Channel x half transfer flag (x = 1 ..7)
	TEIF1 ISR = 0x01 << 3  //+ Channel x transfer error flag (x = 1 ..7)
	GIF2  ISR = 0x01 << 4  //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF2 ISR = 0x01 << 5  //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF2 ISR = 0x01 << 6  //+ Channel x half transfer flag (x = 1 ..7)
	TEIF2 ISR = 0x01 << 7  //+ Channel x transfer error flag (x = 1 ..7)
	GIF3  ISR = 0x01 << 8  //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF3 ISR = 0x01 << 9  //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF3 ISR = 0x01 << 10 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF3 ISR = 0x01 << 11 //+ Channel x transfer error flag (x = 1 ..7)
	GIF4  ISR = 0x01 << 12 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF4 ISR = 0x01 << 13 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF4 ISR = 0x01 << 14 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF4 ISR = 0x01 << 15 //+ Channel x transfer error flag (x = 1 ..7)
	GIF5  ISR = 0x01 << 16 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF5 ISR = 0x01 << 17 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF5 ISR = 0x01 << 18 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF5 ISR = 0x01 << 19 //+ Channel x transfer error flag (x = 1 ..7)
	GIF6  ISR = 0x01 << 20 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF6 ISR = 0x01 << 21 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF6 ISR = 0x01 << 22 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF6 ISR = 0x01 << 23 //+ Channel x transfer error flag (x = 1 ..7)
	GIF7  ISR = 0x01 << 24 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF7 ISR = 0x01 << 25 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF7 ISR = 0x01 << 26 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF7 ISR = 0x01 << 27 //+ Channel x transfer error flag (x = 1 ..7)
)

const (
	GIF1n  = 0
	TCIF1n = 1
	HTIF1n = 2
	TEIF1n = 3
	GIF2n  = 4
	TCIF2n = 5
	HTIF2n = 6
	TEIF2n = 7
	GIF3n  = 8
	TCIF3n = 9
	HTIF3n = 10
	TEIF3n = 11
	GIF4n  = 12
	TCIF4n = 13
	HTIF4n = 14
	TEIF4n = 15
	GIF5n  = 16
	TCIF5n = 17
	HTIF5n = 18
	TEIF5n = 19
	GIF6n  = 20
	TCIF6n = 21
	HTIF6n = 22
	TEIF6n = 23
	GIF7n  = 24
	TCIF7n = 25
	HTIF7n = 26
	TEIF7n = 27
)

const (
	CGIF1  IFCR = 0x01 << 0  //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF1 IFCR = 0x01 << 1  //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF1 IFCR = 0x01 << 2  //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF1 IFCR = 0x01 << 3  //+ Channel x transfer error clear (x = 1 ..7)
	CGIF2  IFCR = 0x01 << 4  //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF2 IFCR = 0x01 << 5  //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF2 IFCR = 0x01 << 6  //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF2 IFCR = 0x01 << 7  //+ Channel x transfer error clear (x = 1 ..7)
	CGIF3  IFCR = 0x01 << 8  //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF3 IFCR = 0x01 << 9  //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF3 IFCR = 0x01 << 10 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF3 IFCR = 0x01 << 11 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF4  IFCR = 0x01 << 12 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF4 IFCR = 0x01 << 13 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF4 IFCR = 0x01 << 14 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF4 IFCR = 0x01 << 15 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF5  IFCR = 0x01 << 16 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF5 IFCR = 0x01 << 17 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF5 IFCR = 0x01 << 18 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF5 IFCR = 0x01 << 19 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF6  IFCR = 0x01 << 20 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF6 IFCR = 0x01 << 21 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF6 IFCR = 0x01 << 22 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF6 IFCR = 0x01 << 23 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF7  IFCR = 0x01 << 24 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF7 IFCR = 0x01 << 25 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF7 IFCR = 0x01 << 26 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF7 IFCR = 0x01 << 27 //+ Channel x transfer error clear (x = 1 ..7)
)

const (
	CGIF1n  = 0
	CTCIF1n = 1
	CHTIF1n = 2
	CTEIF1n = 3
	CGIF2n  = 4
	CTCIF2n = 5
	CHTIF2n = 6
	CTEIF2n = 7
	CGIF3n  = 8
	CTCIF3n = 9
	CHTIF3n = 10
	CTEIF3n = 11
	CGIF4n  = 12
	CTCIF4n = 13
	CHTIF4n = 14
	CTEIF4n = 15
	CGIF5n  = 16
	CTCIF5n = 17
	CHTIF5n = 18
	CTEIF5n = 19
	CGIF6n  = 20
	CTCIF6n = 21
	CHTIF6n = 22
	CTEIF6n = 23
	CGIF7n  = 24
	CTCIF7n = 25
	CHTIF7n = 26
	CTEIF7n = 27
)

const (
	EN      CCR1 = 0x01 << 0  //+ Channel enable
	TCIE    CCR1 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR1 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR1 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR1 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR1 = 0x01 << 5  //+ Circular mode
	PINC    CCR1 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR1 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR1 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR1 = 0x03 << 10 //+ Memory size
	PL      CCR1 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR1 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR1 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR1 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR1 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	EN      CCR2 = 0x01 << 0  //+ Channel enable
	TCIE    CCR2 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR2 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR2 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR2 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR2 = 0x01 << 5  //+ Circular mode
	PINC    CCR2 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR2 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR2 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR2 = 0x03 << 10 //+ Memory size
	PL      CCR2 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR2 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR2 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR2 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR2 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	EN      CCR3 = 0x01 << 0  //+ Channel enable
	TCIE    CCR3 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR3 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR3 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR3 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR3 = 0x01 << 5  //+ Circular mode
	PINC    CCR3 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR3 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR3 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR3 = 0x03 << 10 //+ Memory size
	PL      CCR3 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR3 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR3 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR3 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR3 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	EN      CCR4 = 0x01 << 0  //+ Channel enable
	TCIE    CCR4 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR4 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR4 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR4 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR4 = 0x01 << 5  //+ Circular mode
	PINC    CCR4 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR4 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR4 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR4 = 0x03 << 10 //+ Memory size
	PL      CCR4 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR4 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR4 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR4 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR4 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	EN      CCR5 = 0x01 << 0  //+ Channel enable
	TCIE    CCR5 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR5 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR5 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR5 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR5 = 0x01 << 5  //+ Circular mode
	PINC    CCR5 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR5 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR5 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR5 = 0x03 << 10 //+ Memory size
	PL      CCR5 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR5 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR5 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR5 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR5 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	EN      CCR6 = 0x01 << 0  //+ Channel enable
	TCIE    CCR6 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR6 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR6 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR6 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR6 = 0x01 << 5  //+ Circular mode
	PINC    CCR6 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR6 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR6 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR6 = 0x03 << 10 //+ Memory size
	PL      CCR6 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR6 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR6 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR6 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR6 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	EN      CCR7 = 0x01 << 0  //+ Channel enable
	TCIE    CCR7 = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CCR7 = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CCR7 = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CCR7 = 0x01 << 4  //+ Data transfer direction
	CIRC    CCR7 = 0x01 << 5  //+ Circular mode
	PINC    CCR7 = 0x01 << 6  //+ Peripheral increment mode
	MINC    CCR7 = 0x01 << 7  //+ Memory increment mode
	PSIZE   CCR7 = 0x03 << 8  //+ Peripheral size
	MSIZE   CCR7 = 0x03 << 10 //+ Memory size
	PL      CCR7 = 0x03 << 12 //+ Channel priority level
	MEM2MEM CCR7 = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT CNDTR7 = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR7 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR7 = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	C1S CSELR = 0x0F << 0  //+ DMA channel 1 selection
	C2S CSELR = 0x0F << 4  //+ DMA channel 2 selection
	C3S CSELR = 0x0F << 8  //+ DMA channel 3 selection
	C4S CSELR = 0x0F << 12 //+ DMA channel 4 selection
	C5S CSELR = 0x0F << 16 //+ DMA channel 5 selection
	C6S CSELR = 0x0F << 20 //+ DMA channel 6 selection
	C7S CSELR = 0x0F << 24 //+ DMA channel 7 selection
)

const (
	C1Sn = 0
	C2Sn = 4
	C3Sn = 8
	C4Sn = 12
	C5Sn = 16
	C6Sn = 20
	C7Sn = 24
)
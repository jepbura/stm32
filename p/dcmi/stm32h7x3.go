// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package dcmi provides access to the registers of the DCMI peripheral.
//
// Instances:
//  DCMI  DCMI_BASE  -  DCMI  Digital camera interface
// Registers:
//  0x000 32  CR      control register 1
//  0x004 32  SR      status register
//  0x008 32  RIS     raw interrupt status register
//  0x00C 32  IER     interrupt enable register
//  0x010 32  MIS     masked interrupt status register
//  0x014 32  ICR     interrupt clear register
//  0x018 32  ESCR    embedded synchronization code register
//  0x01C 32  ESUR    embedded synchronization unmask register
//  0x020 32  CWSTRT  crop window start
//  0x024 32  CWSIZE  crop window size
//  0x028 32  DR      data register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package dcmi

const (
	CAPTURE CR = 0x01 << 0  //+ Capture enable
	CM      CR = 0x01 << 1  //+ Capture mode
	CROP    CR = 0x01 << 2  //+ Crop feature
	JPEG    CR = 0x01 << 3  //+ JPEG format
	ESS     CR = 0x01 << 4  //+ Embedded synchronization select
	PCKPOL  CR = 0x01 << 5  //+ Pixel clock polarity
	HSPOL   CR = 0x01 << 6  //+ Horizontal synchronization polarity
	VSPOL   CR = 0x01 << 7  //+ Vertical synchronization polarity
	FCRC    CR = 0x03 << 8  //+ Frame capture rate control
	EDM     CR = 0x03 << 10 //+ Extended data mode
	ENABLE  CR = 0x01 << 14 //+ DCMI enable
	BSM     CR = 0x03 << 16 //+ Byte Select mode
	OEBS    CR = 0x01 << 18 //+ Odd/Even Byte Select (Byte Select Start)
	LSM     CR = 0x01 << 19 //+ Line Select mode
	OELS    CR = 0x01 << 20 //+ Odd/Even Line Select (Line Select Start)
)

const (
	CAPTUREn = 0
	CMn      = 1
	CROPn    = 2
	JPEGn    = 3
	ESSn     = 4
	PCKPOLn  = 5
	HSPOLn   = 6
	VSPOLn   = 7
	FCRCn    = 8
	EDMn     = 10
	ENABLEn  = 14
	BSMn     = 16
	OEBSn    = 18
	LSMn     = 19
	OELSn    = 20
)

const (
	HSYNC SR = 0x01 << 0 //+ HSYNC
	VSYNC SR = 0x01 << 1 //+ VSYNC
	FNE   SR = 0x01 << 2 //+ FIFO not empty
)

const (
	HSYNCn = 0
	VSYNCn = 1
	FNEn   = 2
)

const (
	FRAME_RIS RIS = 0x01 << 0 //+ Capture complete raw interrupt status
	OVR_RIS   RIS = 0x01 << 1 //+ Overrun raw interrupt status
	ERR_RIS   RIS = 0x01 << 2 //+ Synchronization error raw interrupt status
	VSYNC_RIS RIS = 0x01 << 3 //+ VSYNC raw interrupt status
	LINE_RIS  RIS = 0x01 << 4 //+ Line raw interrupt status
)

const (
	FRAME_RISn = 0
	OVR_RISn   = 1
	ERR_RISn   = 2
	VSYNC_RISn = 3
	LINE_RISn  = 4
)

const (
	FRAME_IE IER = 0x01 << 0 //+ Capture complete interrupt enable
	OVR_IE   IER = 0x01 << 1 //+ Overrun interrupt enable
	ERR_IE   IER = 0x01 << 2 //+ Synchronization error interrupt enable
	VSYNC_IE IER = 0x01 << 3 //+ VSYNC interrupt enable
	LINE_IE  IER = 0x01 << 4 //+ Line interrupt enable
)

const (
	FRAME_IEn = 0
	OVR_IEn   = 1
	ERR_IEn   = 2
	VSYNC_IEn = 3
	LINE_IEn  = 4
)

const (
	FRAME_MIS MIS = 0x01 << 0 //+ Capture complete masked interrupt status
	OVR_MIS   MIS = 0x01 << 1 //+ Overrun masked interrupt status
	ERR_MIS   MIS = 0x01 << 2 //+ Synchronization error masked interrupt status
	VSYNC_MIS MIS = 0x01 << 3 //+ VSYNC masked interrupt status
	LINE_MIS  MIS = 0x01 << 4 //+ Line masked interrupt status
)

const (
	FRAME_MISn = 0
	OVR_MISn   = 1
	ERR_MISn   = 2
	VSYNC_MISn = 3
	LINE_MISn  = 4
)

const (
	FRAME_ISC ICR = 0x01 << 0 //+ Capture complete interrupt status clear
	OVR_ISC   ICR = 0x01 << 1 //+ Overrun interrupt status clear
	ERR_ISC   ICR = 0x01 << 2 //+ Synchronization error interrupt status clear
	VSYNC_ISC ICR = 0x01 << 3 //+ Vertical synch interrupt status clear
	LINE_ISC  ICR = 0x01 << 4 //+ line interrupt status clear
)

const (
	FRAME_ISCn = 0
	OVR_ISCn   = 1
	ERR_ISCn   = 2
	VSYNC_ISCn = 3
	LINE_ISCn  = 4
)

const (
	FSC ESCR = 0xFF << 0  //+ Frame start delimiter code
	LSC ESCR = 0xFF << 8  //+ Line start delimiter code
	LEC ESCR = 0xFF << 16 //+ Line end delimiter code
	FEC ESCR = 0xFF << 24 //+ Frame end delimiter code
)

const (
	FSCn = 0
	LSCn = 8
	LECn = 16
	FECn = 24
)

const (
	FSU ESUR = 0xFF << 0  //+ Frame start delimiter unmask
	LSU ESUR = 0xFF << 8  //+ Line start delimiter unmask
	LEU ESUR = 0xFF << 16 //+ Line end delimiter unmask
	FEU ESUR = 0xFF << 24 //+ Frame end delimiter unmask
)

const (
	FSUn = 0
	LSUn = 8
	LEUn = 16
	FEUn = 24
)

const (
	HOFFCNT CWSTRT = 0x3FFF << 0  //+ Horizontal offset count
	VST     CWSTRT = 0x1FFF << 16 //+ Vertical start line count
)

const (
	HOFFCNTn = 0
	VSTn     = 16
)

const (
	CAPCNT CWSIZE = 0x3FFF << 0  //+ Capture count
	VLINE  CWSIZE = 0x3FFF << 16 //+ Vertical line count
)

const (
	CAPCNTn = 0
	VLINEn  = 16
)

const (
	Byte0 DR = 0xFF << 0  //+ Data byte 0
	Byte1 DR = 0xFF << 8  //+ Data byte 1
	Byte2 DR = 0xFF << 16 //+ Data byte 2
	Byte3 DR = 0xFF << 24 //+ Data byte 3
)

const (
	Byte0n = 0
	Byte1n = 8
	Byte2n = 16
	Byte3n = 24
)

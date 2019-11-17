// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package ethernet_ptp provides access to the registers of the Ethernet_PTP peripheral.
//
// Instances:
//  Ethernet_PTP  Ethernet_PTP_BASE  -  -  Ethernet: Precision time protocol
// Registers:
//  0x000 32  PTPTSCR   Ethernet PTP time stamp control register
//  0x004 32  PTPSSIR   Ethernet PTP subsecond increment register
//  0x008 32  PTPTSHR   Ethernet PTP time stamp high register
//  0x00C 32  PTPTSLR   Ethernet PTP time stamp low register
//  0x010 32  PTPTSHUR  Ethernet PTP time stamp high update register
//  0x014 32  PTPTSLUR  Ethernet PTP time stamp low update register
//  0x018 32  PTPTSAR   Ethernet PTP time stamp addend register
//  0x01C 32  PTPTTHR   Ethernet PTP target time high register
//  0x020 32  PTPTTLR   Ethernet PTP target time low register
//  0x028 32  PTPTSSR   Ethernet PTP time stamp status register
//  0x02C 32  PTPPPSCR  Ethernet PTP PPS control register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package ethernet_ptp

const (
	TSE        PTPTSCR = 0x01 << 0  //+ Time stamp enable
	TSFCU      PTPTSCR = 0x01 << 1  //+ Time stamp fine or coarse update
	TSSTI      PTPTSCR = 0x01 << 2  //+ Time stamp system time initialize
	TSSTU      PTPTSCR = 0x01 << 3  //+ Time stamp system time update
	TSITE      PTPTSCR = 0x01 << 4  //+ Time stamp interrupt trigger enable
	TTSARU     PTPTSCR = 0x01 << 5  //+ Time stamp addend register update
	TSSARFE    PTPTSCR = 0x01 << 8  //+ Time stamp snapshot for all received frames enable
	TSSSR      PTPTSCR = 0x01 << 9  //+ Time stamp subsecond rollover: digital or binary rollover control
	TSPTPPSV2E PTPTSCR = 0x01 << 10 //+ Time stamp PTP packet snooping for version2 format enable
	TSSPTPOEFE PTPTSCR = 0x01 << 11 //+ Time stamp snapshot for PTP over ethernet frames enable
	TSSIPV6FE  PTPTSCR = 0x01 << 12 //+ Time stamp snapshot for IPv6 frames enable
	TSSIPV4FE  PTPTSCR = 0x01 << 13 //+ Time stamp snapshot for IPv4 frames enable
	TSSEME     PTPTSCR = 0x01 << 14 //+ Time stamp snapshot for event message enable
	TSSMRME    PTPTSCR = 0x01 << 15 //+ Time stamp snapshot for message relevant to master enable
	TSCNT      PTPTSCR = 0x03 << 16 //+ Time stamp clock node type
	TSPFFMAE   PTPTSCR = 0x01 << 18 //+ Time stamp PTP frame filtering MAC address enable
)

const (
	TSEn        = 0
	TSFCUn      = 1
	TSSTIn      = 2
	TSSTUn      = 3
	TSITEn      = 4
	TTSARUn     = 5
	TSSARFEn    = 8
	TSSSRn      = 9
	TSPTPPSV2En = 10
	TSSPTPOEFEn = 11
	TSSIPV6FEn  = 12
	TSSIPV4FEn  = 13
	TSSEMEn     = 14
	TSSMRMEn    = 15
	TSCNTn      = 16
	TSPFFMAEn   = 18
)

const (
	STSSI PTPSSIR = 0xFF << 0 //+ System time subsecond increment
)

const (
	STSSIn = 0
)

const (
	STS PTPTSHR = 0xFFFFFFFF << 0 //+ System time second
)

const (
	STSn = 0
)

const (
	STSS  PTPTSLR = 0x7FFFFFFF << 0 //+ System time subseconds
	STPNS PTPTSLR = 0x01 << 31      //+ System time positive or negative sign
)

const (
	STSSn  = 0
	STPNSn = 31
)

const (
	TSUS PTPTSHUR = 0xFFFFFFFF << 0 //+ Time stamp update second
)

const (
	TSUSn = 0
)

const (
	TSUSS  PTPTSLUR = 0x7FFFFFFF << 0 //+ Time stamp update subseconds
	TSUPNS PTPTSLUR = 0x01 << 31      //+ Time stamp update positive or negative sign
)

const (
	TSUSSn  = 0
	TSUPNSn = 31
)

const (
	TSA PTPTSAR = 0xFFFFFFFF << 0 //+ Time stamp addend
)

const (
	TSAn = 0
)

const (
	TTSH PTPTTHR = 0xFFFFFFFF << 0 //+ Target time stamp high
)

const (
	TTSHn = 0
)

const (
	TTSL PTPTTLR = 0xFFFFFFFF << 0 //+ Target time stamp low
)

const (
	TTSLn = 0
)

const (
	TSSO  PTPTSSR = 0x01 << 0 //+ Time stamp second overflow
	TSTTR PTPTSSR = 0x01 << 1 //+ Time stamp target time reached
)

const (
	TSSOn  = 0
	TSTTRn = 1
)

const (
	PPSFREQ PTPPPSCR = 0x0F << 0 //+ PPS frequency selection
)

const (
	PPSFREQn = 0
)
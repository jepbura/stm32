// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package ethernet_mac provides access to the registers of the Ethernet_MAC peripheral.
//
// Instances:
//  Ethernet_MAC  Ethernet_MAC_BASE  -  ETH,ETH_WKUP  Ethernet: media access control (MAC)
// Registers:
//  0x000 32  MACCR      Ethernet MAC configuration register
//  0x004 32  MACFFR     Ethernet MAC frame filter register
//  0x008 32  MACHTHR    Ethernet MAC hash table high register
//  0x00C 32  MACHTLR    Ethernet MAC hash table low register
//  0x010 32  MACMIIAR   Ethernet MAC MII address register
//  0x014 32  MACMIIDR   Ethernet MAC MII data register
//  0x018 32  MACFCR     Ethernet MAC flow control register
//  0x01C 32  MACVLANTR  Ethernet MAC VLAN tag register
//  0x02C 32  MACPMTCSR  Ethernet MAC PMT control and status register
//  0x034 32  MACDBGR    Ethernet MAC debug register
//  0x038 32  MACSR      Ethernet MAC interrupt status register
//  0x03C 32  MACIMR     Ethernet MAC interrupt mask register
//  0x040 32  MACA0HR    Ethernet MAC address 0 high register
//  0x044 32  MACA0LR    Ethernet MAC address 0 low register
//  0x048 32  MACA1HR    Ethernet MAC address 1 high register
//  0x04C 32  MACA1LR    Ethernet MAC address1 low register
//  0x050 32  MACA2HR    Ethernet MAC address 2 high register
//  0x054 32  MACA2LR    Ethernet MAC address 2 low register
//  0x058 32  MACA3HR    Ethernet MAC address 3 high register
//  0x05C 32  MACA3LR    Ethernet MAC address 3 low register
//  0x028 32  MACRWUFFR  Ethernet MAC remote wakeup frame filter register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package ethernet_mac

const (
	RE   MACCR = 0x01 << 2  //+ RE
	TE   MACCR = 0x01 << 3  //+ TE
	DC   MACCR = 0x01 << 4  //+ DC
	BL   MACCR = 0x03 << 5  //+ BL
	APCS MACCR = 0x01 << 7  //+ APCS
	RD   MACCR = 0x01 << 9  //+ RD
	IPCO MACCR = 0x01 << 10 //+ IPCO
	DM   MACCR = 0x01 << 11 //+ DM
	LM   MACCR = 0x01 << 12 //+ LM
	ROD  MACCR = 0x01 << 13 //+ ROD
	FES  MACCR = 0x01 << 14 //+ FES
	CSD  MACCR = 0x01 << 16 //+ CSD
	IFG  MACCR = 0x07 << 17 //+ IFG
	JD   MACCR = 0x01 << 22 //+ JD
	WD   MACCR = 0x01 << 23 //+ WD
	CSTF MACCR = 0x01 << 25 //+ CSTF
)

const (
	REn   = 2
	TEn   = 3
	DCn   = 4
	BLn   = 5
	APCSn = 7
	RDn   = 9
	IPCOn = 10
	DMn   = 11
	LMn   = 12
	RODn  = 13
	FESn  = 14
	CSDn  = 16
	IFGn  = 17
	JDn   = 22
	WDn   = 23
	CSTFn = 25
)

const (
	PM   MACFFR = 0x01 << 0  //+ Promiscuous mode
	HU   MACFFR = 0x01 << 1  //+ Hash unicast
	HM   MACFFR = 0x01 << 2  //+ Hash multicast
	DAIF MACFFR = 0x01 << 3  //+ Destination address inverse filtering
	PAM  MACFFR = 0x01 << 4  //+ Pass all multicast
	BFD  MACFFR = 0x01 << 5  //+ Broadcast frames disable
	PCF  MACFFR = 0x01 << 6  //+ Pass control frames
	SAIF MACFFR = 0x01 << 8  //+ Source address inverse filtering
	SAF  MACFFR = 0x01 << 9  //+ Source address filter
	HPF  MACFFR = 0x01 << 10 //+ Hash or perfect filter
	RA   MACFFR = 0x01 << 31 //+ Receive all
)

const (
	PMn   = 0
	HUn   = 1
	HMn   = 2
	DAIFn = 3
	PAMn  = 4
	BFDn  = 5
	PCFn  = 6
	SAIFn = 8
	SAFn  = 9
	HPFn  = 10
	RAn   = 31
)

const (
	HTH MACHTHR = 0xFFFFFFFF << 0 //+ Hash table high
)

const (
	HTHn = 0
)

const (
	HTL MACHTLR = 0xFFFFFFFF << 0 //+ Hash table low
)

const (
	HTLn = 0
)

const (
	MB MACMIIAR = 0x01 << 0  //+ MII busy
	MW MACMIIAR = 0x01 << 1  //+ MII write
	CR MACMIIAR = 0x07 << 2  //+ Clock range
	MR MACMIIAR = 0x1F << 6  //+ MII register
	PA MACMIIAR = 0x1F << 11 //+ PHY address
)

const (
	MBn = 0
	MWn = 1
	CRn = 2
	MRn = 6
	PAn = 11
)

const (
	MD MACMIIDR = 0xFFFF << 0 //+ MII data
)

const (
	MDn = 0
)

const (
	FCB  MACFCR = 0x01 << 0    //+ Flow control busy/back pressure activate
	TFCE MACFCR = 0x01 << 1    //+ Transmit flow control enable
	RFCE MACFCR = 0x01 << 2    //+ Receive flow control enable
	UPFD MACFCR = 0x01 << 3    //+ Unicast pause frame detect
	PLT  MACFCR = 0x03 << 4    //+ Pause low threshold
	ZQPD MACFCR = 0x01 << 7    //+ Zero-quanta pause disable
	PT   MACFCR = 0xFFFF << 16 //+ Pause time
)

const (
	FCBn  = 0
	TFCEn = 1
	RFCEn = 2
	UPFDn = 3
	PLTn  = 4
	ZQPDn = 7
	PTn   = 16
)

const (
	VLANTI MACVLANTR = 0xFFFF << 0 //+ VLAN tag identifier
	VLANTC MACVLANTR = 0x01 << 16  //+ 12-bit VLAN tag comparison
)

const (
	VLANTIn = 0
	VLANTCn = 16
)

const (
	PD     MACPMTCSR = 0x01 << 0  //+ Power down
	MPE    MACPMTCSR = 0x01 << 1  //+ Magic Packet enable
	WFE    MACPMTCSR = 0x01 << 2  //+ Wakeup frame enable
	MPR    MACPMTCSR = 0x01 << 5  //+ Magic packet received
	WFR    MACPMTCSR = 0x01 << 6  //+ Wakeup frame received
	GU     MACPMTCSR = 0x01 << 9  //+ Global unicast
	WFFRPR MACPMTCSR = 0x01 << 31 //+ Wakeup frame filter register pointer reset
)

const (
	PDn     = 0
	MPEn    = 1
	WFEn    = 2
	MPRn    = 5
	WFRn    = 6
	GUn     = 9
	WFFRPRn = 31
)

const (
	MMRPEA  MACDBGR = 0x01 << 0  //+ MAC MII receive protocol engine active
	MSFRWCS MACDBGR = 0x01 << 1  //+ MAC small FIFO read / write controllers status
	RFWRA   MACDBGR = 0x01 << 4  //+ Rx FIFO write controller active
	RFRCS   MACDBGR = 0x01 << 5  //+ Rx FIFO read controller status
	RFFL    MACDBGR = 0x01 << 8  //+ Rx FIFO fill level
	MMTEA   MACDBGR = 0x01 << 16 //+ MAC MII transmit engine active
	MTFCS   MACDBGR = 0x03 << 17 //+ MAC transmit frame controller status
	MTP     MACDBGR = 0x01 << 19 //+ MAC transmitter in pause
	TFRS    MACDBGR = 0x03 << 20 //+ Tx FIFO read status
	TFWA    MACDBGR = 0x01 << 22 //+ Tx FIFO write active
	TFNE    MACDBGR = 0x01 << 24 //+ Tx FIFO not empty
	TFF     MACDBGR = 0x01 << 25 //+ Tx FIFO full
)

const (
	MMRPEAn  = 0
	MSFRWCSn = 1
	RFWRAn   = 4
	RFRCSn   = 5
	RFFLn    = 8
	MMTEAn   = 16
	MTFCSn   = 17
	MTPn     = 19
	TFRSn    = 20
	TFWAn    = 22
	TFNEn    = 24
	TFFn     = 25
)

const (
	PMTS  MACSR = 0x01 << 3 //+ PMT status
	MMCS  MACSR = 0x01 << 4 //+ MMC status
	MMCRS MACSR = 0x01 << 5 //+ MMC receive status
	MMCTS MACSR = 0x01 << 6 //+ MMC transmit status
	TSTS  MACSR = 0x01 << 9 //+ Time stamp trigger status
)

const (
	PMTSn  = 3
	MMCSn  = 4
	MMCRSn = 5
	MMCTSn = 6
	TSTSn  = 9
)

const (
	PMTIM MACIMR = 0x01 << 3 //+ PMT interrupt mask
	TSTIM MACIMR = 0x01 << 9 //+ Time stamp trigger interrupt mask
)

const (
	PMTIMn = 3
	TSTIMn = 9
)

const (
	MACA0H MACA0HR = 0xFFFF << 0 //+ MAC address0 high
	MO     MACA0HR = 0x01 << 31  //+ MO
)

const (
	MACA0Hn = 0
	MOn     = 31
)

const (
	MACA0L MACA0LR = 0xFFFFFFFF << 0 //+ MAC address0 low
)

const (
	MACA0Ln = 0
)

const (
	MACA1H MACA1HR = 0xFFFF << 0 //+ MAC address1 high
	MBC    MACA1HR = 0x3F << 24  //+ Mask byte control
	SA     MACA1HR = 0x01 << 30  //+ Source address
	AE     MACA1HR = 0x01 << 31  //+ Address enable
)

const (
	MACA1Hn = 0
	MBCn    = 24
	SAn     = 30
	AEn     = 31
)

const (
	MACA1LR MACA1LR = 0xFFFFFFFF << 0 //+ MAC address1 low
)

const (
	MACA1LRn = 0
)

const (
	MAC2AH MACA2HR = 0xFFFF << 0 //+ MAC address2 high
	MBC    MACA2HR = 0x3F << 24  //+ Mask byte control
	SA     MACA2HR = 0x01 << 30  //+ Source address
	AE     MACA2HR = 0x01 << 31  //+ Address enable
)

const (
	MAC2AHn = 0
	MBCn    = 24
	SAn     = 30
	AEn     = 31
)

const (
	MACA2L MACA2LR = 0x7FFFFFFF << 0 //+ MAC address2 low
)

const (
	MACA2Ln = 0
)

const (
	MACA3H MACA3HR = 0xFFFF << 0 //+ MAC address3 high
	MBC    MACA3HR = 0x3F << 24  //+ Mask byte control
	SA     MACA3HR = 0x01 << 30  //+ Source address
	AE     MACA3HR = 0x01 << 31  //+ Address enable
)

const (
	MACA3Hn = 0
	MBCn    = 24
	SAn     = 30
	AEn     = 31
)

const (
	MBCA3L MACA3LR = 0xFFFFFFFF << 0 //+ MAC address3 low
)

const (
	MBCA3Ln = 0
)
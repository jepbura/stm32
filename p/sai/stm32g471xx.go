// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package sai provides access to the registers of the SAI peripheral.
//
// Instances:
//  SAI  SAI_BASE  -  SAI  Serial audio interface
// Registers:
//  0x004 32  ACR1    AConfiguration register 1
//  0x008 32  ACR2    AConfiguration register 2
//  0x00C 32  AFRCR   AFRCR
//  0x010 32  ASLOTR  ASlot register
//  0x014 32  AIM     AInterrupt mask register2
//  0x018 32  ASR     AStatus register
//  0x01C 32  ACLRFR  AClear flag register
//  0x020 32  ADR     AData register
//  0x024 32  BCR1    BConfiguration register 1
//  0x028 32  BCR2    BConfiguration register 2
//  0x02C 32  BFRCR   BFRCR
//  0x030 32  BSLOTR  BSlot register
//  0x034 32  BIM     BInterrupt mask register2
//  0x038 32  BSR     BStatus register
//  0x03C 32  BCLRFR  BClear flag register
//  0x040 32  BDR     BData register
//  0x044 32  PDMCR   PDM control register
//  0x048 32  PDMDLY  PDM delay register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package sai

const (
	MODE     ACR1 = 0x03 << 0  //+ Audio block mode
	PRTCFG   ACR1 = 0x03 << 2  //+ Protocol configuration
	DS       ACR1 = 0x07 << 5  //+ Data size
	LSBFIRST ACR1 = 0x01 << 8  //+ Least significant bit first
	CKSTR    ACR1 = 0x01 << 9  //+ Clock strobing edge
	SYNCEN   ACR1 = 0x03 << 10 //+ Synchronization enable
	MONO     ACR1 = 0x01 << 12 //+ Mono mode
	OutDri   ACR1 = 0x01 << 13 //+ Output drive
	SAIAEN   ACR1 = 0x01 << 16 //+ Audio block A enable
	DMAEN    ACR1 = 0x01 << 17 //+ DMA enable
	NODIV    ACR1 = 0x01 << 19 //+ No divider
	MCJDIV   ACR1 = 0x3F << 20 //+ Master clock divider
	OSR      ACR1 = 0x01 << 26 //+ OSR
	MCKEN    ACR1 = 0x01 << 27 //+ MCKEN
)

const (
	MODEn     = 0
	PRTCFGn   = 2
	DSn       = 5
	LSBFIRSTn = 8
	CKSTRn    = 9
	SYNCENn   = 10
	MONOn     = 12
	OutDrin   = 13
	SAIAENn   = 16
	DMAENn    = 17
	NODIVn    = 19
	MCJDIVn   = 20
	OSRn      = 26
	MCKENn    = 27
)

const (
	FTH     ACR2 = 0x07 << 0  //+ FIFO threshold
	FFLUS   ACR2 = 0x01 << 3  //+ FIFO flush
	TRIS    ACR2 = 0x01 << 4  //+ Tristate management on data line
	MUTE    ACR2 = 0x01 << 5  //+ Mute
	MUTEVAL ACR2 = 0x01 << 6  //+ Mute value
	MUTECN  ACR2 = 0x3F << 7  //+ Mute counter
	CPL     ACR2 = 0x01 << 13 //+ Complement bit
	COMP    ACR2 = 0x03 << 14 //+ Companding mode
)

const (
	FTHn     = 0
	FFLUSn   = 3
	TRISn    = 4
	MUTEn    = 5
	MUTEVALn = 6
	MUTECNn  = 7
	CPLn     = 13
	COMPn    = 14
)

const (
	FRL   AFRCR = 0xFF << 0  //+ Frame length
	FSALL AFRCR = 0x7F << 8  //+ Frame synchronization active level length
	FSDEF AFRCR = 0x01 << 16 //+ Frame synchronization definition
	FSPOL AFRCR = 0x01 << 17 //+ Frame synchronization polarity
	FSOFF AFRCR = 0x01 << 18 //+ Frame synchronization offset
)

const (
	FRLn   = 0
	FSALLn = 8
	FSDEFn = 16
	FSPOLn = 17
	FSOFFn = 18
)

const (
	FBOFF  ASLOTR = 0x1F << 0    //+ First bit offset
	SLOTSZ ASLOTR = 0x03 << 6    //+ Slot size
	NBSLOT ASLOTR = 0x0F << 8    //+ Number of slots in an audio frame
	SLOTEN ASLOTR = 0xFFFF << 16 //+ Slot enable
)

const (
	FBOFFn  = 0
	SLOTSZn = 6
	NBSLOTn = 8
	SLOTENn = 16
)

const (
	OVRUDRIE AIM = 0x01 << 0 //+ Overrun/underrun interrupt enable
	MUTEDET  AIM = 0x01 << 1 //+ Mute detection interrupt enable
	WCKCFG   AIM = 0x01 << 2 //+ Wrong clock configuration interrupt enable
	FREQIE   AIM = 0x01 << 3 //+ FIFO request interrupt enable
	CNRDYIE  AIM = 0x01 << 4 //+ Codec not ready interrupt enable
	AFSDETIE AIM = 0x01 << 5 //+ Anticipated frame synchronization detection interrupt enable
	LFSDET   AIM = 0x01 << 6 //+ Late frame synchronization detection interrupt enable
)

const (
	OVRUDRIEn = 0
	MUTEDETn  = 1
	WCKCFGn   = 2
	FREQIEn   = 3
	CNRDYIEn  = 4
	AFSDETIEn = 5
	LFSDETn   = 6
)

const (
	OVRUDR  ASR = 0x01 << 0  //+ Overrun / underrun
	MUTEDET ASR = 0x01 << 1  //+ Mute detection
	WCKCFG  ASR = 0x01 << 2  //+ Wrong clock configuration flag. This bit is read only
	FREQ    ASR = 0x01 << 3  //+ FIFO request
	CNRDY   ASR = 0x01 << 4  //+ Codec not ready
	AFSDET  ASR = 0x01 << 5  //+ Anticipated frame synchronization detection
	LFSDET  ASR = 0x01 << 6  //+ Late frame synchronization detection
	FLVL    ASR = 0x07 << 16 //+ FIFO level threshold
)

const (
	OVRUDRn  = 0
	MUTEDETn = 1
	WCKCFGn  = 2
	FREQn    = 3
	CNRDYn   = 4
	AFSDETn  = 5
	LFSDETn  = 6
	FLVLn    = 16
)

const (
	OVRUDR  ACLRFR = 0x01 << 0 //+ Clear overrun / underrun
	MUTEDET ACLRFR = 0x01 << 1 //+ Mute detection flag
	WCKCFG  ACLRFR = 0x01 << 2 //+ Clear wrong clock configuration flag
	CNRDY   ACLRFR = 0x01 << 4 //+ Clear codec not ready flag
	CAFSDET ACLRFR = 0x01 << 5 //+ Clear anticipated frame synchronization detection flag
	LFSDET  ACLRFR = 0x01 << 6 //+ Clear late frame synchronization detection flag
)

const (
	OVRUDRn  = 0
	MUTEDETn = 1
	WCKCFGn  = 2
	CNRDYn   = 4
	CAFSDETn = 5
	LFSDETn  = 6
)

const (
	DATA ADR = 0xFFFFFFFF << 0 //+ Data
)

const (
	DATAn = 0
)

const (
	MODE     BCR1 = 0x03 << 0  //+ Audio block mode
	PRTCFG   BCR1 = 0x03 << 2  //+ Protocol configuration
	DS       BCR1 = 0x07 << 5  //+ Data size
	LSBFIRST BCR1 = 0x01 << 8  //+ Least significant bit first
	CKSTR    BCR1 = 0x01 << 9  //+ Clock strobing edge
	SYNCEN   BCR1 = 0x03 << 10 //+ Synchronization enable
	MONO     BCR1 = 0x01 << 12 //+ Mono mode
	OutDri   BCR1 = 0x01 << 13 //+ Output drive
	SAIBEN   BCR1 = 0x01 << 16 //+ Audio block B enable
	DMAEN    BCR1 = 0x01 << 17 //+ DMA enable
	NODIV    BCR1 = 0x01 << 19 //+ No divider
	MCJDIV   BCR1 = 0x3F << 20 //+ Master clock divider
	OSR      BCR1 = 0x01 << 26 //+ OSR
	MCKEN    BCR1 = 0x01 << 27 //+ MCKEN
)

const (
	MODEn     = 0
	PRTCFGn   = 2
	DSn       = 5
	LSBFIRSTn = 8
	CKSTRn    = 9
	SYNCENn   = 10
	MONOn     = 12
	OutDrin   = 13
	SAIBENn   = 16
	DMAENn    = 17
	NODIVn    = 19
	MCJDIVn   = 20
	OSRn      = 26
	MCKENn    = 27
)

const (
	FTH     BCR2 = 0x07 << 0  //+ FIFO threshold
	FFLUS   BCR2 = 0x01 << 3  //+ FIFO flush
	TRIS    BCR2 = 0x01 << 4  //+ Tristate management on data line
	MUTE    BCR2 = 0x01 << 5  //+ Mute
	MUTEVAL BCR2 = 0x01 << 6  //+ Mute value
	MUTECN  BCR2 = 0x3F << 7  //+ Mute counter
	CPL     BCR2 = 0x01 << 13 //+ Complement bit
	COMP    BCR2 = 0x03 << 14 //+ Companding mode
)

const (
	FTHn     = 0
	FFLUSn   = 3
	TRISn    = 4
	MUTEn    = 5
	MUTEVALn = 6
	MUTECNn  = 7
	CPLn     = 13
	COMPn    = 14
)

const (
	FRL   BFRCR = 0xFF << 0  //+ Frame length
	FSALL BFRCR = 0x7F << 8  //+ Frame synchronization active level length
	FSDEF BFRCR = 0x01 << 16 //+ Frame synchronization definition
	FSPOL BFRCR = 0x01 << 17 //+ Frame synchronization polarity
	FSOFF BFRCR = 0x01 << 18 //+ Frame synchronization offset
)

const (
	FRLn   = 0
	FSALLn = 8
	FSDEFn = 16
	FSPOLn = 17
	FSOFFn = 18
)

const (
	FBOFF  BSLOTR = 0x1F << 0    //+ First bit offset
	SLOTSZ BSLOTR = 0x03 << 6    //+ Slot size
	NBSLOT BSLOTR = 0x0F << 8    //+ Number of slots in an audio frame
	SLOTEN BSLOTR = 0xFFFF << 16 //+ Slot enable
)

const (
	FBOFFn  = 0
	SLOTSZn = 6
	NBSLOTn = 8
	SLOTENn = 16
)

const (
	OVRUDRIE BIM = 0x01 << 0 //+ Overrun/underrun interrupt enable
	MUTEDET  BIM = 0x01 << 1 //+ Mute detection interrupt enable
	WCKCFG   BIM = 0x01 << 2 //+ Wrong clock configuration interrupt enable
	FREQIE   BIM = 0x01 << 3 //+ FIFO request interrupt enable
	CNRDYIE  BIM = 0x01 << 4 //+ Codec not ready interrupt enable
	AFSDETIE BIM = 0x01 << 5 //+ Anticipated frame synchronization detection interrupt enable
	LFSDETIE BIM = 0x01 << 6 //+ Late frame synchronization detection interrupt enable
)

const (
	OVRUDRIEn = 0
	MUTEDETn  = 1
	WCKCFGn   = 2
	FREQIEn   = 3
	CNRDYIEn  = 4
	AFSDETIEn = 5
	LFSDETIEn = 6
)

const (
	OVRUDR  BSR = 0x01 << 0  //+ Overrun / underrun
	MUTEDET BSR = 0x01 << 1  //+ Mute detection
	WCKCFG  BSR = 0x01 << 2  //+ Wrong clock configuration flag
	FREQ    BSR = 0x01 << 3  //+ FIFO request
	CNRDY   BSR = 0x01 << 4  //+ Codec not ready
	AFSDET  BSR = 0x01 << 5  //+ Anticipated frame synchronization detection
	LFSDET  BSR = 0x01 << 6  //+ Late frame synchronization detection
	FLVL    BSR = 0x07 << 16 //+ FIFO level threshold
)

const (
	OVRUDRn  = 0
	MUTEDETn = 1
	WCKCFGn  = 2
	FREQn    = 3
	CNRDYn   = 4
	AFSDETn  = 5
	LFSDETn  = 6
	FLVLn    = 16
)

const (
	OVRUDR  BCLRFR = 0x01 << 0 //+ Clear overrun / underrun
	MUTEDET BCLRFR = 0x01 << 1 //+ Mute detection flag
	WCKCFG  BCLRFR = 0x01 << 2 //+ Clear wrong clock configuration flag
	CNRDY   BCLRFR = 0x01 << 4 //+ Clear codec not ready flag
	CAFSDET BCLRFR = 0x01 << 5 //+ Clear anticipated frame synchronization detection flag
	LFSDET  BCLRFR = 0x01 << 6 //+ Clear late frame synchronization detection flag
)

const (
	OVRUDRn  = 0
	MUTEDETn = 1
	WCKCFGn  = 2
	CNRDYn   = 4
	CAFSDETn = 5
	LFSDETn  = 6
)

const (
	DATA BDR = 0xFFFFFFFF << 0 //+ Data
)

const (
	DATAn = 0
)

const (
	PDMEN  PDMCR = 0x01 << 0  //+ PDMEN
	MICNBR PDMCR = 0x03 << 4  //+ MICNBR
	CKEN1  PDMCR = 0x01 << 8  //+ CKEN1
	CKEN2  PDMCR = 0x01 << 9  //+ CKEN2
	CKEN3  PDMCR = 0x01 << 10 //+ CKEN3
	CKEN4  PDMCR = 0x01 << 11 //+ CKEN4
)

const (
	PDMENn  = 0
	MICNBRn = 4
	CKEN1n  = 8
	CKEN2n  = 9
	CKEN3n  = 10
	CKEN4n  = 11
)

const (
	DLYM1L PDMDLY = 0x07 << 0  //+ DLYM1L
	DLYM1R PDMDLY = 0x07 << 4  //+ DLYM1R
	DLYM2L PDMDLY = 0x07 << 8  //+ DLYM2L
	DLYM2R PDMDLY = 0x07 << 12 //+ DLYM2R
	DLYM3L PDMDLY = 0x07 << 16 //+ DLYM3L
	DLYM3R PDMDLY = 0x07 << 20 //+ DLYM3R
	DLYM4L PDMDLY = 0x07 << 24 //+ DLYM4L
	DLYM4R PDMDLY = 0x07 << 28 //+ DLYM4R
)

const (
	DLYM1Ln = 0
	DLYM1Rn = 4
	DLYM2Ln = 8
	DLYM2Rn = 12
	DLYM3Ln = 16
	DLYM3Rn = 20
	DLYM4Ln = 24
	DLYM4Rn = 28
)

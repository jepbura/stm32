// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f215

package flash

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	ACR     RACR
	KEYR    RKEYR
	OPTKEYR ROPTKEYR
	SR      RSR
	CR      RCR
	OPTCR   ROPTCR
}

func FLASH() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.FLASH_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type ACR uint32

type RACR struct{ mmio.U32 }

func (r *RACR) LoadBits(mask ACR) ACR { return ACR(r.U32.LoadBits(uint32(mask))) }
func (r *RACR) StoreBits(mask, b ACR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RACR) SetBits(mask ACR)      { r.U32.SetBits(uint32(mask)) }
func (r *RACR) ClearBits(mask ACR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RACR) Load() ACR             { return ACR(r.U32.Load()) }
func (r *RACR) Store(b ACR)           { r.U32.Store(uint32(b)) }

type RMACR struct{ mmio.UM32 }

func (rm RMACR) Load() ACR   { return ACR(rm.UM32.Load()) }
func (rm RMACR) Store(b ACR) { rm.UM32.Store(uint32(b)) }

func LATENCY_(p *Periph) RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(LATENCY)}}
}

func PRFTEN_(p *Periph) RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(PRFTEN)}}
}

func ICEN_(p *Periph) RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(ICEN)}}
}

func DCEN_(p *Periph) RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(DCEN)}}
}

func ICRST_(p *Periph) RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(ICRST)}}
}

func DCRST_(p *Periph) RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(DCRST)}}
}

type KEYR uint32

type RKEYR struct{ mmio.U32 }

func (r *RKEYR) LoadBits(mask KEYR) KEYR { return KEYR(r.U32.LoadBits(uint32(mask))) }
func (r *RKEYR) StoreBits(mask, b KEYR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RKEYR) SetBits(mask KEYR)       { r.U32.SetBits(uint32(mask)) }
func (r *RKEYR) ClearBits(mask KEYR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RKEYR) Load() KEYR              { return KEYR(r.U32.Load()) }
func (r *RKEYR) Store(b KEYR)            { r.U32.Store(uint32(b)) }

type RMKEYR struct{ mmio.UM32 }

func (rm RMKEYR) Load() KEYR   { return KEYR(rm.UM32.Load()) }
func (rm RMKEYR) Store(b KEYR) { rm.UM32.Store(uint32(b)) }

type OPTKEYR uint32

type ROPTKEYR struct{ mmio.U32 }

func (r *ROPTKEYR) LoadBits(mask OPTKEYR) OPTKEYR { return OPTKEYR(r.U32.LoadBits(uint32(mask))) }
func (r *ROPTKEYR) StoreBits(mask, b OPTKEYR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPTKEYR) SetBits(mask OPTKEYR)          { r.U32.SetBits(uint32(mask)) }
func (r *ROPTKEYR) ClearBits(mask OPTKEYR)        { r.U32.ClearBits(uint32(mask)) }
func (r *ROPTKEYR) Load() OPTKEYR                 { return OPTKEYR(r.U32.Load()) }
func (r *ROPTKEYR) Store(b OPTKEYR)               { r.U32.Store(uint32(b)) }

type RMOPTKEYR struct{ mmio.UM32 }

func (rm RMOPTKEYR) Load() OPTKEYR   { return OPTKEYR(rm.UM32.Load()) }
func (rm RMOPTKEYR) Store(b OPTKEYR) { rm.UM32.Store(uint32(b)) }

type SR uint32

type RSR struct{ mmio.U32 }

func (r *RSR) LoadBits(mask SR) SR  { return SR(r.U32.LoadBits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func EOP_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

func OPERR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OPERR)}}
}

func WRPERR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func PGAERR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGAERR)}}
}

func PGPERR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGPERR)}}
}

func PGSERR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGSERR)}}
}

func BSY_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

type CR uint32

type RCR struct{ mmio.U32 }

func (r *RCR) LoadBits(mask CR) CR  { return CR(r.U32.LoadBits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func PG_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PG)}}
}

func SER_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SER)}}
}

func MER_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MER)}}
}

func SNB_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SNB)}}
}

func PSIZE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PSIZE)}}
}

func STRT_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(STRT)}}
}

func EOPIE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EOPIE)}}
}

func ERRIE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ERRIE)}}
}

func LOCK_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LOCK)}}
}

type OPTCR uint32

type ROPTCR struct{ mmio.U32 }

func (r *ROPTCR) LoadBits(mask OPTCR) OPTCR { return OPTCR(r.U32.LoadBits(uint32(mask))) }
func (r *ROPTCR) StoreBits(mask, b OPTCR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPTCR) SetBits(mask OPTCR)        { r.U32.SetBits(uint32(mask)) }
func (r *ROPTCR) ClearBits(mask OPTCR)      { r.U32.ClearBits(uint32(mask)) }
func (r *ROPTCR) Load() OPTCR               { return OPTCR(r.U32.Load()) }
func (r *ROPTCR) Store(b OPTCR)             { r.U32.Store(uint32(b)) }

type RMOPTCR struct{ mmio.UM32 }

func (rm RMOPTCR) Load() OPTCR   { return OPTCR(rm.UM32.Load()) }
func (rm RMOPTCR) Store(b OPTCR) { rm.UM32.Store(uint32(b)) }

func OPTLOCK_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(OPTLOCK)}}
}

func OPTSTRT_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(OPTSTRT)}}
}

func BOR_LEV_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(BOR_LEV)}}
}

func WDG_SW_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(WDG_SW)}}
}

func nRST_STOP_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(nRST_STOP)}}
}

func nRST_STDBY_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(nRST_STDBY)}}
}

func RDP_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(RDP)}}
}

func nWRP_(p *Periph) RMOPTCR {
	return RMOPTCR{mmio.UM32{&p.OPTCR.U32, uint32(nWRP)}}
}

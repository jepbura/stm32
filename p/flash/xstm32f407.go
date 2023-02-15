// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f407

package flash

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	ACR     mmio.R32[ACR]
	KEYR    mmio.R32[uint32]
	OPTKEYR mmio.R32[uint32]
	SR      mmio.R32[SR]
	CR      mmio.R32[CR]
	OPTCR   mmio.R32[OPTCR]
}

func FLASH() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.FLASH_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type ACR uint32

func LATENCY_(p *Periph) mmio.RM32[ACR] { return mmio.RM32[ACR]{&p.ACR, LATENCY} }
func PRFTEN_(p *Periph) mmio.RM32[ACR]  { return mmio.RM32[ACR]{&p.ACR, PRFTEN} }
func ICEN_(p *Periph) mmio.RM32[ACR]    { return mmio.RM32[ACR]{&p.ACR, ICEN} }
func DCEN_(p *Periph) mmio.RM32[ACR]    { return mmio.RM32[ACR]{&p.ACR, DCEN} }
func ICRST_(p *Periph) mmio.RM32[ACR]   { return mmio.RM32[ACR]{&p.ACR, ICRST} }
func DCRST_(p *Periph) mmio.RM32[ACR]   { return mmio.RM32[ACR]{&p.ACR, DCRST} }

type SR uint32

func EOP_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, EOP} }
func OPERR_(p *Periph) mmio.RM32[SR]  { return mmio.RM32[SR]{&p.SR, OPERR} }
func WRPERR_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, WRPERR} }
func PGAERR_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, PGAERR} }
func PGPERR_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, PGPERR} }
func PGSERR_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, PGSERR} }
func BSY_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, BSY} }

type CR uint32

func PG_(p *Periph) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.CR, PG} }
func SER_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, SER} }
func MER_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, MER} }
func SNB_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, SNB} }
func PSIZE_(p *Periph) mmio.RM32[CR] { return mmio.RM32[CR]{&p.CR, PSIZE} }
func STRT_(p *Periph) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.CR, STRT} }
func EOPIE_(p *Periph) mmio.RM32[CR] { return mmio.RM32[CR]{&p.CR, EOPIE} }
func ERRIE_(p *Periph) mmio.RM32[CR] { return mmio.RM32[CR]{&p.CR, ERRIE} }
func LOCK_(p *Periph) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.CR, LOCK} }

type OPTCR uint32

func OPTLOCK_(p *Periph) mmio.RM32[OPTCR]    { return mmio.RM32[OPTCR]{&p.OPTCR, OPTLOCK} }
func OPTSTRT_(p *Periph) mmio.RM32[OPTCR]    { return mmio.RM32[OPTCR]{&p.OPTCR, OPTSTRT} }
func BOR_LEV_(p *Periph) mmio.RM32[OPTCR]    { return mmio.RM32[OPTCR]{&p.OPTCR, BOR_LEV} }
func WDG_SW_(p *Periph) mmio.RM32[OPTCR]     { return mmio.RM32[OPTCR]{&p.OPTCR, WDG_SW} }
func nRST_STOP_(p *Periph) mmio.RM32[OPTCR]  { return mmio.RM32[OPTCR]{&p.OPTCR, nRST_STOP} }
func nRST_STDBY_(p *Periph) mmio.RM32[OPTCR] { return mmio.RM32[OPTCR]{&p.OPTCR, nRST_STDBY} }
func RDP_(p *Periph) mmio.RM32[OPTCR]        { return mmio.RM32[OPTCR]{&p.OPTCR, RDP} }
func nWRP_(p *Periph) mmio.RM32[OPTCR]       { return mmio.RM32[OPTCR]{&p.OPTCR, nWRP} }

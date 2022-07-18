// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32l4x6

package pwr

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1  RCR1
	CR2  RCR2
	CR3  RCR3
	CR4  RCR4
	SR1  RSR1
	SR2  RSR2
	SCR  RSCR
	_    uint32
	PUDC [8]RPUDC
}

func PWR() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB1
}

type CR1 uint32

type RCR1 struct{ mmio.U32 }

func (r *RCR1) LoadBits(mask CR1) CR1 { return CR1(r.U32.LoadBits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

func LPMS_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LPMS)}}
}

func DBP_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DBP)}}
}

func VOS_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(VOS)}}
}

func LPR_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LPR)}}
}

type CR2 uint32

type RCR2 struct{ mmio.U32 }

func (r *RCR2) LoadBits(mask CR2) CR2 { return CR2(r.U32.LoadBits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

func PVDE_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVDE)}}
}

func PLS_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PLS)}}
}

func PVME1_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVME1)}}
}

func PVME2_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVME2)}}
}

func PVME3_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVME3)}}
}

func PVME4_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PVME4)}}
}

func IOSV_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(IOSV)}}
}

func USV_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(USV)}}
}

type CR3 uint32

type RCR3 struct{ mmio.U32 }

func (r *RCR3) LoadBits(mask CR3) CR3 { return CR3(r.U32.LoadBits(uint32(mask))) }
func (r *RCR3) StoreBits(mask, b CR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR3) SetBits(mask CR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR3) ClearBits(mask CR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR3) Load() CR3             { return CR3(r.U32.Load()) }
func (r *RCR3) Store(b CR3)           { r.U32.Store(uint32(b)) }

type RMCR3 struct{ mmio.UM32 }

func (rm RMCR3) Load() CR3   { return CR3(rm.UM32.Load()) }
func (rm RMCR3) Store(b CR3) { rm.UM32.Store(uint32(b)) }

func EWUP1_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP1)}}
}

func EWUP2_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP2)}}
}

func EWUP3_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP3)}}
}

func EWUP4_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP4)}}
}

func EWUP5_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWUP5)}}
}

func RRS_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(RRS)}}
}

func APC_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(APC)}}
}

func EWF_(p *Periph) RMCR3 {
	return RMCR3{mmio.UM32{&p.CR3.U32, uint32(EWF)}}
}

type CR4 uint32

type RCR4 struct{ mmio.U32 }

func (r *RCR4) LoadBits(mask CR4) CR4 { return CR4(r.U32.LoadBits(uint32(mask))) }
func (r *RCR4) StoreBits(mask, b CR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR4) SetBits(mask CR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR4) ClearBits(mask CR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR4) Load() CR4             { return CR4(r.U32.Load()) }
func (r *RCR4) Store(b CR4)           { r.U32.Store(uint32(b)) }

type RMCR4 struct{ mmio.UM32 }

func (rm RMCR4) Load() CR4   { return CR4(rm.UM32.Load()) }
func (rm RMCR4) Store(b CR4) { rm.UM32.Store(uint32(b)) }

func WP1_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP1)}}
}

func WP2_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP2)}}
}

func WP3_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP3)}}
}

func WP4_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP4)}}
}

func WP5_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(WP5)}}
}

func VBE_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(VBE)}}
}

func VBRS_(p *Periph) RMCR4 {
	return RMCR4{mmio.UM32{&p.CR4.U32, uint32(VBRS)}}
}

type SR1 uint32

type RSR1 struct{ mmio.U32 }

func (r *RSR1) LoadBits(mask SR1) SR1 { return SR1(r.U32.LoadBits(uint32(mask))) }
func (r *RSR1) StoreBits(mask, b SR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR1) SetBits(mask SR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR1) ClearBits(mask SR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR1) Load() SR1             { return SR1(r.U32.Load()) }
func (r *RSR1) Store(b SR1)           { r.U32.Store(uint32(b)) }

type RMSR1 struct{ mmio.UM32 }

func (rm RMSR1) Load() SR1   { return SR1(rm.UM32.Load()) }
func (rm RMSR1) Store(b SR1) { rm.UM32.Store(uint32(b)) }

func CWUF1_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(CWUF1)}}
}

func CWUF2_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(CWUF2)}}
}

func CWUF3_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(CWUF3)}}
}

func CWUF4_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(CWUF4)}}
}

func CWUF5_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(CWUF5)}}
}

func CSBF_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(CSBF)}}
}

func WUFI_(p *Periph) RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(WUFI)}}
}

type SR2 uint32

type RSR2 struct{ mmio.U32 }

func (r *RSR2) LoadBits(mask SR2) SR2 { return SR2(r.U32.LoadBits(uint32(mask))) }
func (r *RSR2) StoreBits(mask, b SR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR2) SetBits(mask SR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR2) ClearBits(mask SR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR2) Load() SR2             { return SR2(r.U32.Load()) }
func (r *RSR2) Store(b SR2)           { r.U32.Store(uint32(b)) }

type RMSR2 struct{ mmio.UM32 }

func (rm RMSR2) Load() SR2   { return SR2(rm.UM32.Load()) }
func (rm RMSR2) Store(b SR2) { rm.UM32.Store(uint32(b)) }

func REGLPS_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(REGLPS)}}
}

func REGLPF_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(REGLPF)}}
}

func VOSF_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(VOSF)}}
}

func PVDO_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVDO)}}
}

func PVMO1_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO1)}}
}

func PVMO2_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO2)}}
}

func PVMO3_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO3)}}
}

func PVMO4_(p *Periph) RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PVMO4)}}
}

type SCR uint32

type RSCR struct{ mmio.U32 }

func (r *RSCR) LoadBits(mask SCR) SCR { return SCR(r.U32.LoadBits(uint32(mask))) }
func (r *RSCR) StoreBits(mask, b SCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSCR) SetBits(mask SCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSCR) ClearBits(mask SCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSCR) Load() SCR             { return SCR(r.U32.Load()) }
func (r *RSCR) Store(b SCR)           { r.U32.Store(uint32(b)) }

type RMSCR struct{ mmio.UM32 }

func (rm RMSCR) Load() SCR   { return SCR(rm.UM32.Load()) }
func (rm RMSCR) Store(b SCR) { rm.UM32.Store(uint32(b)) }

func WUF1_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(WUF1)}}
}

func WUF2_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(WUF2)}}
}

func WUF3_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(WUF3)}}
}

func WUF4_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(WUF4)}}
}

func WUF5_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(WUF5)}}
}

func SBF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(SBF)}}
}

type RPUDC struct {
	PUCR RPUCR
	PDCR RPDCR
}

type PUCR uint32

type RPUCR struct{ mmio.U32 }

func (r *RPUCR) LoadBits(mask PUCR) PUCR { return PUCR(r.U32.LoadBits(uint32(mask))) }
func (r *RPUCR) StoreBits(mask, b PUCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPUCR) SetBits(mask PUCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RPUCR) ClearBits(mask PUCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPUCR) Load() PUCR              { return PUCR(r.U32.Load()) }
func (r *RPUCR) Store(b PUCR)            { r.U32.Store(uint32(b)) }

type RMPUCR struct{ mmio.UM32 }

func (rm RMPUCR) Load() PUCR   { return PUCR(rm.UM32.Load()) }
func (rm RMPUCR) Store(b PUCR) { rm.UM32.Store(uint32(b)) }

func PU0_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU0)}}
}

func PU1_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU1)}}
}

func PU2_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU2)}}
}

func PU3_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU3)}}
}

func PU4_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU4)}}
}

func PU5_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU5)}}
}

func PU6_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU6)}}
}

func PU7_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU7)}}
}

func PU8_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU8)}}
}

func PU9_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU9)}}
}

func PU10_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU10)}}
}

func PU11_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU11)}}
}

func PU12_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU12)}}
}

func PU13_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU13)}}
}

func PU14_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU14)}}
}

func PU15_(p *Periph, n int) RMPUCR {
	return RMPUCR{mmio.UM32{&p.PUDC[n].PUCR.U32, uint32(PU15)}}
}

type PDCR uint32

type RPDCR struct{ mmio.U32 }

func (r *RPDCR) LoadBits(mask PDCR) PDCR { return PDCR(r.U32.LoadBits(uint32(mask))) }
func (r *RPDCR) StoreBits(mask, b PDCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPDCR) SetBits(mask PDCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RPDCR) ClearBits(mask PDCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPDCR) Load() PDCR              { return PDCR(r.U32.Load()) }
func (r *RPDCR) Store(b PDCR)            { r.U32.Store(uint32(b)) }

type RMPDCR struct{ mmio.UM32 }

func (rm RMPDCR) Load() PDCR   { return PDCR(rm.UM32.Load()) }
func (rm RMPDCR) Store(b PDCR) { rm.UM32.Store(uint32(b)) }

func PD0_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD0)}}
}

func PD1_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD1)}}
}

func PD2_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD2)}}
}

func PD3_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD3)}}
}

func PD4_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD4)}}
}

func PD5_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD5)}}
}

func PD6_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD6)}}
}

func PD7_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD7)}}
}

func PD8_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD8)}}
}

func PD9_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD9)}}
}

func PD10_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD10)}}
}

func PD11_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD11)}}
}

func PD12_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD12)}}
}

func PD13_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD13)}}
}

func PD14_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD14)}}
}

func PD15_(p *Periph, n int) RMPDCR {
	return RMPDCR{mmio.UM32{&p.PUDC[n].PDCR.U32, uint32(PD15)}}
}

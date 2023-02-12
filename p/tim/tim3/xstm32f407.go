// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f407

package tim3

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1   RCR1
	CR2   RCR2
	SMCR  RSMCR
	DIER  RDIER
	SR    RSR
	EGR   REGR
	CCMR1 RCCMR1
	CCMR2 RCCMR2
	CCER  RCCER
	CNT   RCNT
	PSC   RPSC
	ARR   RARR
	_     uint32
	CCR1  RCCR
	CCR2  RCCR
	CCR3  RCCR
	CCR4  RCCR
	_     uint32
	DCR   RDCR
	DMAR  RDMAR
}

func TIM3() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM3_BASE))) }
func TIM4() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM4_BASE))) }

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

func CEN_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CEN)}}
}

func UDIS_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(UDIS)}}
}

func URS_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(URS)}}
}

func OPM_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(OPM)}}
}

func DIR_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DIR)}}
}

func CMS_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CMS)}}
}

func ARPE_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ARPE)}}
}

func CKD_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CKD)}}
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

func CCDS_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(CCDS)}}
}

func MMS_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(MMS)}}
}

func TI1S_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TI1S)}}
}

type SMCR uint32

type RSMCR struct{ mmio.U32 }

func (r *RSMCR) LoadBits(mask SMCR) SMCR { return SMCR(r.U32.LoadBits(uint32(mask))) }
func (r *RSMCR) StoreBits(mask, b SMCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSMCR) SetBits(mask SMCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RSMCR) ClearBits(mask SMCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RSMCR) Load() SMCR              { return SMCR(r.U32.Load()) }
func (r *RSMCR) Store(b SMCR)            { r.U32.Store(uint32(b)) }

type RMSMCR struct{ mmio.UM32 }

func (rm RMSMCR) Load() SMCR   { return SMCR(rm.UM32.Load()) }
func (rm RMSMCR) Store(b SMCR) { rm.UM32.Store(uint32(b)) }

func SMS_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(SMS)}}
}

func TS_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(TS)}}
}

func MSM_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(MSM)}}
}

func ETF_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(ETF)}}
}

func ETPS_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(ETPS)}}
}

func ECE_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(ECE)}}
}

func ETP_(p *Periph) RMSMCR {
	return RMSMCR{mmio.UM32{&p.SMCR.U32, uint32(ETP)}}
}

type DIER uint32

type RDIER struct{ mmio.U32 }

func (r *RDIER) LoadBits(mask DIER) DIER { return DIER(r.U32.LoadBits(uint32(mask))) }
func (r *RDIER) StoreBits(mask, b DIER)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIER) SetBits(mask DIER)       { r.U32.SetBits(uint32(mask)) }
func (r *RDIER) ClearBits(mask DIER)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDIER) Load() DIER              { return DIER(r.U32.Load()) }
func (r *RDIER) Store(b DIER)            { r.U32.Store(uint32(b)) }

type RMDIER struct{ mmio.UM32 }

func (rm RMDIER) Load() DIER   { return DIER(rm.UM32.Load()) }
func (rm RMDIER) Store(b DIER) { rm.UM32.Store(uint32(b)) }

func UIE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(UIE)}}
}

func CC1IE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC1IE)}}
}

func CC2IE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC2IE)}}
}

func CC3IE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC3IE)}}
}

func CC4IE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC4IE)}}
}

func TIE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(TIE)}}
}

func UDE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(UDE)}}
}

func CC1DE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC1DE)}}
}

func CC2DE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC2DE)}}
}

func CC3DE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC3DE)}}
}

func CC4DE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(CC4DE)}}
}

func TDE_(p *Periph) RMDIER {
	return RMDIER{mmio.UM32{&p.DIER.U32, uint32(TDE)}}
}

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

func UIF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(UIF)}}
}

func CC1IF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC1IF)}}
}

func CC2IF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC2IF)}}
}

func CC3IF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC3IF)}}
}

func CC4IF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC4IF)}}
}

func TIF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TIF)}}
}

func CC1OF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC1OF)}}
}

func CC2OF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC2OF)}}
}

func CC3OF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC3OF)}}
}

func CC4OF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CC4OF)}}
}

type EGR uint32

type REGR struct{ mmio.U32 }

func (r *REGR) LoadBits(mask EGR) EGR { return EGR(r.U32.LoadBits(uint32(mask))) }
func (r *REGR) StoreBits(mask, b EGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REGR) SetBits(mask EGR)      { r.U32.SetBits(uint32(mask)) }
func (r *REGR) ClearBits(mask EGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REGR) Load() EGR             { return EGR(r.U32.Load()) }
func (r *REGR) Store(b EGR)           { r.U32.Store(uint32(b)) }

type RMEGR struct{ mmio.UM32 }

func (rm RMEGR) Load() EGR   { return EGR(rm.UM32.Load()) }
func (rm RMEGR) Store(b EGR) { rm.UM32.Store(uint32(b)) }

func UG_(p *Periph) RMEGR {
	return RMEGR{mmio.UM32{&p.EGR.U32, uint32(UG)}}
}

func CC1G_(p *Periph) RMEGR {
	return RMEGR{mmio.UM32{&p.EGR.U32, uint32(CC1G)}}
}

func CC2G_(p *Periph) RMEGR {
	return RMEGR{mmio.UM32{&p.EGR.U32, uint32(CC2G)}}
}

func CC3G_(p *Periph) RMEGR {
	return RMEGR{mmio.UM32{&p.EGR.U32, uint32(CC3G)}}
}

func CC4G_(p *Periph) RMEGR {
	return RMEGR{mmio.UM32{&p.EGR.U32, uint32(CC4G)}}
}

func TG_(p *Periph) RMEGR {
	return RMEGR{mmio.UM32{&p.EGR.U32, uint32(TG)}}
}

type CCMR1 uint32

type RCCMR1 struct{ mmio.U32 }

func (r *RCCMR1) LoadBits(mask CCMR1) CCMR1 { return CCMR1(r.U32.LoadBits(uint32(mask))) }
func (r *RCCMR1) StoreBits(mask, b CCMR1)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCMR1) SetBits(mask CCMR1)        { r.U32.SetBits(uint32(mask)) }
func (r *RCCMR1) ClearBits(mask CCMR1)      { r.U32.ClearBits(uint32(mask)) }
func (r *RCCMR1) Load() CCMR1               { return CCMR1(r.U32.Load()) }
func (r *RCCMR1) Store(b CCMR1)             { r.U32.Store(uint32(b)) }

type RMCCMR1 struct{ mmio.UM32 }

func (rm RMCCMR1) Load() CCMR1   { return CCMR1(rm.UM32.Load()) }
func (rm RMCCMR1) Store(b CCMR1) { rm.UM32.Store(uint32(b)) }

func CC1S_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(CC1S)}}
}

func OC1FE_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC1FE)}}
}

func OC1PE_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC1PE)}}
}

func OC1M_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC1M)}}
}

func OC1CE_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC1CE)}}
}

func CC2S_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(CC2S)}}
}

func OC2FE_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC2FE)}}
}

func OC2PE_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC2PE)}}
}

func OC2M_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC2M)}}
}

func OC2CE_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(OC2CE)}}
}

func ICPCS_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(ICPCS)}}
}

func IC1F_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(IC1F)}}
}

func IC2PCS_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(IC2PCS)}}
}

func IC2F_(p *Periph) RMCCMR1 {
	return RMCCMR1{mmio.UM32{&p.CCMR1.U32, uint32(IC2F)}}
}

type CCMR2 uint32

type RCCMR2 struct{ mmio.U32 }

func (r *RCCMR2) LoadBits(mask CCMR2) CCMR2 { return CCMR2(r.U32.LoadBits(uint32(mask))) }
func (r *RCCMR2) StoreBits(mask, b CCMR2)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCMR2) SetBits(mask CCMR2)        { r.U32.SetBits(uint32(mask)) }
func (r *RCCMR2) ClearBits(mask CCMR2)      { r.U32.ClearBits(uint32(mask)) }
func (r *RCCMR2) Load() CCMR2               { return CCMR2(r.U32.Load()) }
func (r *RCCMR2) Store(b CCMR2)             { r.U32.Store(uint32(b)) }

type RMCCMR2 struct{ mmio.UM32 }

func (rm RMCCMR2) Load() CCMR2   { return CCMR2(rm.UM32.Load()) }
func (rm RMCCMR2) Store(b CCMR2) { rm.UM32.Store(uint32(b)) }

func CC3S_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(CC3S)}}
}

func OC3FE_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC3FE)}}
}

func OC3PE_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC3PE)}}
}

func OC3M_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC3M)}}
}

func OC3CE_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC3CE)}}
}

func CC4S_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(CC4S)}}
}

func OC4FE_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC4FE)}}
}

func OC4PE_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC4PE)}}
}

func OC4M_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(OC4M)}}
}

func O24CE_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(O24CE)}}
}

func IC3PSC_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(IC3PSC)}}
}

func IC3F_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(IC3F)}}
}

func IC4PSC_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(IC4PSC)}}
}

func IC4F_(p *Periph) RMCCMR2 {
	return RMCCMR2{mmio.UM32{&p.CCMR2.U32, uint32(IC4F)}}
}

type CCER uint32

type RCCER struct{ mmio.U32 }

func (r *RCCER) LoadBits(mask CCER) CCER { return CCER(r.U32.LoadBits(uint32(mask))) }
func (r *RCCER) StoreBits(mask, b CCER)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCER) SetBits(mask CCER)       { r.U32.SetBits(uint32(mask)) }
func (r *RCCER) ClearBits(mask CCER)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCCER) Load() CCER              { return CCER(r.U32.Load()) }
func (r *RCCER) Store(b CCER)            { r.U32.Store(uint32(b)) }

type RMCCER struct{ mmio.UM32 }

func (rm RMCCER) Load() CCER   { return CCER(rm.UM32.Load()) }
func (rm RMCCER) Store(b CCER) { rm.UM32.Store(uint32(b)) }

func CC1E_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC1E)}}
}

func CC1P_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC1P)}}
}

func CC1NP_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC1NP)}}
}

func CC2E_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC2E)}}
}

func CC2P_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC2P)}}
}

func CC2NP_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC2NP)}}
}

func CC3E_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC3E)}}
}

func CC3P_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC3P)}}
}

func CC3NP_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC3NP)}}
}

func CC4E_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC4E)}}
}

func CC4P_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC4P)}}
}

func CC4NP_(p *Periph) RMCCER {
	return RMCCER{mmio.UM32{&p.CCER.U32, uint32(CC4NP)}}
}

type CNT uint32

type RCNT struct{ mmio.U32 }

func (r *RCNT) LoadBits(mask CNT) CNT { return CNT(r.U32.LoadBits(uint32(mask))) }
func (r *RCNT) StoreBits(mask, b CNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCNT) SetBits(mask CNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RCNT) ClearBits(mask CNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCNT) Load() CNT             { return CNT(r.U32.Load()) }
func (r *RCNT) Store(b CNT)           { r.U32.Store(uint32(b)) }

type RMCNT struct{ mmio.UM32 }

func (rm RMCNT) Load() CNT   { return CNT(rm.UM32.Load()) }
func (rm RMCNT) Store(b CNT) { rm.UM32.Store(uint32(b)) }

type PSC uint32

type RPSC struct{ mmio.U32 }

func (r *RPSC) LoadBits(mask PSC) PSC { return PSC(r.U32.LoadBits(uint32(mask))) }
func (r *RPSC) StoreBits(mask, b PSC) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSC) SetBits(mask PSC)      { r.U32.SetBits(uint32(mask)) }
func (r *RPSC) ClearBits(mask PSC)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPSC) Load() PSC             { return PSC(r.U32.Load()) }
func (r *RPSC) Store(b PSC)           { r.U32.Store(uint32(b)) }

type RMPSC struct{ mmio.UM32 }

func (rm RMPSC) Load() PSC   { return PSC(rm.UM32.Load()) }
func (rm RMPSC) Store(b PSC) { rm.UM32.Store(uint32(b)) }

type ARR uint32

type RARR struct{ mmio.U32 }

func (r *RARR) LoadBits(mask ARR) ARR { return ARR(r.U32.LoadBits(uint32(mask))) }
func (r *RARR) StoreBits(mask, b ARR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RARR) SetBits(mask ARR)      { r.U32.SetBits(uint32(mask)) }
func (r *RARR) ClearBits(mask ARR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RARR) Load() ARR             { return ARR(r.U32.Load()) }
func (r *RARR) Store(b ARR)           { r.U32.Store(uint32(b)) }

type RMARR struct{ mmio.UM32 }

func (rm RMARR) Load() ARR   { return ARR(rm.UM32.Load()) }
func (rm RMARR) Store(b ARR) { rm.UM32.Store(uint32(b)) }

type CCR uint32

type RCCR struct{ mmio.U32 }

func (r *RCCR) LoadBits(mask CCR) CCR { return CCR(r.U32.LoadBits(uint32(mask))) }
func (r *RCCR) StoreBits(mask, b CCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) SetBits(mask CCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCCR) ClearBits(mask CCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCCR) Load() CCR             { return CCR(r.U32.Load()) }
func (r *RCCR) Store(b CCR)           { r.U32.Store(uint32(b)) }

type RMCCR struct{ mmio.UM32 }

func (rm RMCCR) Load() CCR   { return CCR(rm.UM32.Load()) }
func (rm RMCCR) Store(b CCR) { rm.UM32.Store(uint32(b)) }

type DCR uint32

type RDCR struct{ mmio.U32 }

func (r *RDCR) LoadBits(mask DCR) DCR { return DCR(r.U32.LoadBits(uint32(mask))) }
func (r *RDCR) StoreBits(mask, b DCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDCR) SetBits(mask DCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDCR) ClearBits(mask DCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDCR) Load() DCR             { return DCR(r.U32.Load()) }
func (r *RDCR) Store(b DCR)           { r.U32.Store(uint32(b)) }

type RMDCR struct{ mmio.UM32 }

func (rm RMDCR) Load() DCR   { return DCR(rm.UM32.Load()) }
func (rm RMDCR) Store(b DCR) { rm.UM32.Store(uint32(b)) }

func DBA_(p *Periph) RMDCR {
	return RMDCR{mmio.UM32{&p.DCR.U32, uint32(DBA)}}
}

func DBL_(p *Periph) RMDCR {
	return RMDCR{mmio.UM32{&p.DCR.U32, uint32(DBL)}}
}

type DMAR uint32

type RDMAR struct{ mmio.U32 }

func (r *RDMAR) LoadBits(mask DMAR) DMAR { return DMAR(r.U32.LoadBits(uint32(mask))) }
func (r *RDMAR) StoreBits(mask, b DMAR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDMAR) SetBits(mask DMAR)       { r.U32.SetBits(uint32(mask)) }
func (r *RDMAR) ClearBits(mask DMAR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDMAR) Load() DMAR              { return DMAR(r.U32.Load()) }
func (r *RDMAR) Store(b DMAR)            { r.U32.Store(uint32(b)) }

type RMDMAR struct{ mmio.UM32 }

func (rm RMDMAR) Load() DMAR   { return DMAR(rm.UM32.Load()) }
func (rm RMDMAR) Store(b DMAR) { rm.UM32.Store(uint32(b)) }

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f412

package syscfg

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	MEMRM      RMEMRM
	PMC        RPMC
	EXTICR     [4]REXTICR
	_          [2]uint32
	CMPCR      RCMPCR
	_          [2]uint32
	I2C_BUFOUT RI2C_BUFOUT
}

func SYSCFG() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB2
}

type MEMRM uint32

type RMEMRM struct{ mmio.U32 }

func (r *RMEMRM) LoadBits(mask MEMRM) MEMRM { return MEMRM(r.U32.LoadBits(uint32(mask))) }
func (r *RMEMRM) StoreBits(mask, b MEMRM)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMEMRM) SetBits(mask MEMRM)        { r.U32.SetBits(uint32(mask)) }
func (r *RMEMRM) ClearBits(mask MEMRM)      { r.U32.ClearBits(uint32(mask)) }
func (r *RMEMRM) Load() MEMRM               { return MEMRM(r.U32.Load()) }
func (r *RMEMRM) Store(b MEMRM)             { r.U32.Store(uint32(b)) }

type RMMEMRM struct{ mmio.UM32 }

func (rm RMMEMRM) Load() MEMRM   { return MEMRM(rm.UM32.Load()) }
func (rm RMMEMRM) Store(b MEMRM) { rm.UM32.Store(uint32(b)) }

func MEM_MODE_(p *Periph) RMMEMRM {
	return RMMEMRM{mmio.UM32{&p.MEMRM.U32, uint32(MEM_MODE)}}
}

type PMC uint32

type RPMC struct{ mmio.U32 }

func (r *RPMC) LoadBits(mask PMC) PMC { return PMC(r.U32.LoadBits(uint32(mask))) }
func (r *RPMC) StoreBits(mask, b PMC) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPMC) SetBits(mask PMC)      { r.U32.SetBits(uint32(mask)) }
func (r *RPMC) ClearBits(mask PMC)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPMC) Load() PMC             { return PMC(r.U32.Load()) }
func (r *RPMC) Store(b PMC)           { r.U32.Store(uint32(b)) }

type RMPMC struct{ mmio.UM32 }

func (rm RMPMC) Load() PMC   { return PMC(rm.UM32.Load()) }
func (rm RMPMC) Store(b PMC) { rm.UM32.Store(uint32(b)) }

func ADC1DC2_(p *Periph) RMPMC {
	return RMPMC{mmio.UM32{&p.PMC.U32, uint32(ADC1DC2)}}
}

type EXTICR uint32

type REXTICR struct{ mmio.U32 }

func (r *REXTICR) LoadBits(mask EXTICR) EXTICR { return EXTICR(r.U32.LoadBits(uint32(mask))) }
func (r *REXTICR) StoreBits(mask, b EXTICR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) SetBits(mask EXTICR)         { r.U32.SetBits(uint32(mask)) }
func (r *REXTICR) ClearBits(mask EXTICR)       { r.U32.ClearBits(uint32(mask)) }
func (r *REXTICR) Load() EXTICR                { return EXTICR(r.U32.Load()) }
func (r *REXTICR) Store(b EXTICR)              { r.U32.Store(uint32(b)) }

type RMEXTICR struct{ mmio.UM32 }

func (rm RMEXTICR) Load() EXTICR   { return EXTICR(rm.UM32.Load()) }
func (rm RMEXTICR) Store(b EXTICR) { rm.UM32.Store(uint32(b)) }

type CMPCR uint32

type RCMPCR struct{ mmio.U32 }

func (r *RCMPCR) LoadBits(mask CMPCR) CMPCR { return CMPCR(r.U32.LoadBits(uint32(mask))) }
func (r *RCMPCR) StoreBits(mask, b CMPCR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCMPCR) SetBits(mask CMPCR)        { r.U32.SetBits(uint32(mask)) }
func (r *RCMPCR) ClearBits(mask CMPCR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RCMPCR) Load() CMPCR               { return CMPCR(r.U32.Load()) }
func (r *RCMPCR) Store(b CMPCR)             { r.U32.Store(uint32(b)) }

type RMCMPCR struct{ mmio.UM32 }

func (rm RMCMPCR) Load() CMPCR   { return CMPCR(rm.UM32.Load()) }
func (rm RMCMPCR) Store(b CMPCR) { rm.UM32.Store(uint32(b)) }

func CMP_PD_(p *Periph) RMCMPCR {
	return RMCMPCR{mmio.UM32{&p.CMPCR.U32, uint32(CMP_PD)}}
}

func READY_(p *Periph) RMCMPCR {
	return RMCMPCR{mmio.UM32{&p.CMPCR.U32, uint32(READY)}}
}

type I2C_BUFOUT uint32

type RI2C_BUFOUT struct{ mmio.U32 }

func (r *RI2C_BUFOUT) LoadBits(mask I2C_BUFOUT) I2C_BUFOUT {
	return I2C_BUFOUT(r.U32.LoadBits(uint32(mask)))
}
func (r *RI2C_BUFOUT) StoreBits(mask, b I2C_BUFOUT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RI2C_BUFOUT) SetBits(mask I2C_BUFOUT)      { r.U32.SetBits(uint32(mask)) }
func (r *RI2C_BUFOUT) ClearBits(mask I2C_BUFOUT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RI2C_BUFOUT) Load() I2C_BUFOUT             { return I2C_BUFOUT(r.U32.Load()) }
func (r *RI2C_BUFOUT) Store(b I2C_BUFOUT)           { r.U32.Store(uint32(b)) }

type RMI2C_BUFOUT struct{ mmio.UM32 }

func (rm RMI2C_BUFOUT) Load() I2C_BUFOUT   { return I2C_BUFOUT(rm.UM32.Load()) }
func (rm RMI2C_BUFOUT) Store(b I2C_BUFOUT) { rm.UM32.Store(uint32(b)) }

func I2C4SCL_(p *Periph) RMI2C_BUFOUT {
	return RMI2C_BUFOUT{mmio.UM32{&p.I2C_BUFOUT.U32, uint32(I2C4SCL)}}
}

func I2C4SDA_(p *Periph) RMI2C_BUFOUT {
	return RMI2C_BUFOUT{mmio.UM32{&p.I2C_BUFOUT.U32, uint32(I2C4SDA)}}
}

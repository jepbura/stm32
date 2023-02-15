// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f215

package syscfg

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	MEMRM  mmio.R32[MEMRM]
	PMC    mmio.R32[PMC]
	EXTICR [4]mmio.R32[uint32]
	_      [2]uint32
	CMPCR  mmio.R32[CMPCR]
}

func SYSCFG() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB2
}

type MEMRM uint32

func MEM_MODE_(p *Periph) mmio.RM32[MEMRM] { return mmio.RM32[MEMRM]{&p.MEMRM, MEM_MODE} }

type PMC uint32

func MII_RMII_SEL_(p *Periph) mmio.RM32[PMC] { return mmio.RM32[PMC]{&p.PMC, MII_RMII_SEL} }

type CMPCR uint32

func CMP_PD_(p *Periph) mmio.RM32[CMPCR] { return mmio.RM32[CMPCR]{&p.CMPCR, CMP_PD} }
func READY_(p *Periph) mmio.RM32[CMPCR]  { return mmio.RM32[CMPCR]{&p.CMPCR, READY} }

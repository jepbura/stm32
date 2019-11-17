// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32l4x6

package exti

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	IMR1   RIMR1
	EMR1   REMR1
	RTSR1  RRTSR1
	FTSR1  RFTSR1
	SWIER1 RSWIER1
	PR1    RPR1
	_      [2]uint32
	IMR2   RIMR2
	EMR2   REMR2
	RTSR2  RRTSR2
	FTSR2  RFTSR2
	SWIER2 RSWIER2
	PR2    RPR2
}

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func EXTI() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.EXTI_BASE))) }

type IMR1 uint32

type RIMR1 struct{ mmio.U32 }

func (r *RIMR1) Bits(mask IMR1) IMR1    { return IMR1(r.U32.Bits(uint32(mask))) }
func (r *RIMR1) StoreBits(mask, b IMR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIMR1) SetBits(mask IMR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RIMR1) ClearBits(mask IMR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIMR1) Load() IMR1             { return IMR1(r.U32.Load()) }
func (r *RIMR1) Store(b IMR1)           { r.U32.Store(uint32(b)) }

type RMIMR1 struct{ mmio.UM32 }

func (rm RMIMR1) Load() IMR1   { return IMR1(rm.UM32.Load()) }
func (rm RMIMR1) Store(b IMR1) { rm.UM32.Store(uint32(b)) }

type EMR1 uint32

type REMR1 struct{ mmio.U32 }

func (r *REMR1) Bits(mask EMR1) EMR1    { return EMR1(r.U32.Bits(uint32(mask))) }
func (r *REMR1) StoreBits(mask, b EMR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REMR1) SetBits(mask EMR1)      { r.U32.SetBits(uint32(mask)) }
func (r *REMR1) ClearBits(mask EMR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *REMR1) Load() EMR1             { return EMR1(r.U32.Load()) }
func (r *REMR1) Store(b EMR1)           { r.U32.Store(uint32(b)) }

type RMEMR1 struct{ mmio.UM32 }

func (rm RMEMR1) Load() EMR1   { return EMR1(rm.UM32.Load()) }
func (rm RMEMR1) Store(b EMR1) { rm.UM32.Store(uint32(b)) }

type RTSR1 uint32

type RRTSR1 struct{ mmio.U32 }

func (r *RRTSR1) Bits(mask RTSR1) RTSR1   { return RTSR1(r.U32.Bits(uint32(mask))) }
func (r *RRTSR1) StoreBits(mask, b RTSR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRTSR1) SetBits(mask RTSR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RRTSR1) ClearBits(mask RTSR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRTSR1) Load() RTSR1             { return RTSR1(r.U32.Load()) }
func (r *RRTSR1) Store(b RTSR1)           { r.U32.Store(uint32(b)) }

type RMRTSR1 struct{ mmio.UM32 }

func (rm RMRTSR1) Load() RTSR1   { return RTSR1(rm.UM32.Load()) }
func (rm RMRTSR1) Store(b RTSR1) { rm.UM32.Store(uint32(b)) }

type FTSR1 uint32

type RFTSR1 struct{ mmio.U32 }

func (r *RFTSR1) Bits(mask FTSR1) FTSR1   { return FTSR1(r.U32.Bits(uint32(mask))) }
func (r *RFTSR1) StoreBits(mask, b FTSR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFTSR1) SetBits(mask FTSR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RFTSR1) ClearBits(mask FTSR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFTSR1) Load() FTSR1             { return FTSR1(r.U32.Load()) }
func (r *RFTSR1) Store(b FTSR1)           { r.U32.Store(uint32(b)) }

type RMFTSR1 struct{ mmio.UM32 }

func (rm RMFTSR1) Load() FTSR1   { return FTSR1(rm.UM32.Load()) }
func (rm RMFTSR1) Store(b FTSR1) { rm.UM32.Store(uint32(b)) }

type SWIER1 uint32

type RSWIER1 struct{ mmio.U32 }

func (r *RSWIER1) Bits(mask SWIER1) SWIER1  { return SWIER1(r.U32.Bits(uint32(mask))) }
func (r *RSWIER1) StoreBits(mask, b SWIER1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSWIER1) SetBits(mask SWIER1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSWIER1) ClearBits(mask SWIER1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSWIER1) Load() SWIER1             { return SWIER1(r.U32.Load()) }
func (r *RSWIER1) Store(b SWIER1)           { r.U32.Store(uint32(b)) }

type RMSWIER1 struct{ mmio.UM32 }

func (rm RMSWIER1) Load() SWIER1   { return SWIER1(rm.UM32.Load()) }
func (rm RMSWIER1) Store(b SWIER1) { rm.UM32.Store(uint32(b)) }

type PR1 uint32

type RPR1 struct{ mmio.U32 }

func (r *RPR1) Bits(mask PR1) PR1     { return PR1(r.U32.Bits(uint32(mask))) }
func (r *RPR1) StoreBits(mask, b PR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPR1) SetBits(mask PR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RPR1) ClearBits(mask PR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPR1) Load() PR1             { return PR1(r.U32.Load()) }
func (r *RPR1) Store(b PR1)           { r.U32.Store(uint32(b)) }

type RMPR1 struct{ mmio.UM32 }

func (rm RMPR1) Load() PR1   { return PR1(rm.UM32.Load()) }
func (rm RMPR1) Store(b PR1) { rm.UM32.Store(uint32(b)) }

type IMR2 uint32

type RIMR2 struct{ mmio.U32 }

func (r *RIMR2) Bits(mask IMR2) IMR2    { return IMR2(r.U32.Bits(uint32(mask))) }
func (r *RIMR2) StoreBits(mask, b IMR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIMR2) SetBits(mask IMR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RIMR2) ClearBits(mask IMR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIMR2) Load() IMR2             { return IMR2(r.U32.Load()) }
func (r *RIMR2) Store(b IMR2)           { r.U32.Store(uint32(b)) }

type RMIMR2 struct{ mmio.UM32 }

func (rm RMIMR2) Load() IMR2   { return IMR2(rm.UM32.Load()) }
func (rm RMIMR2) Store(b IMR2) { rm.UM32.Store(uint32(b)) }

type EMR2 uint32

type REMR2 struct{ mmio.U32 }

func (r *REMR2) Bits(mask EMR2) EMR2    { return EMR2(r.U32.Bits(uint32(mask))) }
func (r *REMR2) StoreBits(mask, b EMR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REMR2) SetBits(mask EMR2)      { r.U32.SetBits(uint32(mask)) }
func (r *REMR2) ClearBits(mask EMR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *REMR2) Load() EMR2             { return EMR2(r.U32.Load()) }
func (r *REMR2) Store(b EMR2)           { r.U32.Store(uint32(b)) }

type RMEMR2 struct{ mmio.UM32 }

func (rm RMEMR2) Load() EMR2   { return EMR2(rm.UM32.Load()) }
func (rm RMEMR2) Store(b EMR2) { rm.UM32.Store(uint32(b)) }

type RTSR2 uint32

type RRTSR2 struct{ mmio.U32 }

func (r *RRTSR2) Bits(mask RTSR2) RTSR2   { return RTSR2(r.U32.Bits(uint32(mask))) }
func (r *RRTSR2) StoreBits(mask, b RTSR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRTSR2) SetBits(mask RTSR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RRTSR2) ClearBits(mask RTSR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRTSR2) Load() RTSR2             { return RTSR2(r.U32.Load()) }
func (r *RRTSR2) Store(b RTSR2)           { r.U32.Store(uint32(b)) }

type RMRTSR2 struct{ mmio.UM32 }

func (rm RMRTSR2) Load() RTSR2   { return RTSR2(rm.UM32.Load()) }
func (rm RMRTSR2) Store(b RTSR2) { rm.UM32.Store(uint32(b)) }

type FTSR2 uint32

type RFTSR2 struct{ mmio.U32 }

func (r *RFTSR2) Bits(mask FTSR2) FTSR2   { return FTSR2(r.U32.Bits(uint32(mask))) }
func (r *RFTSR2) StoreBits(mask, b FTSR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFTSR2) SetBits(mask FTSR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RFTSR2) ClearBits(mask FTSR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFTSR2) Load() FTSR2             { return FTSR2(r.U32.Load()) }
func (r *RFTSR2) Store(b FTSR2)           { r.U32.Store(uint32(b)) }

type RMFTSR2 struct{ mmio.UM32 }

func (rm RMFTSR2) Load() FTSR2   { return FTSR2(rm.UM32.Load()) }
func (rm RMFTSR2) Store(b FTSR2) { rm.UM32.Store(uint32(b)) }

type SWIER2 uint32

type RSWIER2 struct{ mmio.U32 }

func (r *RSWIER2) Bits(mask SWIER2) SWIER2  { return SWIER2(r.U32.Bits(uint32(mask))) }
func (r *RSWIER2) StoreBits(mask, b SWIER2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSWIER2) SetBits(mask SWIER2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSWIER2) ClearBits(mask SWIER2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSWIER2) Load() SWIER2             { return SWIER2(r.U32.Load()) }
func (r *RSWIER2) Store(b SWIER2)           { r.U32.Store(uint32(b)) }

type RMSWIER2 struct{ mmio.UM32 }

func (rm RMSWIER2) Load() SWIER2   { return SWIER2(rm.UM32.Load()) }
func (rm RMSWIER2) Store(b SWIER2) { rm.UM32.Store(uint32(b)) }

type PR2 uint32

type RPR2 struct{ mmio.U32 }

func (r *RPR2) Bits(mask PR2) PR2     { return PR2(r.U32.Bits(uint32(mask))) }
func (r *RPR2) StoreBits(mask, b PR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPR2) SetBits(mask PR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RPR2) ClearBits(mask PR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPR2) Load() PR2             { return PR2(r.U32.Load()) }
func (r *RPR2) Store(b PR2)           { r.U32.Store(uint32(b)) }

type RMPR2 struct{ mmio.UM32 }

func (rm RMPR2) Load() PR2   { return PR2(rm.UM32.Load()) }
func (rm RMPR2) Store(b PR2) { rm.UM32.Store(uint32(b)) }
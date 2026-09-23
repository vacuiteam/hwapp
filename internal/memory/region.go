// Copyright 2026 Vacui Development Team - Apache License 2.0
// https://github.com/vacuiteam/hwapp

package memory

// RegionFlag is a permission flag that manages the guest program and hardware's
// access to specific memory region.
type RegionFlag uint32

// Region consists of its physical address, size and permission flags
// for the hardware and a guest program.
type Region struct {
	PhysAddress       uint64
	Size              uint64
	HardwareFlags     RegionFlag
	GuestProgramFlags RegionFlag
}

const (
	RegionFlagRead    RegionFlag = 1 << 0
	RegionFlagWrite   RegionFlag = 1 << 1
	RegionFlagExecute RegionFlag = 1 << 2
)

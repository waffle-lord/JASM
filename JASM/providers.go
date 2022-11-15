package main

type IHardwareProvider interface {
	GetCpuInfo() CpuInfo
	GetMemoryInfo() MemoryInfo
	GetGpuInfo() GpuInfo
}

type WindowsHardwareProvider struct {
}

func (whp *WindowsHardwareProvider) GetCpuInfo() *CpuInfo {
	return nil
}

func (whp *WindowsHardwareProvider) GetMemoryInfo() *MemoryInfo {
	return nil
}

func (whp *WindowsHardwareProvider) GetGpuInfo() *GpuInfo {
	return nil
}

type LinuxHardwareProvider struct {
}

func (whp *LinuxHardwareProvider) GetCpuInfo() *CpuInfo {
	return nil
}

func (whp *LinuxHardwareProvider) GetMemoryInfo() *MemoryInfo {
	return nil
}

func (whp *LinuxHardwareProvider) GetGpuInfo() *GpuInfo {
	return nil
}

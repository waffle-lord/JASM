package main

type CpuInfo struct {
	Model     string
	CoreCount int
	CoreLoads []float64
	TotalLoad float64
}

type MemoryInfo struct {
	Total       float64
	Available   float64
	Used        float64
	PercentUsed int
}

type GpuInfo struct {
	// TODO - this, blah
}

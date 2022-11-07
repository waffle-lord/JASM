package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// App struct
type App struct {
	ctx context.Context
}

type CpuInfo struct {
	Model     string
	CoreCount int
	CoreLoads []float64
	TotalLoad float64
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// SetContext is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetCpuData() *CpuInfo {
	fmt.Printf("%v :: Getting CPU data\n", time.Now().Unix())

	info, _ := cpu.Info()

	cLoads, _ := cpu.Percent(time.Second*1, true)

	tLoad, _ := cpu.Percent(time.Second*1, false)

	return &CpuInfo{
		Model:     info[0].ModelName,
		CoreCount: int(info[0].Cores),
		CoreLoads: cLoads,
		TotalLoad: tLoad[0],
	}
}

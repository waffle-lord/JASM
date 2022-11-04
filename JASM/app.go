package main

import (
	"context"
	// "fmt"

	// "time"
	// "github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/shirou/gopsutil/v3/cpu"
)

// App struct
type App struct {
	ctx context.Context
}

type CpuInfo struct {
	Model     string
	CoreCount int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetCpuData() *CpuInfo {
	info, _ := cpu.Info()

	return &CpuInfo{
		Model:     info[0].ModelName,
		CoreCount: int(info[0].Cores),
	}
}

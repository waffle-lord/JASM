package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	ticker      time.Ticker
	tickChannel chan bool
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

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) refreshData() {
	for {
		select {
		case tick := <-a.ticker.C:
			// update frontend data
			fmt.Printf("Tick  :: %s\n", tick)
			fmt.Println("Event :: dataRefresh")
			runtime.EventsEmit(a.ctx, "dataRefresh")

		case <-a.tickChannel:
			fmt.Println("tickChannel closing")
			close(a.tickChannel)
		}
	}
}

func (a *App) StartTimer() {
	a.tickChannel = make(chan bool)
	a.ticker = *time.NewTicker(time.Second * 1)
	go a.refreshData()
	fmt.Println("Refresh Timer Started")
}

func (a *App) StopTimer() {
	a.ticker.Stop()
	a.tickChannel <- true
	fmt.Println("Refresh Timer Stopped")
}

func (a *App) GetCpuData() *CpuInfo {
	fmt.Println("Getting CPU data")

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

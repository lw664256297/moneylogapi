package sd

import (
	"fmt"
	"net/http"

	// . "moneylogapi/handler"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// SysInfo 系统状态返回值
type SysInfo struct {
	ApiStatus string `json:"apiStatus"`
	ApiCpu    string `json:"apiCpu"`
	ApiDisk   string `json:"apiDisk"`
	ApiRAM    string `json:"apiRAM"`
}

// HealthCheck 健康状态
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

// DiskCheck 硬盘状态
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)
	status := http.StatusOK
	text := "良好"
	if usedPercent >= 95 {
		text = "危机"
	} else if usedPercent >= 90 {
		text = "警告"
	}
	message := fmt.Sprintf("%s - 使用空间: %dMB (%dGB) / %dMB (%dGB) | 使用率: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}

// CPUCheck cpu运行状况
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15
	status := http.StatusOK
	text := "良好"
	if l5 >= float64(cores-1) {
		text = "危机"
	} else if l5 >= float64(cores-2) {
		text = "警告"
	}
	message := fmt.Sprintf("%s - 平均负载: %.2f, %.2f, %.2f | 核心数: %d", text, l1, l5, l15, cores)
	c.String(status, "\n"+message)
}

// RAMCheck 内存使用情况
func RAMCheck(c *gin.Context) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)
	status := http.StatusOK
	text := "良好"
	if usedPercent >= 95 {
		text = "危机"
	} else if usedPercent >= 90 {
		text = "警告"
	}

	message := fmt.Sprintf("%s - 可用空间: %dMB (%dGB) / %dMB (%dGB) | 使用率: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}

// SYSInfoCheck 所有状态
// func SYSInfoCheck(c *gin.Context) {
// 	rsp := SysInfo{
// 		ApiStatus: "OK",
// 		ApiCpu:    CPUCheck(c),
// 		ApiDisk:   DiskCheck(c),
// 		ApiRAM:    RAMCheck(c),
// 	}

// 	SendResponse(c, nil, rsp)
// }

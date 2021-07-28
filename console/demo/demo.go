package demo

import (
	"BearApp/internal/bootstrap"
)

// Run 背景
func Run() error {
	bootstrap.WriteLog("INFO", "Demo Job")
	return nil
}

package main

import (
	"BearApp/internal/server"
	"BearApp/model"

	"os"
)

func init() {
	os.Setenv("TZ", "Asia/Taipei")
}

func main() {
	// go schedule.Run()
	server.Run()

	model.SetupTable()

}

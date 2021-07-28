package model

import (
	"BearApp/internal/bootstrap"
	"log"
	"os"
	"strings"
)

// SetupTable 設置資料表
func SetupTable() {
	//
	modelList := []IModel{
		new(Account), new(Operation),
	}

	env := bootstrap.GetAppConf().App.Env
	if env == "local" || env == "docker" {
		for _, m := range modelList {
			log.Print(m)
			db, err := NewModelDB(m, true)
			if err != nil {
				bootstrap.WriteLog("ERROR", "DB連線失敗, "+err.Error())
				os.Exit(1)
				return
			}

			err = db.AutoMigrate(m).Error
			if err != nil {
				bootstrap.WriteLog("ERROR", "建立資料表失敗, "+err.Error())
				os.Exit(1)
				return
			}
		}

		return
	}

	missingTable := []string{}
	for _, m := range modelList {
		db, err := NewModelDB(m, false)
		if err != nil {
			bootstrap.WriteLog("ERROR", "DB連線失敗, "+err.Error())
			os.Exit(1)
			return
		}

		if !db.HasTable(m.TableName()) {
			missingTable = append(missingTable, m.TableName())
		}
	}

	if len(missingTable) > 0 {
		bootstrap.WriteLog("ERROR", "缺少資料表: "+strings.Join(missingTable, ", "))
		os.Exit(1)
		return
	}
}

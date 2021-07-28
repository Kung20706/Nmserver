package bootstrap

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// GetAppEnv 取環境變數
func GetAppEnv() string {
	env := strings.TrimSpace(os.Getenv("PROJECT_ENV"))
	if env == "" {
		env = "default"
	}
	return env
}

// GetAppSite 取客端變數
func GetAppSite() string {
	site := strings.TrimSpace(os.Getenv("PROJECT_SITE"))
	if site == "" {
		site = "default"
	}
	return site
}

// GetAppRoot 取專案的根目錄
func GetAppRoot() string {
	var root string
	if os.Getenv("PROJECT_ROOT") == "" {
		execRoot, err := os.Getwd()
		if err != nil {
			WriteLog("WARNING", fmt.Sprintf("🎃  GetAppRoot 取根目錄失敗 (%v) 🎃", err))
		}
		root = execRoot
	} else {
		root = os.Getenv("PROJECT_ROOT")
	}

	return root
}

// GetAppConf 取專案的設定檔
func GetAppConf() *Config {
	if Conf != nil {
		return Conf
	}
	return LoadConfig()
}

// WriteLog 寫Log記錄檔案
func WriteLog(tag string, msg string) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			log.Println(time.Now().Format("[2006-01-02 15:04:05]")+"【ERROR】 WriteLog: 寫Log檔案發生意外！", catchErr)
		}
	}()
	//設定時間
	now := time.Now()

	// 組合字串
	logStr := now.Format("[2006-01-02 15:04:05]") + "【" + tag + "】" + msg + "\n"
	log.Print(logStr)

	// 設定檔案位置
	fileName := "Beartest.log"
	folderPath := GetAppRoot() + now.Format("/storage/logs/2006-01-02/15/")

	//檢查今日log檔案是否存在
	if _, err := os.Stat(folderPath + fileName); os.IsNotExist(err) {
		//建立資料夾
		os.MkdirAll(folderPath, 0777)
		//建立檔案
		_, err := os.Create(folderPath + fileName)
		if err != nil {
			log.Printf("❌ WriteLog: 建立檔案錯誤 [%v] ❌ \n----> %s\n", err, msg)
			return
		}
	}

	//開啟檔案準備寫入
	logFile, err := os.OpenFile(folderPath+fileName, os.O_RDWR|os.O_APPEND, 0777)
	defer logFile.Close()
	if err != nil {
		log.Printf("❌ WriteLog: 開啟檔案錯誤 [%v] ❌ \n----> %s\n", err, msg)
		return
	}

	_, err = logFile.WriteString(logStr)

	if err != nil {
		log.Printf("❌ WriteLog: 寫入檔案錯誤 [%v] ❌ \n----> %s\n", err, msg)
	}
}

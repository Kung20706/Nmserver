package helper

import (
	"encoding/json"
	"log"
	"runtime"
	"strconv"
	"strings"
)

// MyCaller 取call我的人
func MyCaller() (name string) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			log.Println("🎃  helper.MyCaller 發生錯誤!", catchErr, " 🎃")
			return
		}
	}()

	fpcs := make([]uintptr, 1)
	n := runtime.Callers(4, fpcs)
	if n == 0 {
		return
	}

	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return
	}

	funcFile, funcLine := fun.FileLine(fun.Entry())
	specs := strings.Split(fun.Name(), ".")
	if len(specs) == 0 {
		name = "檔案: " + funcFile + " #行數: " + strconv.Itoa(funcLine)
		return
	}

	name = "💡 檔案: " + funcFile + " 📍 Func名稱: " + specs[len(specs)-1] + ", 📍 行數: " + strconv.Itoa(funcLine)
	return
}

// InArrayInt64 是否在陣列中
func InArrayInt64(array []int64, val int64) (ok bool) {
	for _, data := range array {
		if data == val {
			ok = true
			return
		}
	}
	return
}

//ToJSON 將資料解成byte
func ToJSON(v interface{}) (j []byte, err error) {
	j, err = json.Marshal(v)
	return
}

// FromJSON 從JSON轉回資料
func FromJSON(j []byte, v interface{}) (err error) {
	err = json.Unmarshal(j, &v)
	return
}

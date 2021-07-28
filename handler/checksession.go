package Accountapi
import (	
	"log"
constant "BearApp/constant"
)
func CheckInternal(token string)  int{
	
	if token == constant.Internaltoken{
		log.Print(token)
		return 1
	}//預計這裡會去check 是否為會員 再拉出一份app的會員
	return 0
}
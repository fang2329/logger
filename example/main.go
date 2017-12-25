package main


import (
	"../common_log"
	)


func main(){
	logger :=common_log.NewRealStLogger(0)
	logger.DEBUG("HELLO")
	logger.ERROR("ERROR")
	logger.INFO("INFO")
}

package main


import (
	"../common_log"
	"sync"
	)


func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	go log_d(&wg)
	go log_e(&wg)
	go log_i(&wg)
	wg.Wait()
}

func log_d( wg *sync.WaitGroup){
    defer wg.Done()
    for i := 0; i < 1000;i++{
 	common_log.LOG_DEBUG("HELLO %d",i)
    }
}

func log_e( wg *sync.WaitGroup){
    defer wg.Done()
    for i := 0; i < 1000;i++{
    	common_log.LOG_ERROR("ERROR %d",i)
    }
}

func log_i( wg *sync.WaitGroup){
    defer wg.Done()
    for i := 0;i < 1000;i++{
    	 common_log.LOG_INFO("INFO %d",i)
    }
}

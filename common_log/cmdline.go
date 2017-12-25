package common_log

import "flag"

var(
	logD = flag.String("logdir","../log/","log directory name")
	maxFileNum = flag.Int("num",10,"everyday log file num")
	maxFileCap = flag.Int("cap",1024*1024*50,"max log data ")
	delDay = flag.Uint("days",3,"log dir save days")
)

func init()  {
	flag.Parse()
}
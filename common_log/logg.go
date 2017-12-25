package common_log

var logger = NewRealStLogger(0)

func  LOG_DEBUG(format string,args ...interface{})  {
	logger.DEBUG(format,args...)
}

func LOG_INFO(format string,args ...interface{})  {
	logger.INFO(format,args...)
}

func LOG_WARNING(format string,args ...interface{})  {
	logger.WARNING(format,args...)
}

func LOG_ERROR(format string,args ...interface{})  {
	logger.ERROR(format,args...)
}

func LOG_CRITIC(format string,args ...interface{})  {
	logger.CRITIC(format,args...)
}

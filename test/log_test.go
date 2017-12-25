package log_test


import(
	 "../common_log"
	"testing"
)

func Test_debug(t *testing.T){
    logger :=common_log.NewRealStLogger(0)
    logger.DEBUG("test logger")
}

func Test_info(t *testing.T){
    logger :=common_log.NewRealStLogger(0)
    logger.INFO("test logger")
}

func Test_warning(t *testing.T){
    logger :=common_log.NewRealStLogger(0)
    logger.WARNING("test logger")
}

func Test_error(t *testing.T){
    logger :=common_log.NewRealStLogger(0)
    logger.ERROR("test logger")
}

func Test_critic(t *testing.T){
    logger :=common_log.NewRealStLogger(0)
    logger.CRITIC("test logger")
}

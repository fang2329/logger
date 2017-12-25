package log_test


import(
	 "../common_log"
	"testing"
)

func Test_debug(t *testing.T){
    common_log.LOG_DEBUG("test logger")
}

func Test_info(t *testing.T){
     common_log.LOG_INFO("test logger")
}

func Test_warning(t *testing.T){
     common_log.LOG_WARNING("test logger")
}

func Test_error(t *testing.T){
     common_log.LOG_ERROR("test logger")
}

func Test_critic(t *testing.T){
     common_log.LOG_CRITIC("test logger")
}

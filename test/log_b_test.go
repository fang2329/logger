package log_test

import(
    "testing"
    "../common_log"
)


func Benchmark_debug(b *testing.B){
    logger :=common_log.NewRealStLogger(0)
    for i := 0;i < b.N;i++{
	logger.DEBUG("benchmark test")
    }
}


func Benchmark_info(b *testing.B){
    logger :=common_log.NewRealStLogger(0)
    for i := 0;i < b.N;i++{
        logger.INFO("benchmark test")
    }
}

func Benchmark_warning(b *testing.B){
    logger :=common_log.NewRealStLogger(0)
    for i := 0;i < b.N;i++{
        logger.WARNING("benchmark test")
    }
}

func Benchmark_error(b *testing.B){
    logger :=common_log.NewRealStLogger(0)
    for i := 0;i < b.N;i++{
        logger.ERROR("benchmark test")
    }
}

func Benchmark_critic(b *testing.B){
    logger :=common_log.NewRealStLogger(0)
    for i := 0;i < b.N;i++{
        logger.CRITIC("benchmark test")
    }
}

package log_test

import(
    "testing"
    "../common_log"
)


func Benchmark_debug(b *testing.B){
    for i := 0;i < b.N;i++{
	common_log.LOG_DEBUG("benchmark test")
    }
}


func Benchmark_info(b *testing.B){
    for i := 0;i < b.N;i++{
        common_log.LOG_INFO("benchmark test")
    }
}

func Benchmark_warning(b *testing.B){
    for i := 0;i < b.N;i++{
        common_log.LOG_WARNING("benchmark test")
    }
}

func Benchmark_error(b *testing.B){
    for i := 0;i < b.N;i++{
        common_log.LOG_ERROR("benchmark test")
    }
}

func Benchmark_critic(b *testing.B){
    for i := 0;i < b.N;i++{
        common_log.LOG_CRITIC("benchmark test")
    }
}

package test

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

var limiter *rate.Limiter

func init() {
	//第一个参数是r Limit。代表每秒可以向Token桶中产生多少token。Limit实际上是float64的别名。
	//第二个参数是b int。b代表Token桶的容量大小。
	limiter = rate.NewLimiter(10, 20) //其令牌桶大小为80, 以每秒100个Token的速率向桶中放置Token。

	//limiter.SetLimit(50) //改变放入Token的速率

	//limit := rate.Every(100 * time.Millisecond);
	//limiter = rate.NewLimiter(limit, 1);
}

func TestRate(t *testing.T) {
	//for i := 0; i < 100; i++ {
	//	go func(d int) {
	//		allow := limiter.Allow()
	//		fmt.Println(d, allow)
	//	}(i)
	//}
	//time.Sleep(10 * time.Second)

	for i := 0; i < 100; i++ {
		go func(d int) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			wait := limiter.Wait(ctx)
			fmt.Println(d, wait)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

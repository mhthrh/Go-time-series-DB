package Server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mhthrh/TimeSeriesDb/Influx"
	"net/http"
	"strconv"
	"time"
)

var (
	d chan map[string]interface{}
	e chan error
)

func init() {
	d = make(chan map[string]interface{})
	e = make(chan error)
	c := Influx.New("http://localhost:8086", "XJeyY9FihSITmTRyrslmpCDrxEG5KRrCfqUfUqxek_jEpHXqkHkreWEvbhsC-Am-Bh0q0nfoiKLQvfFpkXSftQ==")
	go c.Write("my0rg", "BlackBucket", Influx.InfluxChan{
		Ctx:     context.Background(),
		Data:    &d,
		ErrChan: &e,
	})
}
func RunServer() http.Handler {
	start := time.Now()
	router := gin.New()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
		d <- map[string]interface{}{"name": "ping", "age": time.Now().Sub(start), "response": "pong"}
	})
	router.GET("/prime", func(ctx *gin.Context) {
		start := time.Now()
		defer func() {
			d <- map[string]interface{}{"name": "prime", "age": time.Now().Sub(start), "response": ""}
		}()
		fmt.Println(ctx.Query("number"), " ", "received")
		number, err := strconv.ParseInt(ctx.Query("number"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid input")
			return
		}
		if !isPrime(int(number)) {
			ctx.JSON(http.StatusOK, fmt.Sprintf("%d is not a prime number", number))
			return
		}
		ctx.JSON(http.StatusOK, fmt.Sprintf("%d is a prime number", number))
	})
	return router
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

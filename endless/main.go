package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func preSigUsr1() {
	log.Println("重启之前。。。。。")
}

func postSigUsr1() {
	log.Println("重启之后。。。。。")
}

func main() {

	fmt.Printf("PID: %d\n", os.Getpid())

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// 这个5秒的延时。是为了演示操作方便，实际上线一定注释掉
		sleep := c.DefaultQuery("sleep", "0s")
		duration, _ := time.ParseDuration(sleep)
		time.Sleep(duration)
		c.String(http.StatusOK, "hello xiaosheng !")
	})

	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	//if err := endless.ListenAndServe(":8083", router); err != nil {
	//	log.Fatalf("listen: %s\n", err)
	//}

	srv := endless.NewServer("localhost:4244", router)

	srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGHUP] = append(
		srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGHUP],
		preSigUsr1)

	srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGHUP] = append(
		srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGHUP],
		postSigUsr1)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

	log.Println("Server exiting...")

}

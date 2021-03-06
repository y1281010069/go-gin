package main

import (
    "fmt"
    "log"
    "syscall"

    "github.com/fvbock/endless"

    "github.com/y1281010069/go-gin/routers"
    "github.com/y1281010069/go-gin/pkg/setting"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/y1281010069/go-gin

// @license.name MIT
// @license.url https://github.com/y1281010069/go-gin/blob/master/LICENSE

// @BasePath /
func main() {
    endless.DefaultReadTimeOut = setting.ReadTimeout
    endless.DefaultWriteTimeOut = setting.WriteTimeout
    endless.DefaultMaxHeaderBytes = 1 << 20
    endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

    server := endless.NewServer(endPoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getpid())
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err: %v", err)
    }
}
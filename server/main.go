package main

import (
	"flag"
	"fmt"
	"goProject/student/server/config"
	"goProject/student/server/grpc"
	"goProject/student/server/models/psql"
	"goProject/student/server/router"
	"os"
	"os/signal"
	"syscall"
)

func initDB() {
	var tomlFile string
	flag.StringVar(&tomlFile, "config", "docs/test.toml", "服务配置文件")
	// 解析配置文件
	tomlConfig, err := config.UnmarshalConfig(tomlFile)
	if err != nil {
		fmt.Println("UnmarshalConfig: err:%v\n", err)
		return
	}
	psql.NewStudentDBConn("test_db", tomlConfig)
}

func main() {
	initDB()
	go grpc.Start("127.0.0.1:8080")
	//defer grpc.Stop()
	engine, err := router.InitEngine("release")
	if err != nil {
		panic(err)
	}

	go engine.Run(":8080")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	<-signalChan
	// 收到退出信号 准备退出
	grpc.Stop()
	//r := gin.Default()
	//r.GET("/students/server/:id", handlers.GetById)
	//r.Run(":8080")
}

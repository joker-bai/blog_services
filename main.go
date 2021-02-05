package main

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/internal/model"
	"code.coolops.cn/blog_services/internal/routers"
	"code.coolops.cn/blog_services/pkg/logger"
	setting2 "code.coolops.cn/blog_services/pkg/setting"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	// 获取命令行参数
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	// 初始化配置
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// 初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	// 初始化日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

// @title 博客系统
// @version 1.0
// @description 简单的博客系统
// @termsOfService https://github.com/qiaokebaba/blog_service

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	//log.Fatal("开始启动")
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	go func() {
		err := s.ListenAndServe()
		if err !=nil && err !=http.ErrServerClosed {
			log.Fatalf("s.ListenAndServer err: %v",err)
		}
	}()
	// 等待中断信号
	quit := make(chan  os.Signal)
	// 接收syscall.SIGINT和syscall.SIGTERM信号
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Fatalf("Shuting down server ......")

	// 最大时间控制，用于通知该服务端它有5秒的时间来处理原请求
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx);err !=nil{
		log.Fatalf("Server force to shutdown: %v",err)
	}
	log.Fatalf("Server existing ......")
}

// 获取命令行参数
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
	//log.Fatal(config,port,runMode)
	return nil
}

// 初始化配置文件
func setupSetting() error {
	setting, err := setting2.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Middleware", &global.MiddlewareSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	// 初始化Email
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

// 初始化数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

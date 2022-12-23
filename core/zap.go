package core

import (
	"fmt"
	"gin-project/core/internal"
	"gin-project/global"
	"gin-project/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GVA_CONFIG.Zap.ShowLine {
		// logger.WithOptions(zap.AddCaller())用于向日志添加调用者信息，以帮助追踪日志消息的来源
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

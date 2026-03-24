package core

import (
	"fmt"
	"os"
	"server/core/internal"
	"server/global"
	"server/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	exists, _ := utils.PathExists(global.BIGO_CONFIG.Zap.Director)
	if !exists {
		fmt.Printf("create %v directory\n", global.BIGO_CONFIG.Zap.Director)
		if err := os.Mkdir(global.BIGO_CONFIG.Zap.Director, os.ModePerm); err != nil {
			panic(err)
		}
	}
	// 获取所配置日志级别之上的全部日志级别
	levels := global.BIGO_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	// 将上述创建的多个 Core 合并成一个大的 Core，会进行日志级别分发
	logger = zap.New(zapcore.NewTee(cores...))
	// 当日志级别达到 Error 或更高时，自动记录堆栈跟踪信息，便于排查严重错误
	opts := []zap.Option{zap.AddStacktrace(zapcore.ErrorLevel)}
	if global.BIGO_CONFIG.Zap.ShowLine {
		// 这会让日志输出中包含调用日志函数的文件名和行
		opts = append(opts, zap.AddCaller())
	}
	logger = logger.WithOptions(opts...)
	return logger
}

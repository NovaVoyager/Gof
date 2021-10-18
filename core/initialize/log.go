package initialize

import (
	"fmt"
	"github.com/miaogu-go/Gof/global"
	"github.com/miaogu-go/Gof/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

//LoadLog 加载日志
func LoadLog() {
	checkWithCreateLogDir()
	getLevelFromConf()
	global.GofLog = getLogger()
}

//getLevelFromConf 在配置文件获取日志级别
func getLevelFromConf() {
	//初始化配置文件level
	switch global.GofConfig.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

//checkWithCreateLogDir 检查和创建日志目录
func checkWithCreateLogDir() {
	if ok, _ := utils.PathExists(global.GofConfig.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GofConfig.Zap.Director)
		_ = os.Mkdir(global.GofConfig.Zap.Director, os.ModePerm)
	}
}

//getCore 获取zapcore
func getCore() zapcore.Core {
	writeSync := getLogWrite()
	cores := []zapcore.Core{
		zapcore.NewCore(getEncoder(), writeSync, level),
	}
	if global.GofConfig.Zap.LogInConsole {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level)
		cores = append(cores, consoleCore)
	}
	core := zapcore.NewTee(cores...)
	return core
}

//getLogger 获取日志实例
func getLogger() *zap.Logger {
	core := getCore()
	var logger *zap.Logger

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(core, zap.AddStacktrace(level))
	} else {
		logger = zap.New(core)
	}
	if global.GofConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

//getEncodeConfig 获取日志编码配置
func getEncodeConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  global.GofConfig.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     getCustomTimeFormat,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	/*switch {
	case global.GofConfig.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GofConfig.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GofConfig.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GofConfig.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}*/
	return config
}

//getEncoder 获取编码器
func getEncoder() zapcore.Encoder {
	if global.GofConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncodeConfig())
	}
	return zapcore.NewConsoleEncoder(getEncodeConfig())
}

//getCustomTimeFormat 获取自定义时间输出格式
func getCustomTimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

//getLogWrite 获取日志写入文件对象
func getLogWrite() zapcore.WriteSyncer {
	lumberJackLogger := lumberjack.Logger{
		Filename:   path.Join(global.GofConfig.Zap.Director, "app.log"),
		MaxSize:    global.GofConfig.Zap.MaxSize,
		MaxAge:     global.GofConfig.Zap.MaxAge,
		MaxBackups: global.GofConfig.Zap.MaxBackups,
	}
	return zapcore.AddSync(&lumberJackLogger)
}

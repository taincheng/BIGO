package internal

import (
	"os"
	"server/global"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	core := &ZapCore{
		level: level,
	}
	syncer := core.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	core.Core = zapcore.NewCore(global.BIGO_CONFIG.Zap.Encoder(), syncer, levelEnabler)
	return core
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		global.BIGO_CONFIG.Zap.Director,
		z.level.String(),
		global.BIGO_CONFIG.Zap.KeepDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	if global.BIGO_CONFIG.Zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

package core

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/zhaogongchengsi/starter-gin/config"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

type FileRotatelog struct {
	config *config.Zap
}

func NewFileRotatelog(config *config.Zap) *FileRotatelog {
	return &FileRotatelog{config: config}
}

func (r *FileRotatelog) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(r.config.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(r.config.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if r.config.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

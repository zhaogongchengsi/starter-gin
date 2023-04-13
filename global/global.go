package global

import (
	"github.com/zhaogongchengsi/starter-gin/config"
	"github.com/zhaogongchengsi/starter-gin/core"
	"github.com/zhaogongchengsi/starter-gin/core/store"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	AppConfig = &config.Config{}
)

var (
	//Server           *http.Server  = nil

	Db *gorm.DB = nil

	//Redis            *redis.Client = nil

	Logger *zap.Logger = nil
	// CaptchaStore 存储验证码
	CaptchaStore *store.CaptchaBucket = store.NewCaptchaBucket()
	// Blacklist token 黑名单
	Blacklist *core.Bucket = core.NewBucket(time.Duration(10))
)

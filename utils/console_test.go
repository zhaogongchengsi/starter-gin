package utils_test

import (
	"github.com/zhaogongchengsi/starter-gin/utils"
	"testing"
)

func TestLog(t *testing.T) {
	utils.NewLog().Success("Success %s", "成功")
	utils.NewLog().Warning("Warning %s", "警告")
}

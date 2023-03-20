package utils_test

import (
	"errors"
	"testing"

	"github.com/zhaogongchengsi/starter-gin/utils"
)

func TestPwdStrengthCheck(t *testing.T) {
	_, err := utils.PwdStrengthCheck("1234567", 6, 10)

	if err != nil {
		if !errors.Is(utils.ErrPasswordTooShort, err) {
			t.Errorf("Password length is too short and not detected %s", err)
		}
	}

	_, err = utils.PwdStrengthCheck("1234567", 3, 10)

	if err != nil {
		if !errors.Is(utils.ErrPasswordTooLong, err) {
			t.Errorf("Password length is too short and not detected %s", err)
		}
	}

	le1, err := utils.PwdStrengthCheck("1234567", 6, 10)
	if err != nil {
		t.Errorf("Password length is too short and not detected %s", err)
	}
	le2, err := utils.PwdStrengthCheck("1234567abc", 6, 10)
	if err != nil {
		t.Errorf("Password length is too short and not detected %s", err)
	}
	le3, err := utils.PwdStrengthCheck("1234567abc@", 6, 20)

	if err != nil {
		t.Errorf("Password length is too short and not detected %s", err)
	}

	if le1 != 1 && le2 != 2 && le3 != 3 {
		t.Errorf("Password length detection level error 1234567:%v , 1234567abc:%v 1234567abc@:%v", le1, le2, le3)
	}

}

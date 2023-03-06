package utils

import (
	"errors"
	"regexp"
)

const (
	// 密码长度小于 6 位
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

var ErrPasswordTooLong = errors.New("err:password is too long")
var ErrPasswordTooShort = errors.New("err:password is too short")

func PwdStrengthCheck(pwd string, minlen, maxlen int) (int, error) {
	var level int = levelD
	if len(pwd) < minlen {
		return level, ErrPasswordTooShort
	}

	if len(pwd) > maxlen {
		return level, ErrPasswordTooLong
	}

	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	return level, nil
}

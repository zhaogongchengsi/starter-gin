package utils

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	// BlackText 黑色字
	BlackText = 30 + iota
	// ReaText 红色字
	ReaText
	// GreenText 绿色字
	GreenText
	// YellowText 黄色字
	YellowText
	// BlueText 蓝色字
	BlueText
	// PurpleText 紫色字
	PurpleText
	// SkyBlueText 天蓝字
	SkyBlueText
	// WhiteText 白色字
	WhiteText
)

const (
	BlackBg = 40 + iota
	ReaBg
	GreenBg
	YellowBg
	BlueBg
	PurpleBg
	SkyBlueBg
	WhiteBg
)

var (
	ErrColorExist = errors.New("err: color does not exist")
)

func createColor(tc int, bc ...int) (string, error) {
	if tc < BlackText || tc > WhiteText {
		return "", ErrColorExist
	}
	if len(bc) < 1 {
		return strconv.Itoa(tc) + "m", nil
	}
	bgc := bc[0]
	if bgc < BlackBg || bgc > WhiteBg {
		return "", ErrColorExist
	}
	return fmt.Sprintf("%v;%vm", bgc, tc), nil
}

type Log struct {
}

func NewLog() Log {
	return Log{}
}

func (Log) print(c int, format string, text ...any) {
	color, _ := createColor(c)
	fmt.Printf("\033[%s%s \033[0m", color, fmt.Sprintf(format, text...))
}

func (Log) printWithBg(c, b int, format string, text ...any) {
	color, _ := createColor(c, b)
	fmt.Printf("\033[%s%s \033[0m", color, fmt.Sprintf(format, text...))
}

func (l Log) Success(format string, text ...any) {
	l.print(GreenText, format, text...)
}

func (l Log) Warning(format string, text ...any) {
	l.print(YellowText, format, text...)
}

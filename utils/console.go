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

func CreateColor(tc int, bc ...int) (string, error) {
	return createColor(tc, bc...)
}

var Log = NewConsole()

type Console struct {
}

func NewConsole() Console {
	return Console{}
}

func (Console) print(c int, format string, text ...any) {
	color, _ := createColor(c)
	fmt.Printf("\033[%s%s\033[0m", color, fmt.Sprintf(format, text...))
}

func (Console) printWithBg(c, b int, format string, text ...any) {
	color, _ := createColor(c, b)
	fmt.Printf("\033[%s%s\033[0m", color, fmt.Sprintf(format, text...))
}

func (l Console) Success(format string, text ...any) {
	l.print(GreenText, format, text...)
}

func (l Console) Warning(format string, text ...any) {
	l.print(YellowText, format, text...)
}

func (l Console) Error(format string, text ...any) {
	l.print(ReaText, format, text...)
}

func (l Console) Info(format string, text ...any) {
	l.print(BlueText, format, text...)
}

func Success(format string, text ...any) {
	Log.Success(format, text...)
}

func Warning(format string, text ...any) {
	Log.Warning(format, text...)
}

func Error(format string, text ...any) {
	Log.Error(format, text...)
}

func Info(format string, text ...any) {
	Log.Info(format, text...)
}

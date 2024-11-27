package cosgolog

import (
	"fmt"
)

const (
	Red	= "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Magenta = "\033[35m"
	Cyan = "\033[36m"
	Reset = "\033[0m"
)

func red(v ...any) string {
	return Red + fmt.Sprint(v...) + Reset
}

func redf(format string, v ...any) string {
	return Red + fmt.Sprintf(format, v...) + Reset
}

func green(v ...any) string {
	return Green + fmt.Sprint(v...) + Reset
}

func greenf(format string, v ...any) string {
	return Green + fmt.Sprintf(format, v...) + Reset
}

func yellow(v ...any) string {
	return Yellow + fmt.Sprint(v...) + Reset
}

func yellowf(format string, v ...any) string {
	return Yellow + fmt.Sprintf(format, v...) + Reset
}

func blue(v ...any) string {
	return Blue + fmt.Sprint(v...) + Reset
}

func bluef(format string, v ...any) string {
	return Blue + fmt.Sprintf(format, v...) + Reset
}

func magenta(v ...any) string {
	return Magenta + fmt.Sprint(v...) + Reset
}

func magentaf(format string, v ...any) string {
	return Magenta + fmt.Sprintf(format, v...) + Reset
}

func cyan(v ...any) string {
	return Cyan + fmt.Sprint(v...) + Reset
}

func cyanf(format string, v ...any) string {
	return Cyan + fmt.Sprintf(format, v...) + Reset
}

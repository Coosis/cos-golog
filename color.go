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

func red(prefix string, v ...any) string {
	return Red + prefix + fmt.Sprint(v...) + Reset
}

func redf(prefix string, format string, v ...any) string {
	return Red + prefix + fmt.Sprintf(format, v...) + Reset
}

func green(prefix string, v ...any) string {
	return Green + prefix + fmt.Sprint(v...) + Reset
}

func greenf(prefix string, format string, v ...any) string {
	return Green + prefix + fmt.Sprintf(format, v...) + Reset
}

func yellow(prefix string, v ...any) string {
	return Yellow + prefix + fmt.Sprint(v...) + Reset
}

func yellowf(prefix string, format string, v ...any) string {
	return Yellow + prefix + fmt.Sprintf(format, v...) + Reset
}

func blue(prefix string, v ...any) string {
	return Blue + prefix + fmt.Sprint(v...) + Reset
}

func bluef(prefix string, format string, v ...any) string {
	return Blue + prefix + fmt.Sprintf(format, v...) + Reset
}

func magenta(prefix string, v ...any) string {
	return Magenta + prefix + fmt.Sprint(v...) + Reset
}

func magentaf(prefix string, format string, v ...any) string {
	return Magenta + prefix + fmt.Sprintf(format, v...) + Reset
}

func cyan(prefix string, v ...any) string {
	return Cyan + prefix + fmt.Sprint(v...) + Reset
}

func cyanf(prefix string, format string, v ...any) string {
	return Cyan + prefix + fmt.Sprintf(format, v...) + Reset
}
